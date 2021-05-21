// +build !nopython

package steps

import (
	"context"
	"fmt"
	"runtime"
	"strconv"

	machine "github.com/c3sr/machine/info"
	nvidiasmi "github.com/c3sr/nvidia-smi"

	dl "github.com/c3sr/dlframework"
	"github.com/c3sr/dlframework/framework/feature"
	"github.com/c3sr/dlframework/framework/options"
	"github.com/c3sr/dlframework/framework/predictor"
	"github.com/c3sr/go-python3"
	"github.com/c3sr/pipeline"
	"github.com/c3sr/tracer"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorgonia.org/tensor"
)

type predictGeneral struct {
	base
	predictor predictor.Predictor
	methods   string
}

func NewPredictGeneral(predictor predictor.Predictor, methods string) pipeline.Step {
	res := predictGeneral{
		base: base{
			info: "predict_general_step",
		},
		predictor: predictor,
		methods:   methods,
	}
	res.doer = res.do

	return res
}

func (p predictGeneral) do(ctx context.Context, in0 interface{}, pipelineOpts *pipeline.Options) interface{} {
	iData, ok := in0.([]interface{})
	if !ok {
		return errors.Errorf("expecting []interface{} for predict general step, but got %v", in0)
	}

	data, err := p.castToTensorType(iData)
	if err != nil {
		return err
	}

	if p.predictor == nil {
		return errors.New("the predict general step was created with a nil predictor")
	}

	opts, err := p.predictor.GetPredictionOptions()
	if err != nil {
		return err
	}

	framework, model, err := p.predictor.Info()
	if err != nil {
		return err
	}

	if opentracing.SpanFromContext(ctx) == nil {
		return errors.New("there is no parent span in the context for the predict general step")
	}

	predictTags := opentracing.Tags{
		"trace_source":      "steps",
		"step_name":         "predict_general",
		"model_name":        model.GetName(),
		"model_version":     model.GetVersion(),
		"framework_name":    framework.GetName(),
		"framework_version": framework.GetVersion(),
		"batch_size":        opts.BatchSize(),
		"feature_limit":     opts.FeatureLimit(),
		"device":            opts.Devices().String(),
		"trace_level":       opts.TraceLevel().String(),
		"uses_gpu":          opts.UsesGPU(),
	}
	predictTags["kernel_os"] = machine.Info.KernelOS
	predictTags["os_name"] = machine.Info.OSName
	if opts.GPUMetrics() != "" {
		predictTags["gpu_metrics"] = opts.GPUMetrics()
	}
	if opts.UsesGPU() {
		deviceId := opts.Devices()[0].ID()
		if deviceId > len(nvidiasmi.Info.GPUS) {
			log.WithField("device_id", deviceId).WithField("num_gpus", len(nvidiasmi.Info.GPUS)).Error("unexpected number of gpus")
		} else {
			gpuInfo := nvidiasmi.Info.GPUS[deviceId]
			predictTags["gpu_driver_version"] = nvidiasmi.Info.DriverVersion
			predictTags["gpu_id"] = gpuInfo.ID
			predictTags["gpu_pci_bus"] = gpuInfo.PciBus
			predictTags["gpu_product_name"] = gpuInfo.ProductName
			predictTags["gpu_product_brand"] = gpuInfo.ProductBrand
			predictTags["gpu_persistence_mode"] = gpuInfo.PersistenceMode
		}
	}

	span, newCtx := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, p.Info(), predictTags)

	err = p.predictor.Predict(newCtx, data, options.WithOptions(opts))
	if err != nil {
		return err
	}

	out, err := p.predictor.ReadPredictedFeaturesAsMap(newCtx)
	if err != nil {
		return err
	}

	span.Finish()

	return p.postprocess(ctx, out)
}

func (p predictGeneral) castToTensorType(inputs []interface{}) (interface{}, error) {
	data := make([][]tensor.Tensor, len(inputs))
	for i, input := range inputs {
		v, ok := input.([]tensor.Tensor)
		if !ok {
			return nil, errors.Errorf("unable to cast to []tensor.Tensor in %v step", p.info)
		}
		data[i] = v
	}
	res := make([]tensor.Tensor, len(data[0]))
	for i, _ := range res {
		tmp := make([]tensor.Tensor, len(inputs))
		for j, ten := range data {
			tmp[j] = ten[i]
		}
		var joined tensor.Tensor = tensor.New(tensor.Of(tmp[0].Dtype()), tensor.WithShape(tmp[0].Shape()...))
		if err := tensor.Copy(joined, tmp[0]); err != nil {
			return nil, err
		}
		joined, err := tensor.Concat(0, joined, tmp[1:]...)
		if err != nil {
			return nil, errors.Errorf("unable to concat tensors in %v step", p.info)
		}
		joined.Reshape(append([]int{len(tmp)}, tmp[0].Shape()...)...)
		res[i] = joined
	}
	return res, nil
}

func (p predictGeneral) postprocess(ctx context.Context, in0 interface{}) interface{} {
	if opentracing.SpanFromContext(ctx) != nil {
		span, _ := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, "postprocess_general_step", opentracing.Tags{
			"trace_source": "steps",
			"step_name":    "postprocess_general",
		})
		defer span.Finish()
	}

	table, ok := in0.(map[string]interface{})
	if !ok {
		return errors.Errorf("unable to cast to map[string]interface{} in %v step", p.info)
	}

	if _, exist := table["outputs"]; !exist {
		return errors.Errorf("outputs is not a key in %v step", p.info)
	}

	tensors, ok := (table["outputs"]).([]tensor.Tensor)
	if !ok {
		return errors.Errorf("unable to cast to []tensor.Tensor in %v step", p.info)
	}

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	pyState := python3.PyGILState_Ensure()
	defer python3.PyGILState_Release(pyState)

	python3.PyRun_SimpleString(p.methods)
	pyMain := python3.PyImport_AddModule("__main__")
	pyDict := python3.PyModule_GetDict(pyMain)
	pyPostprocess := python3.PyDict_GetItemString(pyDict, "postprocess")

	pyCtx := python3.PyDict_New()
	defer pyCtx.DecRef()

	python3.PyRun_SimpleString(`
def convert(shape, ptr, data_type):
  import ctypes
  import numpy as np
  type_dict = {
    'float32' : ctypes.c_float,
    'float64' : ctypes.c_double,
    'uint8'   : ctypes.c_uint8,
    'uint16'  : ctypes.c_uint16,
    'uint32'  : ctypes.c_uint32,
    'uint64'  : ctypes.c_uint64,
    'int8'    : ctypes.c_int8,
    'int16'   : ctypes.c_int16,
    'int32'   : ctypes.c_int32,
    'int64'   : ctypes.c_int64
  }
  return np.ctypeslib.as_array(ctypes.cast(ptr, ctypes.POINTER(type_dict[data_type])), shape)`)
	pyConvert := python3.PyDict_GetItemString(pyDict, "convert")

	pyTensors := python3.PyList_New(0)
	for _, t := range tensors {
		dense, ok := t.(*tensor.Dense)
		if !ok {
			return errors.New("expecting a dense tensor")
		}
		pyShapeTuple := python3.PyTuple_New(len(dense.Shape()))
		for i, v := range dense.Shape() {
			cur := python3.PyLong_FromGoInt(v)
			// The reference is stolen hence no need to DecRef
			python3.PyTuple_SetItem(pyShapeTuple, i, cur)
		}

		dType := fmt.Sprintf("%v", dense.Dtype().Type)
		ptr, _ := strconv.ParseUint(fmt.Sprintf("%v", dense.Uintptr()), 0, 64)

		pyDataPtr := python3.PyLong_FromUnsignedLongLong(ptr)
		defer pyDataPtr.DecRef()

		pyDataType := python3.PyUnicode_FromString(dType)
		defer pyDataType.DecRef()

		ret := pyConvert.CallFunctionObjArgs(pyShapeTuple, pyDataPtr, pyDataType)
		defer ret.DecRef()

		python3.PyList_Append(pyTensors, ret)
	}

	defer pyTensors.DecRef()

	pyOutputs := pyPostprocess.CallFunctionObjArgs(pyCtx, pyTensors)
	defer pyOutputs.DecRef()

	modality, err := p.predictor.Modality()
	if err != nil {
		return err
	}

	var features []dl.Features

	switch modality {
	case dl.ImageClassificationModality:
		if _, exist := table["labels"]; !exist {
			return errors.Errorf("labels is not a key in %v step", p.info)
		}
		labels, ok := (table["labels"]).([]string)
		if !ok {
			return errors.Errorf("unable to cast to []string in %v step", p.info)
		}
		features = p.convertImageClassification(pyOutputs, labels)
	case dl.ImageObjectDetectionModality:
		if _, exist := table["labels"]; !exist {
			return errors.Errorf("labels is not a key in %v step", p.info)
		}
		labels, ok := (table["labels"]).([]string)
		if !ok {
			return errors.Errorf("unable to cast to []string in %v step", p.info)
		}
		features = p.convertImageObjectDetection(pyOutputs, labels)
	case dl.ImageInstanceSegmentationModality:
		if _, exist := table["labels"]; !exist {
			return errors.Errorf("labels is not a key in %v step", p.info)
		}
		labels, ok := (table["labels"]).([]string)
		if !ok {
			return errors.Errorf("unable to cast to []string in %v step", p.info)
		}
		features = p.convertImageInstanceSegmentation(pyOutputs, labels)
	case dl.ImageSemanticSegmentationModality:
		features = p.convertImageSemanticSegmentation(pyOutputs)
	case dl.ImageEnhancementModality:
		features = p.convertImageEnhancement(pyOutputs)
	default:
		return errors.New("unsupported modality")
	}

	runtime.KeepAlive(tensors)

	lst := make([]interface{}, len(features))
	for i := 0; i < len(features); i++ {
		lst[i] = features[i]
	}

	return lst
}

// convertImageClassification expects pyOutputs to be probabilities
// probabilities[i][j]: probability for the i'th input to be the j'th item
func (p predictGeneral) convertImageClassification(pyOutputs *python3.PyObject, labels []string) []dl.Features {
	probabilities := make([][]float32, python3.PyList_Size(pyOutputs))
	for i, _ := range probabilities {
		// borrowed reference
		cur := python3.PyList_GetItem(pyOutputs, i)

		length := python3.PyList_Size(cur)
		probabilities[i] = make([]float32, length)

		for j := 0; j < length; j++ {
			// borrowed reference
			val := python3.PyList_GetItem(cur, j)

			probabilities[i][j] = float32(python3.PyFloat_AsDouble(val))
		}
	}

	return feature.CreateClassificationFeaturesCanonical(probabilities, labels)
}

// convertImageObjectDetection expects pyOutputs to be a tuple, (probabilities, classes, boxes)
// probabilities[i][j]: probability for the j'th box in the i'th input to be classes[i][j]
// classes[i][j]: predicted class for the j'th box in the i'th input, which is the argmax among all labels
// boxes[i][j][0:4]: bounding boxes for the j'th box in the i'th input, following (Ymin, Xmin, Ymax, Xmax) \in [0, 1]^4
func (p predictGeneral) convertImageObjectDetection(pyOutputs *python3.PyObject, labels []string) []dl.Features {
	// borrowed references
	pyProb := python3.PyTuple_GetItem(pyOutputs, 0)
	pyClass := python3.PyTuple_GetItem(pyOutputs, 1)
	pyBox := python3.PyTuple_GetItem(pyOutputs, 2)

	probabilities := make([][]float32, python3.PyList_Size(pyProb))
	classes := make([][]float32, python3.PyList_Size(pyProb))
	boxes := make([][][4]float32, python3.PyList_Size(pyProb))

	for i, _ := range probabilities {
		// borrowed reference
		curProb := python3.PyList_GetItem(pyProb, i)
		curClass := python3.PyList_GetItem(pyClass, i)
		curBox := python3.PyList_GetItem(pyBox, i)

		length := python3.PyList_Size(curProb)
		probabilities[i] = make([]float32, length)
		classes[i] = make([]float32, length)
		boxes[i] = make([][4]float32, length)

		for j := 0; j < length; j++ {
			probabilities[i][j] = float32(python3.PyFloat_AsDouble(python3.PyList_GetItem(curProb, j)))
			classes[i][j] = float32(python3.PyFloat_AsDouble(python3.PyList_GetItem(curClass, j)))
			box := python3.PyList_GetItem(curBox, j)
			for k := 0; k < 4; k++ {
				boxes[i][j][k] = float32(python3.PyFloat_AsDouble(python3.PyList_GetItem(box, k)))
			}
		}
	}

	return feature.CreateBoundingBoxFeaturesCanonical(probabilities, classes, boxes, labels)
}

// convertImageSemanticSegmentation expects pyOutputs to be a 3D masks[B][H][W]
// masks[i][j][k]: mask for the i'th image at coordinate (j, k)
func (p predictGeneral) convertImageSemanticSegmentation(pyOutputs *python3.PyObject) []dl.Features {
	masks := make([][][]int64, python3.PyList_Size(pyOutputs))

	for i, _ := range masks {
		// borrowed reference
		cur := python3.PyList_GetItem(pyOutputs, i)

		height := python3.PyList_Size(cur)
		masks[i] = make([][]int64, height)

		for j := 0; j < height; j++ {
			// borrowed reference
			curRow := python3.PyList_GetItem(cur, j)
			width := python3.PyList_Size(curRow)
			masks[i][j] = make([]int64, width)

			for k := 0; k < width; k++ {
				// borrowed reference
				val := python3.PyList_GetItem(cur, k)
				masks[i][j][k] = int64(python3.PyFloat_AsDouble(val))
			}
		}
	}

	return feature.CreateSemanticSegmentFeaturesCanonical(masks)
}

// convertImageEnhancement expects pyOutputs to be a 4D list images[B][H][W][C]
// images[i][j][k][l]: pixel for the i'th image at channel l at coordinate (j, k)
func (p predictGeneral) convertImageEnhancement(pyOutputs *python3.PyObject) []dl.Features {
	images := make([][][][]float32, python3.PyList_Size(pyOutputs))

	for i, _ := range images {
		// borrowed reference
		curImage := python3.PyList_GetItem(pyOutputs, i)
		height := python3.PyList_Size(curImage)
		images[i] = make([][][]float32, height)

		for j := 0; j < height; j++ {
			// borrowed reference
			curH := python3.PyList_GetItem(curImage, j)
			width := python3.PyList_Size(curH)
			images[i][j] = make([][]float32, width)

			for k := 0; k < width; k++ {
				// borrowed reference
				curW := python3.PyList_GetItem(curH, k)
				channel := python3.PyList_Size(curW)
				images[i][j][k] = make([]float32, channel)

				for l := 0; l < channel; l++ {
					val := python3.PyList_GetItem(curW, l)
					images[i][j][k][l] = float32(python3.PyFloat_AsDouble(val))
				}
			}
		}
	}

	return feature.CreateRawImageFeaturesCanonical(images)
}

// convertImageInstanceSegmentation expects pyOutputs to be a tuple, (probabilities, classes, boxes, masks)
// probabilities[i][j]: probability for the j'th box in the i'th input to be classes[i][j]
// classes[i][j]: predicted class for the j'th box in the i'th input, which is the argmax among all labels
// boxes[i][j][0:4]: bounding boxes for the j'th box in the i'th input, following (Ymin, Xmin, Ymax, Xmax) \in ([0, H) or [0, W))^4
// masks[i][j][k][l]: masks for the j'th item in the i'th input at coordinate (k, l) in the box
func (p predictGeneral) convertImageInstanceSegmentation(pyOutputs *python3.PyObject, labels []string) []dl.Features {
	// borrowed references
	pyProb := python3.PyTuple_GetItem(pyOutputs, 0)
	pyClass := python3.PyTuple_GetItem(pyOutputs, 1)
	pyBox := python3.PyTuple_GetItem(pyOutputs, 2)
	pyMask := python3.PyTuple_GetItem(pyOutputs, 3)

	probabilities := make([][]float32, python3.PyList_Size(pyProb))
	classes := make([][]float32, python3.PyList_Size(pyProb))
	boxes := make([][][4]float32, python3.PyList_Size(pyProb))
	masks := make([][][][]float32, python3.PyList_Size(pyProb))

	for i, _ := range probabilities {
		// borrowed reference
		curProb := python3.PyList_GetItem(pyProb, i)
		curClass := python3.PyList_GetItem(pyClass, i)
		curBox := python3.PyList_GetItem(pyBox, i)
		curMask := python3.PyList_GetItem(pyMask, i)

		length := python3.PyList_Size(curProb)
		probabilities[i] = make([]float32, length)
		classes[i] = make([]float32, length)
		boxes[i] = make([][4]float32, length)
		masks[i] = make([][][]float32, length)

		for j := 0; j < length; j++ {
			probabilities[i][j] = float32(python3.PyFloat_AsDouble(python3.PyList_GetItem(curProb, j)))
			classes[i][j] = float32(python3.PyFloat_AsDouble(python3.PyList_GetItem(curClass, j)))
			box := python3.PyList_GetItem(curBox, j)
			for k := 0; k < 4; k++ {
				boxes[i][j][k] = float32(python3.PyFloat_AsDouble(python3.PyList_GetItem(box, k)))
			}
			mask := python3.PyList_GetItem(curMask, j)
			height := python3.PyList_Size(mask)
			masks[i][j] = make([][]float32, height)

			for k := 0; k < height; k++ {
				cur := python3.PyList_GetItem(mask, k)
				width := python3.PyList_Size(cur)
				masks[i][j][k] = make([]float32, width)

				for l := 0; l < width; l++ {
					masks[i][j][k][l] = float32(python3.PyFloat_AsDouble(python3.PyList_GetItem(cur, l)))
				}
			}
		}
	}

	return feature.CreateInstanceSegmentFeaturesCanonical(probabilities, classes, boxes, masks, labels)
}

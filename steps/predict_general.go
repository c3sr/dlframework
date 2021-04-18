// +build !nopython

package steps

import (
	"context"
	"fmt"
	"runtime"
	"strconv"

	machine "github.com/c3sr/machine/info"
	nvidiasmi "github.com/c3sr/nvidia-smi"

	"github.com/DataDog/go-python3"
	dl "github.com/c3sr/dlframework"
	"github.com/c3sr/dlframework/framework/feature"
	"github.com/c3sr/dlframework/framework/options"
	"github.com/c3sr/dlframework/framework/predictor"
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
		errors.New("there is no parent span in the context for the predict general step")
	}

	predictTags := opentracing.Tags{
		"trace_source":      "steps",
		"step_name":         "predict general",
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

	span, newCtx := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, "postprocess_general_step", predictTags)

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
	data := make([]*tensor.Dense, len(inputs))
	for ii, input := range inputs {
		v, ok := input.(*tensor.Dense)
		if !ok {
			return nil, errors.Errorf("unable to cast to dense tensor in %v step", p.info)
		}
		data[ii] = v
	}
	return data, nil
}

func (p predictGeneral) postprocess(ctx context.Context, in0 interface{}) interface{} {
	if opentracing.SpanFromContext(ctx) != nil {
		span, _ := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, p.Info(), opentracing.Tags{
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

	python3.PyRun_SimpleString("def convert(shape, ptr):\n  import ctypes\n  import numpy as np\n  return np.ctypeslib.as_array(ctypes.cast(ptr, ctypes.POINTER(ctypes.c_float)), shape)")
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
		ptr, _ := strconv.ParseUint(fmt.Sprintf("%v", &(t.Data().([]float32))[0]), 0, 64)

		pyDataPtr := python3.PyLong_FromUnsignedLongLong(ptr)
		defer pyDataPtr.DecRef()

		ret := pyConvert.CallFunctionObjArgs(pyShapeTuple, pyDataPtr)
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
	// case dl.ImageObjectDetectionModality:
	// 	if _, exist := table["labels"]; !exist {
	// 		return errors.Errorf("labels is not a key in %v step", p.info)
	// 	}
	// 	labels, ok := (table["labels"]).([]string)
	// 	if !ok {
	// 		return errors.Errorf("unable to cast to []string in %v step", p.info)
	// 	}
	// 	features = p.convertImageObjectDetection(pyOutputs, labels)
	// case dl.ImageInstanceSegmentationModality:
	// 	if _, exist := table["labels"]; !exist {
	// 		return errors.Errorf("labels is not a key in %v step", p.info)
	// 	}
	// 	labels, ok := (table["labels"]).([]string)
	// 	if !ok {
	// 		return errors.Errorf("unable to cast to []string in %v step", p.info)
	// 	}
	// 	features = p.convertImageInstanceSegmentation(pyOutputs, labels)
	// case dl.ImageSemanticSegmentationModality:
	// 	if _, exist := table["labels"]; !exist {
	// 		return errors.Errorf("labels is not a key in %v step", p.info)
	// 	}
	// 	labels, ok := (table["labels"]).([]string)
	// 	if !ok {
	// 		return errors.Errorf("unable to cast to []string in %v step", p.info)
	// 	}
	// 	features = p.convertImageSemanticSegmentation(pyOutputs, labels)
	// case dl.ImageEnhancementModality:
	// 	features = p.convertImageEnhancement(pyOutputs)
	default:
		return errors.New("unsupported modality")
	}

	runtime.KeepAlive(tensors)

	lst := make([]interface{}, len(features))
	for i := 0; i < len(features); i++ {
		fmt.Println(features[i][0].GetProbability(), features[i][0].GetClassification().GetLabel())
		lst[i] = features[i]
	}

	return lst
}

// convertImageClassification expects pyOutputs to be a 2d list where the first dimension is batchsize
// and the second dimension is number of labels. The element is probability for each labels
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

package steps

import (
	"context"

	machine "github.com/c3sr/machine/info"
	nvidiasmi "github.com/c3sr/nvidia-smi"

	"github.com/c3sr/dlframework/framework/options"
	"github.com/c3sr/dlframework/framework/predictor"
	"github.com/c3sr/pipeline"
	"github.com/c3sr/tracer"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorgonia.org/tensor"
)

type predict struct {
	base
	predictor predictor.Predictor
}

func NewPredict(predictor predictor.Predictor) pipeline.Step {
	res := predict{
		base: base{
			info: "predict_step",
		},
	}
	res.predictor = predictor
	res.doer = res.do

	return res
}

func (p predict) do(ctx context.Context, in0 interface{}, pipelineOpts *pipeline.Options) interface{} {
	iData, ok := in0.([]interface{})
	if !ok {
		return errors.Errorf("expecting []interface{} for predict image step, but got %v", in0)
	}

	data, err := p.castToTensorType(iData)
	if err != nil {
		return err
	}

	if p.predictor == nil {
		return errors.New("the predict image step was created with a nil predictor")
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
		errors.New("there is no parent span in the context for the predict step")
	}

	predictTags := opentracing.Tags{
		"trace_source":      "steps",
		"step_name":         "predict",
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

	span, ctx := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, p.Info(), predictTags)
	defer span.Finish()

	err = p.predictor.Predict(ctx, data, options.WithOptions(opts))
	if err != nil {
		return err
	}

	features, err := p.predictor.ReadPredictedFeatures(ctx)
	if err != nil {
		return err
	}

	lst := make([]interface{}, len(iData))
	for ii := 0; ii < len(iData); ii++ {
		lst[ii] = features[ii]
	}

	return lst
}

func (p predict) castToTensorType(inputs []interface{}) (interface{}, error) {
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

func (p predict) castToElementType(inputs []interface{}) (interface{}, error) {
	_, model, _ := p.predictor.Info()
	switch t := model.GetElementType(); t {
	case "raw_image":
		data := make([][]byte, len(inputs))
		for ii, input := range inputs {
			r, err := toByteSlice(input)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to cast to uint8 slice in %v step", p.info)
			}
			data[ii] = r
		}
		return data, nil
	case "uint8":
		data := make([][]uint8, len(inputs))
		for ii, input := range inputs {
			r, err := toUInt8Slice(input)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to cast to uint8 slice in %v step", p.info)
			}
			data[ii] = r
		}
		return data, nil
	case "float32":
		data := make([][]float32, len(inputs))
		for ii, input := range inputs {
			r, err := toFloat32Slice(input)
			if err != nil {
				return nil, errors.Wrapf(err, "unable to cast to float32 slice in %v step", p.info)
			}
			data[ii] = r
		}
		return data, nil
	default:
		return nil, errors.Errorf("unsupported element type %v", t)
	}
}

func (p predict) Close() error {
	return nil
}

package steps

import (
	"context"

	"github.com/oliamb/cutter"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/c3sr/dlframework/framework/predictor"
	"github.com/c3sr/image/types"
	"github.com/c3sr/pipeline"
	"github.com/c3sr/tracer"
)

type cropDimension struct {
	base
	options predictor.PreprocessOptions
}

func NewCropDimension(options predictor.PreprocessOptions) pipeline.Step {
	res := cropDimension{
		base: base{
			info: "crop_dimension_step",
		},
		options: options,
	}
	res.doer = res.do
	return res
}

func (p cropDimension) do(ctx context.Context, in0 interface{}, opts *pipeline.Options) interface{} {
	if opentracing.SpanFromContext(ctx) != nil {
		span, _ := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, p.Info(), opentracing.Tags{
			"trace_source": "steps",
			"step_name":    "crop_image",
		})
		defer span.Finish()
	}

	in, ok := in0.(types.Image)
	if !ok {
		return errors.Errorf("expecting a io.Reader or dataset element for read image step, but got %v", in0)
	}

	cropDims := p.options.CropDims
	if cropDims == nil {
		return in
	}

	croppedImg, err := cutter.Crop(in, cutter.Config{
		Width:  cropDims[0],
		Height: cropDims[1],
		Mode:   p.options.CropMethod,
	})
	if err != nil {
		return errors.Errorf("unable to crop image")
	}

	return croppedImg
}

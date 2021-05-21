package steps

/*

import (
	"context"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/c3sr/dldataset"
	_ "github.com/c3sr/dldataset/vision"
	"github.com/c3sr/pipeline"
	"github.com/c3sr/tracer"
)

type getDataset struct {
	base
	dataset dldataset.Dataset
}

func NewGetDataset(dataset dldataset.Dataset) pipeline.Step {
	res := getDataset{
		dataset: dataset,
		base: base{
			info: "get_dataset_step",
		},
	}
	res.doer = res.do
	return res
}

func (p getDataset) do(ctx context.Context, in0 interface{}, opts *pipeline.Options) interface{} {
	if span, newCtx := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, p.Info(), opentracing.Tags{
		"trace_source": "steps",
		"step_name":    "predict",
	}); span != nil {
		ctx = newCtx
		defer span.Finish()
	}

	in, ok := in0.(string)
	if !ok {
		return errors.Errorf("expecting a string for get dataset step, but got %v", in0)
	}

	lbl, err := p.dataset.Get(ctx, in)
	if err != nil {
		return err
	}

	return lbl
}

func (p getDataset) Close() error {
	return nil
}
*/

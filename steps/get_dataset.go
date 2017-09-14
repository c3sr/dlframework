package steps

import (
	"golang.org/x/net/context"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rai-project/dldataset"
	_ "github.com/rai-project/dldataset/vision"
	"github.com/rai-project/pipeline"
)

type getDataset struct {
	base
	dataset dldataset.Dataset
}

func NewGetDataset(dataset dldataset.Dataset) pipeline.Step {
	res := getDataset{
		dataset: dataset,
		base: base{
			info: "GetDataset",
		},
	}
	res.doer = res.do
	return res
}

func (p getDataset) do(ctx context.Context, in0 interface{}) interface{} {
	if span, newCtx := opentracing.StartSpanFromContext(ctx, "Get Dataset Step"); span != nil {
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
package steps

import (
	"bytes"
	"context"
	"io"
	"net/http"

	dl "github.com/c3sr/dlframework"
	"github.com/c3sr/pipeline"
	"github.com/c3sr/tracer"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type readURL struct {
	base
}

func NewReadURL() pipeline.Step {
	res := readURL{
		base: base{
			info: "read_url_step",
		},
	}
	res.doer = res.do
	return res
}

func (p readURL) do(ctx context.Context, in0 interface{}, opts *pipeline.Options) interface{} {
	var span opentracing.Span
	if opentracing.SpanFromContext(ctx) != nil {
		span, ctx = tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, p.Info(), opentracing.Tags{
			"trace_source": "steps",
			"step_name":    "read_url",
		})
		defer span.Finish()
	}

	url := ""
	switch in := in0.(type) {
	case string:
		url = in
	case *dl.URLsRequest_URL:
		if in == nil {
			return errors.New("cannot read nil url")
		}
		url = in.GetData()
	default:
		return errors.Errorf("expecting a string or *dl.URLsRequest_URL for read url Step, but got %v", in0)
	}

	if span != nil {
		span.SetTag("url", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.Errorf("bad response code: %d", resp.StatusCode)
	}

	res := new(bytes.Buffer)
	_, err = io.Copy(res, resp.Body)
	if err != nil {
		return errors.Errorf("unable to copy body")
	}
	return res
}

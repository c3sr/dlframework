package steps

import (
	"context"
	"encoding/json"
	"os/exec"
	"strings"

	dl "github.com/c3sr/dlframework"
	"github.com/c3sr/dlframework/framework/predictor"
	"github.com/c3sr/pipeline"
	"github.com/c3sr/tracer"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"gorgonia.org/tensor"
)

type preprocessGeneral struct {
	base
	options predictor.PreprocessOptions
	methods string
}

func NewPreprocessGeneral(options predictor.PreprocessOptions, methods string) pipeline.Step {
	res := preprocessGeneral{
		base: base{
			info: "preprocess_general_step",
		},
		options: options,
		methods: methods,
	}
	res.doer = res.do
	return res
}

func (p *preprocessGeneral) do(ctx context.Context, in0 interface{}, pipelineOptions *pipeline.Options) interface{} {
	if opentracing.SpanFromContext(ctx) != nil {
		span, _ := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, p.Info(), opentracing.Tags{
			"trace_source": "steps",
			"step_name":    "preprocess_general",
		})
		defer span.Finish()
	}

	src := ""
	switch in := in0.(type) {
	case string:
		src = in
	case *dl.URLsRequest_URL:
		if in == nil {
			return errors.New("cannot read nil url")
		}
		src = in.GetData()
	default:
		return errors.Errorf("expecting a string but got %v", in0)
	}

	if len(p.methods) == 0 {
		return src
	}

	preprocessMethod := strings.Split(p.methods, " ")

	cmd := exec.Command(preprocessMethod[0], preprocessMethod[1:len(preprocessMethod)]...)
	cmd.Stdin = strings.NewReader(src)

	jsonString, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	elementType := strings.ToLower(p.options.ElementType)

	switch elementType {
	case "float32":
		var data [][][]float32
		err = json.Unmarshal(jsonString, &data)
		if err != nil {
			return errors.Errorf("invalid return from preprocess methods.")
		}
		var flattenData []float32
		for i, _ := range data {
			for j, _ := range data[i] {
				flattenData = append(flattenData, data[i][j]...)
			}
		}
		outTensor := tensor.New(
			tensor.WithShape(len(data), len(data[0]), len(data[0][0])),
			tensor.WithBacking(flattenData),
		)
		return outTensor
	case "uint8":
		var data [][][]uint8
		err = json.Unmarshal(jsonString, &data)
		if err != nil {
			return errors.Errorf("invalid return from preprocess methods.")
		}
		var flattenData []uint8
		for i, _ := range data {
			for j, _ := range data[i] {
				flattenData = append(flattenData, data[i][j]...)
			}
		}
		outTensor := tensor.New(
			tensor.WithShape(len(data), len(data[0]), len(data[0][0])),
			tensor.WithBacking(flattenData),
		)
		return outTensor
	}

	return errors.Errorf("unsupported element type %v", elementType)
}

// +build !nopython

package steps

import "C"

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"unsafe"

	dl "github.com/c3sr/dlframework"
	"github.com/c3sr/dlframework/framework/predictor"
	"github.com/c3sr/go-python3"
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

func (p preprocessGeneral) do(ctx context.Context, in0 interface{}, pipelineOptions *pipeline.Options) interface{} {
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

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	pyState := python3.PyGILState_Ensure()
	defer python3.PyGILState_Release(pyState)

	python3.PyRun_SimpleString(p.methods)
	pyMain := python3.PyImport_AddModule("__main__")
	pyDict := python3.PyModule_GetDict(pyMain)
	pyPreprocess := python3.PyDict_GetItemString(pyDict, "preprocess")

	pyCtx := python3.PyDict_New()
	defer pyCtx.DecRef()

	pySrc := python3.PyUnicode_FromString(src)
	defer pySrc.DecRef()

	npArray := pyPreprocess.CallFunctionObjArgs(pyCtx, pySrc)
	defer npArray.DecRef()

	npShapeObj := npArray.GetAttrString("shape")
	defer npShapeObj.DecRef()

	npShapeRepr := npShapeObj.Repr()
	defer npShapeRepr.DecRef()

	npShape := python3.PyUnicode_AsUTF8(npShapeRepr)

	python3.PyRun_SimpleString(`
def contiguous(x):
  import numpy as np
  x = np.ascontiguousarray(x, dtype = np.float32)
  return x.ctypes.data`)

	pyContiguous := python3.PyDict_GetItemString(pyDict, "contiguous")

	npDataObj := pyContiguous.CallFunctionObjArgs(npArray)
	defer npDataObj.DecRef()

	npDataRepr := npDataObj.Repr()
	defer npDataRepr.DecRef()

	npDataPtr := python3.PyUnicode_AsUTF8(npDataRepr)

	shape, err := p.parseShape(npShape)
	if err != nil {
		return err
	}

	elementType := strings.ToLower(p.options.ElementType)

	switch elementType {
	case "float32":
		flattenData, err := p.parseDataAsFloat(shape, npDataPtr)
		if err != nil {
			return err
		}

		outTensor := tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(flattenData),
		)

		return outTensor
	case "uint8":
		flattenData, err := p.parseDataAsUInt8(shape, npDataPtr)
		if err != nil {
			return err
		}

		outTensor := tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(flattenData),
		)
		return outTensor
	}

	return errors.Errorf("unsupported element type %v", elementType)
}

func (p preprocessGeneral) parseShape(s string) (res []int, err error) {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			continue
		}
		j := i + 1
		for j < len(s) && s[j] >= '0' && s[j] <= '9' {
			j++
		}
		cur, err := strconv.Atoi(s[i:j])
		if err != nil {
			return res, err
		}
		res = append(res, cur)
		i = j - 1
	}
	return res, err
}

func (p preprocessGeneral) parseDataAsFloat(shape []int, s string) ([]float32, error) {
	sz := 1
	for _, v := range shape {
		sz *= v
	}
	res := make([]float32, sz)

	ptr, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	slice := (*[1 << 30]C.float)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
	for i := 0; i < sz; i++ {
		res[i] = float32(slice[i])
	}
	return res, nil
}

func (p preprocessGeneral) parseDataAsUInt8(shape []int, s string) ([]uint8, error) {
	sz := 1
	for _, v := range shape {
		sz *= v
	}
	res := make([]uint8, sz)

	ptr, err := strconv.Atoi(s)
	if err != nil {
		return nil, err
	}

	slice := (*[1 << 30]C.float)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
	for i := 0; i < sz; i++ {
		res[i] = uint8(slice[i])
	}
	return res, nil
}

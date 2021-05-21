// +build !nopython

package steps

import "C"

import (
	"context"
	"runtime"
	"strconv"
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
	python3.PyRun_SimpleString(`
def contiguous(x):
  import numpy as np
  return np.ascontiguousarray(x)`)
	python3.PyRun_SimpleString(`
def get_address(x):
  return x.ctypes.data`)
	pyMain := python3.PyImport_AddModule("__main__")
	pyDict := python3.PyModule_GetDict(pyMain)
	pyPreprocess := python3.PyDict_GetItemString(pyDict, "preprocess")
	pyContiguous := python3.PyDict_GetItemString(pyDict, "contiguous")
	pyGetAddress := python3.PyDict_GetItemString(pyDict, "get_address")

	pyCtx := python3.PyDict_New()
	defer pyCtx.DecRef()

	pySrc := python3.PyUnicode_FromString(src)
	defer pySrc.DecRef()

	var res []tensor.Tensor

	pyPreprocessedData := pyPreprocess.CallFunctionObjArgs(pyCtx, pySrc)
	defer pyPreprocessedData.DecRef()

	convertToTensor := func(npArrayRaw *python3.PyObject) (tensor.Tensor, error) {
		npArrayContiguous := pyContiguous.CallFunctionObjArgs(npArrayRaw)
		defer npArrayContiguous.DecRef()

		npShapeObj := npArrayContiguous.GetAttrString("shape")
		defer npShapeObj.DecRef()

		npShapeRepr := npShapeObj.Repr()
		defer npShapeRepr.DecRef()

		npShape := python3.PyUnicode_AsUTF8(npShapeRepr)

		npDtypeObj := npArrayContiguous.GetAttrString("dtype")
		defer npDtypeObj.DecRef()

		npDtypeRepr := npDtypeObj.Repr()
		defer npDtypeRepr.DecRef()

		npDtype := python3.PyUnicode_AsUTF8(npDtypeRepr)
		// npDtype = "dtype('*')"
		npDtype = npDtype[7 : len(npDtype)-2]

		npDataObj := pyGetAddress.CallFunctionObjArgs(npArrayContiguous)
		defer npDataObj.DecRef()

		npDataRepr := npDataObj.Repr()
		defer npDataRepr.DecRef()

		npDataPtr := python3.PyUnicode_AsUTF8(npDataRepr)

		return p.parseData(npShape, npDataPtr, npDtype)
	}

	if python3.PyTuple_Check(pyPreprocessedData) == false {
		// Only one input
		ten, err := convertToTensor(pyPreprocessedData)
		if err != nil {
			return err
		}
		res = append(res, ten)
	} else {
		// Several inputs
		for i := 0; i < python3.PyTuple_Size(pyPreprocessedData); i++ {
			ten, err := convertToTensor(python3.PyTuple_GetItem(pyPreprocessedData, i))
			if err != nil {
				return err
			}
			res = append(res, ten)
		}
	}

	return res
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

func (p preprocessGeneral) parseData(npShape string, npDataPtr string, npDtype string) (tensor.Tensor, error) {
	shape, err := p.parseShape(npShape)
	if err != nil {
		return nil, err
	}

	ptr, err := strconv.Atoi(npDataPtr)
	if err != nil {
		return nil, err
	}

	sz := 1
	for _, v := range shape {
		sz *= v
	}

	switch npDtype {
	case "float32":
		cData := (*[1 << 30]float32)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]float32, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "float64":
		cData := (*[1 << 30]float64)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]float64, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "uint8":
		cData := (*[1 << 30]uint8)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]uint8, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "uint16":
		cData := (*[1 << 30]uint16)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]uint16, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "uint32":
		cData := (*[1 << 30]uint32)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]uint32, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "uint64":
		cData := (*[1 << 30]uint64)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]uint64, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "int8":
		cData := (*[1 << 30]int8)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]int8, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "int16":
		cData := (*[1 << 30]int16)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]int16, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "int32":
		cData := (*[1 << 30]int32)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]int32, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	case "int64":
		cData := (*[1 << 30]int64)(unsafe.Pointer(uintptr(ptr)))[:sz:sz]
		data := make([]int64, sz)
		copy(data, cData)
		return tensor.New(
			tensor.WithShape(shape...),
			tensor.WithBacking(data),
		), nil
	}

	return nil, errors.Errorf("unsupported element type %v", npDtype)
}

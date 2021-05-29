package plotting

import (
	"os"

	"github.com/mailru/easyjson"

	"github.com/unknwon/com"
	"github.com/pkg/errors"
	"github.com/c3sr/dlframework/evaluation"
	model "github.com/uber/jaeger/model/json"
)

func ReadTraceFile(path string) ([]model.Span, error) {
	if !com.IsFile(path) {
		return nil, errors.Errorf("the trace file %v was not found", path)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to open the trace file %v", path)
	}
	defer f.Close()

	var trace evaluation.TraceInformation
	err = easyjson.UnmarshalFromReader(f, &trace)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to decode the trace file %v", path)
	}

	if len(trace.Traces) == 0 {
		return nil, errors.Wrapf(err, "no traces were found in %v", path)
	}

	return trace.Traces[0].Spans, nil
}

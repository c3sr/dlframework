package agent

import "github.com/c3sr/dlframework/framework/predictor"

type PredictorLifetime struct {
	Predictor      *predictor.Predictor
	ReferenceCount uint64
}

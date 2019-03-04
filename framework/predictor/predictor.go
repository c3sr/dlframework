package predictor

import (
	"io"

	"context"

	"github.com/rai-project/dlframework"
	"github.com/rai-project/dlframework/framework/options"
)

type Predictor interface {
	// Gets framework and model manifests
	Info() (dlframework.FrameworkManifest, dlframework.ModelManifest, error)
	// Downloads model from manifest
	Download(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) error
	// Load model from manifest
	Load(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) (Predictor, error)
	// Returns the prediction options
	GetPredictionOptions(ctx context.Context) (*options.Options, error)
	// Returns the preprocess options
	GetPreprocessOptions(ctx context.Context) (PreprocessOptions, error)
	// Returns the handle to features
	Predict(ctx context.Context, data [][]float32, opts ...options.Option) error
	// Returns the features
	ReadPredictedFeatures(ctx context.Context) ([]dlframework.Features, error)
	// Clears the internal state of a predictor
	Reset(ctx context.Context) error

	io.Closer
}

package predictor

import (
	"io"
	"context"

	"github.com/c3sr/dlframework"
	"github.com/c3sr/dlframework/framework/options"
)

type Predictor interface {
	// Gets framework and model manifests
	Info() (dlframework.FrameworkManifest, dlframework.ModelManifest, error)
	// Gets predictor's Modality
	Modality() (dlframework.Modality, error)
	// Downloads model from manifest
	Download(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) error
	// Load model from manifest
	Load(ctx context.Context, model dlframework.ModelManifest, opts ...options.Option) (Predictor, error)
	// Returns the prediction options
	GetPredictionOptions() (*options.Options, error)
	// Returns the preprocess options
	GetPreprocessOptions() (PreprocessOptions, error)
	// Returns the handle to features
	Predict(ctx context.Context, data interface{}, opts ...options.Option) error
	// Returns the features
	ReadPredictedFeatures(ctx context.Context) ([]dlframework.Features, error)
  // Returns the raw inference result as go tensors
	ReadPredictedFeaturesAsMap(ctx context.Context) (map[string]interface{}, error)
	// Clears the internal state of a predictor
	Reset(ctx context.Context) error

	io.Closer
}

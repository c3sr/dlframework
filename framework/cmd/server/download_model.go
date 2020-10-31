package server

import (
	"context"

	"github.com/pkg/errors"
	"github.com/c3sr/dlframework/framework/agent"
	dlcmd "github.com/c3sr/dlframework/framework/cmd"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

func downloadModels(ctx context.Context) error {
	predictor, err := agent.GetPredictor(framework)
	if err != nil {
		return errors.Wrapf(err,
			"⚠️ failed to get predictor for %s. make sure you have "+
				"imported the framework's predictor package",
			framework.MustCanonicalName(),
		)

	}
	models := framework.Models()
	pb := dlcmd.NewProgress("download models", len(models))
	var g errgroup.Group
	for _, model := range models {
		model := model // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			err := predictor.Download(ctx, model) // Download function is the same for all the predictors
			if err != nil {
				return errors.Wrapf(err, "failed to download %s model", model.MustCanonicalName())
			}
			log.Infof("downloaded model %v", model.MustCanonicalName())
			pb.Increment()
			return nil
		})
	}

	err = g.Wait()
	pb.Finish()

	return err
}

var downloadModelsCmd = &cobra.Command{
	Use:   "models",
	Short: "Download MLModelScope models",
	RunE: func(c *cobra.Command, args []string) error {
		ctx := context.Background()
		return downloadModels(ctx)
	},
}

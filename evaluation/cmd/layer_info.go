package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/c3sr/dlframework/evaluation"
	"github.com/spf13/cobra"
)

var (
	listRuns   bool
	latencyBar bool
	latencyBox bool
	memoryBar  bool
)

var layerInfoCmd = &cobra.Command{
	Use:     "info",
	Aliases: []string{},
	Short:   "Get model layer information from framework traces in a database",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if databaseName == "" {
			databaseName = defaultDatabaseName["layer"]
		}
		err := rootSetup()
		if err != nil {
			return err
		}
		if modelName == "all" && outputFormat == "json" && outputFileName == "" {
			outputFileName = filepath.Join(mlArcWebAssetsPath, "layers")
		}
		if overwrite && isExists(outputFileName) {
			os.RemoveAll(outputFileName)
		}
		if plotPath == "" {
			plotPath = evaluation.TempFile("", "layer_plot_*.html")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		run := func() error {
			evals, err := getEvaluations()
			if err != nil {
				return err
			}

			summary0, err := evals.SummaryLayerInformations(performanceCollection)
			if err != nil {
				return err
			}

			summary1 := evaluation.SummaryLayerLatencyInformations(summary0)
			if sortOutput {
				sort.Slice(summary1, func(ii, jj int) bool {
					return summary1[ii].Duration > summary1[jj].Duration
				})
			}

			if plotAll || latencyBar {
				plotPath = outputFileName + "_latency_bar.html"
				err := summary1.WriteBarPlot(plotPath)
				if err != nil {
					return err
				}
				fmt.Println("Created plot in " + plotPath)
			}
			if plotAll || latencyBox {
				plotPath = outputFileName + "_latency_box.html"
				err = summary1.WriteBoxPlot(plotPath)
				if err != nil {
					return err
				}
				fmt.Println("Created plot in " + plotPath)
			}
			if plotAll || memoryBar {
				plotPath = outputFileName + "_allocated_memory.html"
				summary2 := evaluation.SummaryLayerAllocatedMemoryInformations(summary0)
				err = summary2.WriteBarPlot(plotPath)
				if err != nil {
					return err
				}
				fmt.Println("Created plot in " + plotPath)
			}

			if topLayers != -1 {
				if topLayers >= len(summary0) {
					topLayers = len(summary0)
				}
				summary0 = summary0[:topLayers]
			}

			if listRuns {
				writer := NewWriter(evaluation.SummaryLayerInformation{})
				defer writer.Close()
				for _, v := range summary0 {
					writer.Row(v)
				}
			}

			writer := NewWriter(evaluation.SummaryMeanLayerInformation{})
			defer writer.Close()
			for _, v := range summary0 {
				writer.Row(evaluation.SummaryMeanLayerInformation(v))
			}
			return nil
		}

		return forallmodels(run)
	},
}

func init() {
	layerInfoCmd.PersistentFlags().BoolVar(&listRuns, "list_runs", false, "list evaluations")
	layerInfoCmd.PersistentFlags().BoolVar(&memoryBar, "memory_bar", false, "plot allocated memory as a bar plot")
	layerInfoCmd.PersistentFlags().BoolVar(&latencyBar, "latency_bar", false, "plot latency as a bar plot")
	layerInfoCmd.PersistentFlags().BoolVar(&latencyBox, "latency_box", false, "plot latency as a box plot")
}

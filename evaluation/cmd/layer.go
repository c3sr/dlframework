package cmd

import (
	"github.com/spf13/cobra"
)

var (
	topLayers int
)

var layerCmd = &cobra.Command{
	Use: "layer",
	Aliases: []string{
		"layers",
	},
	Short: "Get evaluation layer analysis from framework traces in a database",
}

func init() {
	layerCmd.PersistentFlags().IntVar(&topLayers, "top_layers", -1, "consider only the top k layers ranked by latency")

	layerCmd.AddCommand(layerInfoCmd)
	layerCmd.AddCommand(layerAggreInfoCmd)
}

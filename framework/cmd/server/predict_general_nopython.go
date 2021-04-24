// +build nopython

package server

import (
	"path/filepath"

	sourcepath "github.com/GeertJohan/go-sourcepath"
	"github.com/k0kubun/pp/v3"
	"github.com/spf13/cobra"
	"github.com/unknwon/com"
)

var (
	inputsFilePath      string
	numInputParts       int
	numWarmUpInputParts int
)

func runPredictGeneralCmd(c *cobra.Command, args []string) error {
	pp.Println("To run predict general, you need to build the agent without nopython flag")
	return nil
}

var predictGeneralCmd = &cobra.Command{
	Use:     "general",
	Short:   "Evaluate using general methods",
	Aliases: []string{"generals"},
	RunE: func(c *cobra.Command, args []string) error {
		if modelName == "all" {
			for _, model := range framework.Models() {
				modelName = model.Name
				modelVersion = model.Version
				runPredictGeneralCmd(c, args)
			}
			return nil
		}
		return runPredictGeneralCmd(c, args)
	},
}

func init() {
	sourcePath := sourcepath.MustAbsoluteDir()
	defaultInputsPath := filepath.Join(sourcePath, "..", "_fixtures", "urlsfile")
	if !com.IsFile(defaultInputsPath) {
		defaultInputsPath = ""
	}
	predictGeneralCmd.PersistentFlags().StringVar(&inputsFilePath, "inputs_file_path", defaultInputsPath, "the path of the file containing the inputs to perform the evaluations on.")
	predictGeneralCmd.PersistentFlags().IntVar(&numInputParts, "num_input_parts", -1, "the number of input parts to process. Setting input parts to a value other than -1 means that only the first num_input_parts * partition_list_size tensors are infered from the dataset. This is useful while performing performance evaluations, where only a few hundred evaluation samples are useful")
	predictGeneralCmd.PersistentFlags().IntVar(&numWarmUpInputParts, "num_warmup_input_parts", 1, "the number of input parts to process during the warmup period. This is ignored if num_file_parts=-1")
}

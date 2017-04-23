package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type mergeOptions struct {
	outputPath string
	inputPaths []string
}

func init() {
	opts := mergeOptions{}

	// mergeCmd represents the merge command
	mergeCmd := &cobra.Command{
		Use:   "merge",
		Short: "Merges one or more PDF files into one",
		Long:  "Merges one or more PDF files into one",
		Run: func(cmd *cobra.Command, args []string) {
			runMerge(opts)
		},
	}

	emptyStringArray := []string{}

	flags := mergeCmd.Flags()
	flags.StringVarP(&opts.outputPath, "output", "o", "", "Output PDF path")
	flags.StringArrayVarP(&opts.inputPaths, "input", "i", emptyStringArray, "Input path(s)")

	RootCmd.AddCommand(mergeCmd)
}

func runMerge(opts mergeOptions) {
	fmt.Println("Merge called", opts.outputPath, opts.inputPaths)
}

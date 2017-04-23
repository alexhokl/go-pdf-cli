package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/unidoc/unidoc/pdf"
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
			if err := runMerge(opts); err != nil {
				fmt.Println(err.Error())
				//fmt.Errorf(err.Error())
				return
			}
		},
	}

	emptyStringArray := []string{}

	flags := mergeCmd.Flags()
	flags.StringVarP(&opts.outputPath, "output", "o", "", "Output PDF path")
	flags.StringArrayVarP(&opts.inputPaths, "input", "i", emptyStringArray, "Input path(s)")

	RootCmd.AddCommand(mergeCmd)
}

func runMerge(opts mergeOptions) error {
	if opts.outputPath == "" {
		return errors.New("Output path was not specified\n")
	}
	if len(opts.inputPaths) == 0 {
		return errors.New("Input path(s) were not specified\n")
	}

	writer := pdf.NewPdfWriter()

	for _, inputPath := range opts.inputPaths {
		fmt.Println(inputPath)
		f, err := os.Open(inputPath)
		if err != nil {
			return err
		}
		defer f.Close()

		pdfReader, err := pdf.NewPdfReader(f)
		if err != nil {
			return err
		}

		isEncrypted, err := pdfReader.IsEncrypted()
		if err != nil {
			return err
		}

		if isEncrypted {
			_, err = pdfReader.Decrypt([]byte(""))
			if err != nil {
				return err
			}
		}

		numPages, err := pdfReader.GetNumPages()
		if err != nil {
			return err
		}

		for i := 0; i < numPages; i++ {
			pageNum := i + 1

			page, err := pdfReader.GetPage(pageNum)
			if err != nil {
				return err
			}

			err = writer.AddPage(page)
			if err != nil {
				return err
			}
		}
	}

	fWrite, err := os.Create(opts.outputPath)
	if err != nil {
		return err
	}
	defer fWrite.Close()

	err = writer.Write(fWrite)
	if err != nil {
		return err
	}

	fmt.Println("Merged ", opts.inputPaths, " into ", opts.outputPath)

	return nil
}

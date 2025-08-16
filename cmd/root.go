package cmd

import (
	"fmt"
	"os"

	"github.com/hawkuash/sumpto/cmd/compress"
	"github.com/hawkuash/sumpto/cmd/convert"
	"github.com/hawkuash/sumpto/cmd/scale"
	"github.com/hawkuash/sumpto/internal/files"
	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "sumpto",
	Short: "Sumpto is an app to mess up ur media collection",
	Long:  `Will write it when i've completed it`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Idk, try harder for once")
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().
		StringVarP(&files.Input, "input", "i", "", "paths to files meant for processing")
	RootCmd.MarkPersistentFlagRequired("input")

	RootCmd.PersistentFlags().
		StringSliceVarP(&files.FormatList, "format", "f", nil, "filters search results by presented file formats")

	RootCmd.PersistentFlags().
		IntVarP(&files.Quality, "quality", "q", 100, "quality param for supported formats")

	RootCmd.PersistentFlags().
		BoolVarP(&files.Recursive, "recursive", "r", true, "recursive flag indicates if search in subdirectories must be done")

	RootCmd.AddCommand(scale.ScaleCmd)
	RootCmd.AddCommand(convert.ConvertCmd)
	RootCmd.AddCommand(compress.CompressCmd)
}

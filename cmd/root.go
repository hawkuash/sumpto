package cmd

import (
	"fmt"
	"os"

	"github.com/hawkuash/sumpto/cmd/convert"
	"github.com/hawkuash/sumpto/cmd/scale"
	"github.com/hawkuash/sumpto/files"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
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
	RootCmd.PersistentFlags().StringVarP(&files.Paths, "paths", "p", "", "paths to files meant for processing")
	RootCmd.MarkPersistentFlagRequired("paths")

	RootCmd.PersistentFlags().BoolVarP(&files.Recursive, "recursive", "r", true, "recursive flag indicates if search in subdirectories must be done")
	RootCmd.AddCommand(scale.ScaleCmd)
	RootCmd.AddCommand(convert.ConvertCmd)
}

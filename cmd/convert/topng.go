package convert

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/convert"
	"github.com/hawkuash/sumpto/files"
	"github.com/spf13/cobra"
)

// topngCmd represents the topng command
var TopngCmd = &cobra.Command{
	Use:   "to-png",
	Short: "Converts supported formats to PNG",
	Long:  `Converts supported formats to PNG`,
	Run: func(cmd *cobra.Command, args []string) {
		vips.Startup(nil)
		defer vips.Shutdown()
		for _, file := range files.GenerateFiles(files.Input, files.Recursive, convert.SetPNGConvertExtensions(files.Format_list)) {
			convert.ConvertToPNG(file)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topngCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topngCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		vips.Startup(nil)
		defer vips.Shutdown()
		for _, file := range files.GenerateFiles(files.Input, files.Recursive, convert.Supported_formats_png) {
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

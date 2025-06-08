package convert

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/convert"
	"github.com/hawkuash/sumpto/files"
	"github.com/spf13/cobra"
)

// tojpegCmd represents the tojpeg command
var TojpegCmd = &cobra.Command{
	Use:   "to-jpeg",
	Short: "Converts supported formats to JPEG",
	Long:  `Converts supported formats to JPEG`,
	Run: func(cmd *cobra.Command, args []string) {
		vips.Startup(nil)
		defer vips.Shutdown()
		for _, file := range files.GenerateFiles(files.Input, files.Recursive, convert.Supported_formats_jpeg) {
			convert.ConvertToJPEG(file)
		}
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// tojpegCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// tojpegCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

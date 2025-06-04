package scale

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/files"
	"github.com/hawkuash/sumpto/scale"
	"github.com/spf13/cobra"
)

var (
	overwrite   bool
	format_list []string
	limit       int
)

// ScaleCmd represents the scale command
var ScaleCmd = &cobra.Command{
	Use:   "scale",
	Short: " ",
	Long:  `scale command`,
	Run: func(scmd *cobra.Command, args []string) {
		vips.Startup(nil)
		defer vips.Shutdown()
		for _, file := range files.GenerateFiles(files.Input, files.Recursive, scale.SetScaleExtensions(format_list)) {
			scale.ScaleImage(file, overwrite, limit)
		}
	},
}

func init() {
	ScaleCmd.Flags().StringSliceVarP(&format_list, "format", "f", nil, "filters search results by presented file formats")
	ScaleCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "declares if file should be overwritten")
	ScaleCmd.Flags().IntVarP(&limit, "dimension-limit", "l", 2160, "lower limit of image dimensions for downscaling and upper - for upscaling, but no upscaling for now")
}

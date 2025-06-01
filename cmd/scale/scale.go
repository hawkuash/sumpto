package scale

import (
	"fmt"

	"github.com/hawkuash/sumpto/files"
	"github.com/hawkuash/sumpto/scale"
	"github.com/spf13/cobra"
)

var search_options scale.Search_options

var overwrite bool

// ScaleCmd represents the scale command
var ScaleCmd = &cobra.Command{
	Use:   "scale",
	Short: " ",
	Long:  `scale command`,
	Run: func(scmd *cobra.Command, args []string) {
		for _, file := range files.GenerateFiles(files.Paths, files.Recursive, scale.SetScaleExtensions(search_options)) {
			fmt.Println(file)
		}
	},
}

func init() {
	ScaleCmd.Flags().BoolVar(&search_options.Jpeg, "jpeg", false, "sets JPEG format for search")
	ScaleCmd.Flags().BoolVar(&search_options.Png, "png", false, "sets PNG format for search")
	ScaleCmd.Flags().BoolVar(&search_options.All, "all", false, "sets all fupported formats for search")
	ScaleCmd.MarkFlagsOneRequired("jpeg", "png", "all")

	ScaleCmd.Flags().BoolVarP(&overwrite, "overwrite", "o", false, "declares if file should be overwritten")
}

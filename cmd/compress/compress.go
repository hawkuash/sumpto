package compress

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/hawkuash/sumpto/compress"
	"github.com/hawkuash/sumpto/files"
	"github.com/spf13/cobra"
)

var (
	overwrite bool
)

var CompressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compresses supported formats",
	Long:  `Compresses supported formats using quality parameter`,
	Run: func(cmd *cobra.Command, args []string) {
		vips.Startup(nil)
		defer vips.Shutdown()
		for _, file := range files.GenerateFiles(files.Input, files.Recursive, compress.SetCompressExtensions(files.FormatList)) {
			compress.CompressImage(file, overwrite)
		}
	},
}

func init() {
	CompressCmd.Flags().
		BoolVarP(&overwrite, "overwrite", "o", false, "declares if file should be overwritten")
}

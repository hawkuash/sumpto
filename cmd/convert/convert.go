package convert

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ConvertCmd represents the convert command
var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A parent command for all conversion",
	Long:  `A parent command for all conversion`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(
			"Good you called it. To convert image, call corresponding subcommand",
		)
	},
}

func init() {
	ConvertCmd.AddCommand(ToJPEGCmd)
	ConvertCmd.AddCommand(ToPNGCmd)
}

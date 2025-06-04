package convert

import (
	"fmt"

	"github.com/spf13/cobra"
)

// convertCmd represents the convert command
var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "A parent command for conversion",
	Long:  `Wait for an update, sry`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Good you called it. To convert image, call corresponding subcommand")
	},
}

func init() {
	ConvertCmd.AddCommand(TojpegCmd)
	ConvertCmd.AddCommand(TopngCmd)
}

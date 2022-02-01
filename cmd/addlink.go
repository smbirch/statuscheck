package cmd

import (
	"fmt"

	"github.com/smbirch/statuscheck/links"
	"github.com/spf13/cobra"
)

// addlinkCmd represents the addlink command
var addlinkCmd = &cobra.Command{
	Use:   "addlink",
	Short: "Adds a link to the list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("addlink called")
		links.AddLink()

	},
}

func init() {
	rootCmd.AddCommand(addlinkCmd)

}

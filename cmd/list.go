package cmd

import (
	"fmt"

	"github.com/smbirch/statuscheck/links"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists out the URLs to check the status of",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("These are the current URLs that will be checked:")
		list := links.ReadList()
		for _, l := range list {
			fmt.Println(l)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

}

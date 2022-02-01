package cmd

import (
	"fmt"

	"github.com/smbirch/statuscheck/links"
	"github.com/spf13/cobra"
)

// loopCmd represents the loop command
var loopCmd = &cobra.Command{
	Use:   "loop",
	Short: "Will continuosly check sites for status",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Looping...")
		links.LoopLinks()
	},
}

func init() {
	rootCmd.AddCommand(loopCmd)
}

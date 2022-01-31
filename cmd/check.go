package cmd

import (
	"fmt"

	"github.com/smbirch/statuscheck/links"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks the uptime of websites from a list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")

		list := links.GetLinks()

		c := make(chan string)

		for _, l := range list {
			go links.CheckLink(l, c)
		}

		for i := 0; i < len(list); i++ {
			fmt.Println(<-c)
		}

	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

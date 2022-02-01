package cmd

import (
	"fmt"
	"time"

	"github.com/smbirch/statuscheck/links"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks the uptime of websites from a list",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nPlease wait: Checking status of websites") //<-Just goofing, these sleep calls have no bearing on the logic at all
		time.Sleep(300 * time.Millisecond)
		fmt.Println(".")
		time.Sleep(700 * time.Millisecond)
		fmt.Println(".")
		time.Sleep(700 * time.Millisecond)
		fmt.Println(".")
		time.Sleep(300 * time.Millisecond)

		list := links.ReadList()
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
}

/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"time"

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

		list := links.GetLinks()

		c := make(chan string)

		for _, l := range list {
			go links.CheckLink(l, c)
		}

		for l := range c {
			go func(link string) {
				time.Sleep(5 * time.Second)
				links.CheckLink(link, c)
			}(l)
		}
	},
}

func init() {
	rootCmd.AddCommand(loopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loopCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

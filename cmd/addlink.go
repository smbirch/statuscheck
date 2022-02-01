/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
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
		// fmt.Println("Addning %s to list", )
		links.AddLink()

	},
}

func init() {
	rootCmd.AddCommand(addlinkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addlinkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addlinkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

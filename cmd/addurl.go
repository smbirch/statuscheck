/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// addurlCmd represents the addurl command
var addurlCmd = &cobra.Command{
	Use:   "addurl",
	Short: "Add your own URL to be checked",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("\nEnter the URL you would like to add: ")
		fmt.Println("(Format: http://sitename.com)")
		fmt.Println("")
		//TODO:
		//make addlink function to add links
		var input string
		fmt.Scanln(&input)
		if len(input) == 0 {
			fmt.Println("You entered: \n...")
			fmt.Println("Please provide at least one URL")
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(addurlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addurlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addurlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

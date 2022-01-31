/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/

package main

import (
	"time"

	"github.com/smbirch/statuscheck/cmd"
)

func main() {
	// links := []string{
	// 	"http://google.com",
	// 	"http://facebook.com",
	// 	"http://twitter.com",
	// 	"http://golang.org",
	// 	"http://github.com",
	// 	"http://amazon.com",
	// 	"http://stackoverflow.com",
	// 	"http://apple.com",
	// 	"http://netflix.com",
	// }

	cmd.Execute()

	links := GetLinks()

	c := make(chan string)

	for _, l := range links {
		go CheckLink(l, c)

	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			CheckLink(link, c)
		}(l)
	}
}

// cli app
// add option for flags to check any website, or ability to check it while app is open
// request this list of sites and present it on opening app

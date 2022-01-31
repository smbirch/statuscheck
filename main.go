package main

import (
	"fmt"

	"github.com/smbirch/statuscheck/cmd"
	"github.com/smbirch/statuscheck/links"
)

func main() {

	cmd.Execute()

	list := links.GetLinks()

	c := make(chan string)

	for _, l := range list {
		go links.CheckLink(l, c)
	}

	for i := 0; i < len(list); i++ {
		fmt.Println(<-c)
	}
}

// cli app
// add option for flags to check any website, or ability to check it while app is open
// request this list of sites and present it on opening app

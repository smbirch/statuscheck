package main

import (
	"github.com/smbirch/statuscheck/cmd"
	"github.com/smbirch/statuscheck/links"
)

func main() {

	cmd.Execute()
	links.BuildList()
}

package main

import (
	"github.com/smbirch/statuscheck/cmd"
	"github.com/smbirch/statuscheck/links"
)

func main() {
	links.BuildList()
	cmd.Execute()

}

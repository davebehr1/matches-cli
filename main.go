package main

import (
	"github.com/davebehr1/spanassessment/cmd"
)

func main() {
	root := cmd.Initialize()
	root.Execute()
}

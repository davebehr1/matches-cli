package main

import (
	"github.com/davebehr1/spanassessment/cmd"
	"github.com/davebehr1/spanassessment/pkg/scanmatches"
)

func main() {
	rankTable := scanmatches.NewRankTable()
	root := cmd.Initialize(&rankTable)
	root.Execute()
}

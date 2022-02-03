package main

import (
	"github.com/davebehr1/spanassessment/cmd"
	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	"github.com/davebehr1/spanassessment/pkg/writeranktable"
)

func main() {
	scanner := scanmatches.NewRankTable()
	writer := writeranktable.NewWriter()
	root := cmd.Initialize(&scanner, &writer)
	err := root.Execute()
	if err != nil {
		panic(err)
	}
}

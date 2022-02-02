package main

import (
	"github.com/davebehr1/spanassessment/cmd"
	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	"github.com/davebehr1/spanassessment/pkg/writematches"
)

func main() {
	rankTable := scanmatches.NewRankTable()
	writer := writematches.NewWriter()
	root := cmd.Initialize(&rankTable, &writer)
	root.Execute()
}

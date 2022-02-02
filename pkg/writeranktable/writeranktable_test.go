package writeranktable_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	"github.com/davebehr1/spanassessment/pkg/writeranktable"
	. "github.com/onsi/gomega"
)

func TestWriteRanktable(t *testing.T) {
	g := NewGomegaWithT(t)
	outputFilepath := "../ranktable_test.txt"

	writer := writeranktable.NewWriter()
	rankTable := scanmatches.NewRankTable()
	matches, err := rankTable.ScanFromFile("../../matches.txt")
	g.Expect(err).To(BeNil())

	var finalRankTable string
	for index, team := range matches {
		finalRankTable += fmt.Sprintf("%v. %s, %d pts \n", index+1, team.Team, team.Rank)
	}

	err = writer.WriteToFile(outputFilepath, finalRankTable)
	g.Expect(err).To(BeNil())

	file, err := os.Open(outputFilepath)
	g.Expect(err).To(BeNil())

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	file.Close()

	g.Expect(len(text)).To(Equal(7))

	t.Cleanup(func() {
		err := os.Remove(outputFilepath)
		g.Expect(err).To(BeNil())
	})
}

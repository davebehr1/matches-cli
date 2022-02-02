package scanmatches_test

import (
	"bytes"
	"testing"

	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	. "github.com/onsi/gomega"
)

func TestScanMatches(t *testing.T) {
	t.Run("Scan matches from file", func(t *testing.T) {
		g := NewGomegaWithT(t)
		rankTable := scanmatches.NewRankTable()
		matches, err := rankTable.ReadFromFile("../../matches.txt")
		g.Expect(err).To(BeNil())
		g.Expect(len(matches)).To(Equal(7))

	})
	t.Run("Scan matches from stdin", func(t *testing.T) {
		g := NewGomegaWithT(t)
		rankTable := scanmatches.NewRankTable()

		match, err := rankTable.ReadFromStdin(bytes.NewBufferString("Lions 3, Cheetas 0 \n"))
		g.Expect(err).To(BeNil())
		g.Expect(match).To(Equal("Lions 3, Cheetas 0 \n"))
	})
}

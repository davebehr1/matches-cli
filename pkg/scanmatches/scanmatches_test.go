package scanmatches_test

import (
	"bytes"
	"testing"

	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

func TestScanMatches(t *testing.T) {
	t.Run("Scan matches from file", func(t *testing.T) {
		g := NewGomegaWithT(t)
		rankTable := scanmatches.NewRankTable()
		matches, err := rankTable.ScanFromFile("../../matches.txt")
		g.Expect(err).To(BeNil())
		g.Expect(len(matches)).To(Equal(7))

		g.Expect(matches).To(ContainElements(
			PointTo(MatchFields(IgnoreExtras, Fields{
				"Rank": Equal(11), "Team": Equal("Lions"),
			})),
			// PointTo(MatchAllFields(Fields{
			// 	"Team": Equal("Tarantulas"), "Rank": Equal(6),
			// })),
			// PointTo(MatchAllFields(Fields{
			// 	"Team": Equal("FC Awesome"), "Rank": Equal(1),
			// })),
			// PointTo(MatchAllFields(Fields{
			// 	"Team": Equal("Snakes"), "Rank": Equal(1),
			// })),
			// PointTo(MatchAllFields(Fields{
			// 	"Team": Equal("Grouches"), "Rank": Equal(0),
			// })),
			// PointTo(MatchAllFields(Fields{
			// 	"Team": Equal("Pumas"), "Rank": Equal(0),
			// })),
			// PointTo(MatchAllFields(Fields{
			// 	"Team": Equal("Cheetas"), "Rank": Equal(0),
			// })),
		))

	})
	t.Run("Scan matches from stdin", func(t *testing.T) {
		g := NewGomegaWithT(t)
		rankTable := scanmatches.NewRankTable()

		match, err := rankTable.ScanFromStdin(bytes.NewBufferString("Lions 3, Cheetas 0 \n"))
		g.Expect(err).To(BeNil())
		g.Expect(match).To(Equal("Lions 3, Cheetas 0 \n"))
	})
}

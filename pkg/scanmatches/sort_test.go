package scanmatches_test

import (
	"testing"

	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	. "github.com/onsi/gomega"
)

func TestSort(t *testing.T) {
	g := NewGomegaWithT(t)

	rankTable := scanmatches.NewRankTable()
	rankTable.Table["Lions"] = 6
	rankTable.Table["Snakes"] = 1
	rankTable.Table["Cheetahs"] = 11
	rankTable.Table["Pumas"] = 0
	rankTable.Table["Grouches"] = 0

	rank := rankTable.Sort()
	g.Expect(rank).To(Equal([]scanmatches.TeamRank{{
		Team: "Cheetahs",
		Rank: 11,
	}, {
		Team: "Lions",
		Rank: 6,
	}, {
		Team: "Snakes",
		Rank: 1,
	}, {
		Team: "Pumas",
		Rank: 0,
	}, {
		Team: "Grouches",
		Rank: 0,
	}}))

}

package scanmatches_test

import (
	"testing"

	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	. "github.com/onsi/gomega"
)

func TestProcessMatch(t *testing.T) {
	type testCase struct {
		input     string
		finalRank map[string]int
	}

	for _, test := range []testCase{
		{
			input: "Lions 3, Snakes 3",
			finalRank: map[string]int{
				"Lions":  1,
				"Snakes": 1,
			},
		},
		{
			input: "Lions 1, Snakes 3",
			finalRank: map[string]int{
				"Lions":  0,
				"Snakes": 3,
			},
		},
		{
			input: "Lions 3, Snakes 1",
			finalRank: map[string]int{
				"Lions":  3,
				"Snakes": 0,
			},
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			g := NewGomegaWithT(t)
			table := scanmatches.RankTable{Table: make(map[string]int)}
			err := table.Process(test.input)
			g.Expect(err).To(BeNil())
			g.Expect(len(table.Table)).To(Equal(2))
			g.Expect(table.Table).To(Equal(test.finalRank))
		})
	}

}

package scanmatches_test

import (
	"testing"

	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	. "github.com/onsi/gomega"
)

func TestProcessMatch(t *testing.T) {
	type testCase struct {
		input       string
		winningTeam string
	}

	for _, test := range []testCase{
		{
			input:       "Lions 3, Snakes 3",
			winningTeam: "draw",
		},
		{
			input:       "Lions 1, Snakes 3",
			winningTeam: "Snakes",
		},
		{
			input:       "Lions 3, Snakes 1",
			winningTeam: "Lions",
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			g := NewGomegaWithT(t)
			table := scanmatches.RankTable{Table: make(map[string]int)}
			err := table.Process(test.input)
			g.Expect(err).To(BeNil())
			g.Expect(len(table.Table)).To(Equal(2))
		})
	}

}

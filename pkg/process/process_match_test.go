package process_test

import (
	"testing"

	"github.com/davebehr1/spanassessment/pkg/process"
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
			result, err := process.Process(test.input)
			g.Expect(err).To(BeNil())
			g.Expect(result).To(Equal(test.winningTeam))
		})
	}

}

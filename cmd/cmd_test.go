package cmd_test

import (
	"bytes"
	"testing"

	"github.com/davebehr1/spanassessment/cmd"
	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	. "github.com/onsi/gomega"
)

func TestCmd(t *testing.T) {
	t.Run("Run the generate command without flags", func(t *testing.T) {
		g := NewGomegaWithT(t)

		bf := new(bytes.Buffer)

		rankTable := scanmatches.NewRankTable()

		RootCommand := cmd.Initialize(&rankTable)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(Equal("matches from stdin"))
	})
	t.Run("Run the generate command with flags", func(t *testing.T) {
		g := NewGomegaWithT(t)

		bf := new(bytes.Buffer)
		rankTable := scanmatches.NewRankTable()

		RootCommand := cmd.Initialize(&rankTable)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt", "--f=matches.txt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(Equal("matches.txt"))
	})

	t.Run("Run the generate command with incorrect flags", func(t *testing.T) {
		g := NewGomegaWithT(t)

		bf := new(bytes.Buffer)

		rankTable := scanmatches.NewRankTable()

		RootCommand := cmd.Initialize(&rankTable)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt", "--d=matches.txt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(ContainSubstring("unknown flag"))
	})

}

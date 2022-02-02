package cmd_test

import (
	"bytes"
	"testing"

	"github.com/davebehr1/spanassessment/cmd"
	"github.com/davebehr1/spanassessment/pkg/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
)

func TestCmd(t *testing.T) {
	t.Run("Run the generate command without flags", func(t *testing.T) {
		g := NewGomegaWithT(t)
		ctrl := gomock.NewController(t)
		ScanMatches := mocks.NewMockScanMatches(ctrl)

		ScanMatches.EXPECT().ScanFromStdin(gomock.Any()).Return("Lions 3, Snakes 3", nil)

		bf := new(bytes.Buffer)
		RootCommand := cmd.Initialize(ScanMatches)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(Equal("Lions 3, Snakes 3"))
	})
	t.Run("Run the generate command with flags", func(t *testing.T) {
		g := NewGomegaWithT(t)
		ctrl := gomock.NewController(t)
		ScanMatches := mocks.NewMockScanMatches(ctrl)

		ScanMatches.EXPECT().ScanFromFile(gomock.Any()).Return([]string{"Lions 3, Snakes 3"}, nil)

		bf := new(bytes.Buffer)
		RootCommand := cmd.Initialize(ScanMatches)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt", "--f=matches.txt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(Equal("Lions 3, Snakes 3"))
	})

	t.Run("Run the generate command with incorrect flags", func(t *testing.T) {
		g := NewGomegaWithT(t)

		ctrl := gomock.NewController(t)
		ScanMatches := mocks.NewMockScanMatches(ctrl)

		bf := new(bytes.Buffer)
		RootCommand := cmd.Initialize(ScanMatches)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt", "--d=matches.txt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(ContainSubstring("unknown flag"))
	})

}

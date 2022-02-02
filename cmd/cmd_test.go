package cmd_test

import (
	"bytes"
	"testing"

	"github.com/davebehr1/spanassessment/cmd"
	"github.com/davebehr1/spanassessment/pkg/mocks"
	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/gomega"
)

func TestCmd(t *testing.T) {
	mocks := initMocks(t)

	t.Run("Run the generate command without flags", func(t *testing.T) {
		g := NewGomegaWithT(t)

		mocks.scanner.EXPECT().ScanFromStdin(gomock.Any()).Return([]*scanmatches.TeamRank{
			{
				Team: "Lions",
				Rank: 3,
			}, {
				Team: "Snakes",
				Rank: 3,
			}}, nil)

		bf := new(bytes.Buffer)
		RootCommand := cmd.Initialize(mocks.scanner, mocks.writer)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(Equal("1. Lions, 3 pts \n2. Snakes, 3 pts \n"))
	})
	t.Run("Run the generate command with file flag", func(t *testing.T) {
		g := NewGomegaWithT(t)

		mocks.scanner.EXPECT().ScanFromFile(gomock.Any()).Return([]*scanmatches.TeamRank{
			{
				Team: "Lions",
				Rank: 3,
			}, {
				Team: "Snakes",
				Rank: 3,
			}}, nil)

		bf := new(bytes.Buffer)
		RootCommand := cmd.Initialize(mocks.scanner, mocks.writer)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt", "--f=matches.txt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(Equal("1. Lions, 3 pts \n2. Snakes, 3 pts \n"))
	})

	t.Run("Run the generate command with incorrect flags", func(t *testing.T) {
		g := NewGomegaWithT(t)

		bf := new(bytes.Buffer)
		RootCommand := cmd.Initialize(mocks.scanner, mocks.writer)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt", "--d=matches.txt"})
		RootCommand.Execute()

		result := bf.String()
		g.Expect(result).To(ContainSubstring("unknown flag"))
	})

	t.Run("Run the generate command with file and output flag", func(t *testing.T) {
		g := NewGomegaWithT(t)

		mocks.scanner.EXPECT().ScanFromFile(gomock.Any()).Return([]*scanmatches.TeamRank{
			{
				Team: "Lions",
				Rank: 3,
			}, {
				Team: "Snakes",
				Rank: 3,
			}}, nil)

		mocks.writer.EXPECT().WriteToFile(gomock.Any(), gomock.Any()).Return(nil)

		RootCommand := cmd.Initialize(mocks.scanner, mocks.writer)

		bf := new(bytes.Buffer)
		RootCommand.SetOut(bf)
		RootCommand.SetErr(bf)
		RootCommand.SetArgs([]string{"grt", "--f=matches.txt", "--o=rankTable.txt"})
		err := RootCommand.Execute()
		g.Expect(err).To(BeNil())
	})

}

type Mocks struct {
	scanner *mocks.MockScanMatches
	writer  *mocks.MockWriteRankTable
}

func initMocks(t *testing.T) *Mocks {
	ctrl := gomock.NewController(t)
	ScanMatches := mocks.NewMockScanMatches(ctrl)
	WriteRankTable := mocks.NewMockWriteRankTable(ctrl)

	return &Mocks{
		scanner: ScanMatches,
		writer:  WriteRankTable,
	}
}

package cmd

import (
	"bufio"
	"fmt"

	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	"github.com/davebehr1/spanassessment/pkg/writeranktable"
	"github.com/spf13/cobra"
)

func NewGenerateRankTableCmd(scan scanmatches.ScanMatches, write writeranktable.WriteRankTable) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generateranktable",
		Aliases: []string{"grt"},
		Short:   "ingest match results and generate a rank table",
		Long:    "ingest match results via a text file or stdin and generate a rank table",
		RunE: func(cmd *cobra.Command, args []string) error {
			var matches []*scanmatches.TeamRank
			var err error
			if cmd.Flags().Changed("f") {
				var filePath string
				filePath, err := cmd.Flags().GetString("f")
				if err != nil {
					return err
				}
				matches, err = scan.ScanFromFile(filePath)
				if err != nil {
					return err
				}
			} else {
				matches, err = scan.ScanFromStdin(bufio.NewReader(cmd.InOrStdin()))
				if err != nil {
					return err
				}
			}
			var finalRankTable string
			for index, team := range matches {
				finalRankTable += fmt.Sprintf("%v. %s, %d pts \n", index+1, team.Team, team.Rank)
			}

			if cmd.Flags().Changed("o") {
				outputFilePath, err := cmd.Flags().GetString("o")
				if err != nil {
					return err
				}
				err = write.WriteToFile(outputFilePath, finalRankTable)
				if err != nil {
					return err
				}
			} else {

				_, err = fmt.Fprint(cmd.OutOrStdout(), finalRankTable)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().String("f", "", "matches.txt")
	cmd.Flags().String("o", "", "rankTable.txt")
	return cmd
}

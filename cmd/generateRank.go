package cmd

import (
	"bufio"
	"fmt"
	"strings"

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

			finalRankTable := []string{}
			for index, team := range matches {
				finalRankTable = append(finalRankTable, fmt.Sprintf("%v. %s, %d pts", index+1, team.Team, team.Rank))
			}

			if cmd.Flags().Changed("o") {
				outputFilePath, err := cmd.Flags().GetString("o")
				if err != nil {
					return err
				}
				err = write.WriteToFile(outputFilePath, strings.Join(finalRankTable, "\n"))
				if err != nil {
					return err
				}
			} else {

				_, err = fmt.Fprint(cmd.OutOrStdout(), strings.Join(finalRankTable, "\n"))
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

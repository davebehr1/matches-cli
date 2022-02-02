package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davebehr1/spanassessment/pkg/scanmatches"
	"github.com/spf13/cobra"
)

func NewGenerateRankTableCmd(scan scanmatches.ScanMatches) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "generateranktable",
		Aliases: []string{"grt"},
		Short:   "ingest match results and generate a rank table",
		Long:    "ingest match results via a text file or stdin and generate a rank table",
		RunE: func(cmd *cobra.Command, args []string) error {

			if cmd.Flags().Changed("f") {
				var filePath string
				filePath, err := cmd.Flags().GetString("f")
				if err != nil {
					return err
				}

				matches, err := scan.ScanFromFile(filePath)
				if err != nil {
					return err
				}
				var finalRankTable string
				for index, team := range matches {
					finalRankTable += fmt.Sprintf("%v. %s, %d pts \n", index+1, team.Team, team.Rank)
				}

				_, err = fmt.Fprint(cmd.OutOrStdout(), finalRankTable)
				if err != nil {
					return err
				}
			} else {
				matches, err := scan.ScanFromStdin(bufio.NewReader(os.Stdin))
				if err != nil {
					return err
				}
				var finalRankTable string
				for index, team := range matches {
					finalRankTable += fmt.Sprintf("%v. %s, %d pts \n", index+1, team.Team, team.Rank)
				}

				_, err = fmt.Fprint(cmd.OutOrStdout(), finalRankTable)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().String("f", "", "matches.txt")
	return cmd
}

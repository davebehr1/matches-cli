package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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

				_, err = fmt.Fprint(cmd.OutOrStdout(), strings.Join(matches, ","))
				if err != nil {
					return err
				}
			} else {
				match, err := scan.ScanFromStdin(bufio.NewReader(os.Stdin))
				if err != nil {
					return err
				}

				_, err = fmt.Fprint(cmd.OutOrStdout(), match)
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	cmd.Flags().String("f", "", "macthes.txt")
	return cmd
}

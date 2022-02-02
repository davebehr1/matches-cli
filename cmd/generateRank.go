package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewGenerateRankTableCmd() *cobra.Command {
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
				fmt.Fprintf(cmd.OutOrStdout(), filePath)
			} else {
				var result string
				result = "matches from stdin"
				fmt.Fprintf(cmd.OutOrStdout(), result)
			}
			return nil
		},
	}

	cmd.Flags().String("f", "", "macthes.txt")
	return cmd
}

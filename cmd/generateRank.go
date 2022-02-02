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
				_, err = fmt.Fprint(cmd.OutOrStdout(), filePath)
				if err != nil {
					return err
				}
			} else {
				result := "matches from stdin"
				_, err := fmt.Fprint(cmd.OutOrStdout(), result)
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

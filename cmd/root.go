package cmd

import (
	"github.com/spf13/cobra"
)

var RootCommand *cobra.Command

func NewRootCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "root",
		Short: "This is the Root Command",
	}
	return cmd
}

func Execute() {
	cobra.CheckErr(RootCommand.Execute())
}

func Initialize() *cobra.Command {
	RootCommand := NewRootCmd()

	RootCommand.AddCommand(NewGenerateRankTableCmd())

	return RootCommand
}

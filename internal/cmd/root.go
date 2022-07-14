package cmd

import (
	v "github.com/samuelmachado/cli-poc-dock/internal/cmd/version"
	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

//Root Command
func Root(vf v.FullVersion) {

	var flags cmd.Flags

	rootCmd = &cobra.Command{
		Use:   "dock",
		Short: "dock",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	rootCmd.PersistentFlags().StringVarP(
		&flags.Dir, "dir", "d", "./",
		"The output dir: -d ~/Docs",
	)

	rootCmd.PersistentFlags().StringVarP(
		&flags.FormatType, "format-type", "t", "test",
		"The output format: -t test",
	)

	rootCmd.AddCommand(version(vf))

	rootCmd.Execute()

}

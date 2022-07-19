package cmd

import (
	"github.com/samuelmachado/cli-poc-dock/internal/cmd/sources"
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
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	// Global flags
	// global flags should work for all commands
	rootCmd.PersistentFlags().StringVarP(
		&flags.FormatType, "output type", "o", "json",
		"The output format: -o json",
	)
	//

	// Global commands
	// These are generic commands available in the CLI
	rootCmd.AddCommand(version(vf))

	// Custom commands
	// any command intended to perform an action in DOCK applications
	addCustomCommands(rootCmd, &flags)

	rootCmd.Execute()

}

func addCustomCommands(rootCmd *cobra.Command, flags *cmd.Flags) {
	rootCmd.AddCommand(sources.Caradhras(flags))
	//
	//
	//
	//
}

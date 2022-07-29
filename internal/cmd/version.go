package cmd

import (
	"fmt"

	v "github.com/samuelmachado/cli-poc-dock/pkg/structs/version"
	"github.com/spf13/cobra"
)

//version of the binary built
func version(vf v.FullVersion) (versionCmd *cobra.Command) {

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of CLI DOCK",
		Long:  `Print the semantical version of CLI DOCK built`,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Fprintf(versionCmd.OutOrStdout(), "version: %s\nbuilded at: %s\ncommit hash: %s\n", vf.Version, vf.Date, vf.Commit)
		},
	}

	return versionCmd
}

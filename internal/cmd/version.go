package cmd

import (
	"fmt"

	v "github.com/samuelmachado/cli-poc-dock/internal/cmd/version"
	"github.com/spf13/cobra"
)

//version of the binary built
func version(vf v.FullVersion) (versionCmd *cobra.Command) {

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version of CLI DOCK",
		Long:  `Print the semantical version of CLI DOCK built`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("version: %s\nbuilded at: %s\ncommit hash: %s\n", vf.Version, vf.Date, vf.Commit)
		},
	}

	return versionCmd
}

package sources

import (
	"fmt"

	"github.com/samuelmachado/cli-poc-dock/internal/domains/caradhas"
	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"

	"github.com/spf13/cobra"
)

//Caradhras is the command that handle this source
func Caradhras(flags *cmd.Flags) (caradhrasCmd *cobra.Command) {

	caradhrasCmd = &cobra.Command{
		Use:   "caradhas",
		Short: "caradhas source",
		Long:  `caradhas source handler`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	caradhrasCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Get info about the version of Caradhas API",
		Long:  `Get information about the version of Caradhas API`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(caradhas.Info())
		},
	})

	return caradhrasCmd
}

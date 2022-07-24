package sources

import (
	"github.com/samuelmachado/cli-poc-dock/internal/domains/profile"
	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
	"github.com/spf13/cobra"
)

func Profile(_ *cmd.Flags) (caradhrasCmd *cobra.Command) {
	manager := profile.NewProfileManager()
	caradhrasCmd = &cobra.Command{
		Use:   "profile",
		Short: "use to manage your profile",
		Long:  "",
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	caradhrasCmd.AddCommand(&cobra.Command{
		Use:   "create",
		Short: "create a new profile",
		Long:  "",
		Run: func(_ *cobra.Command, _ []string) {
			manager.Create()
		},
	})

	caradhrasCmd.AddCommand(&cobra.Command{
		Use:   "select",
		Short: "Get info about the version of Caradhras SDK",
		Long:  `Get information about the version of Caradhras SDK`,
		Run: func(_ *cobra.Command, _ []string) {

		},
	})

	return caradhrasCmd
}

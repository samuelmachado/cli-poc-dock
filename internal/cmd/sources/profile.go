package sources

import (
	"fmt"

	"github.com/samuelmachado/cli-poc-dock/internal/domains/profile"
	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
	"github.com/spf13/cobra"
)

func Profile(_ *cmd.Flags) (caradhrasCmd *cobra.Command) {
	manager := profile.NewProfileManager()

	caradhrasCmd = &cobra.Command{
		Use:   "profile",
		Short: "used to manage a profile",
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
		Short: "select a specific profile",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			profileName := args[0]
			manager.Select(profileName)

			fmt.Fprintf(cmd.OutOrStdout(), "%s selected", profileName)
		},
	})

	caradhrasCmd.AddCommand(&cobra.Command{
		Use:   "delete",
		Short: "remove selected profile",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			profileName := args[0]
			manager.Delete(profileName)

			fmt.Fprintf(cmd.OutOrStdout(), "%s deleted", profileName)
		},
	})

	return caradhrasCmd
}

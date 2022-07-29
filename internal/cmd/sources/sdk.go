package sources

import (
	"fmt"

	"github.com/samuelmachado/cli-poc-dock/internal/domains/sdk"
	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"

	"github.com/spf13/cobra"
)

//Sdk is the command that handle this source
func Sdk(flags *cmd.Flags) (sdkCmd *cobra.Command) {
	manager := sdk.NewSdkManager()

	sdkCmd = &cobra.Command{
		Use:   "sdk",
		Short: "use to access the dock SDK",
		Long:  `command used to access SDK methods. You can also check the official documentation by going to https://lighthouse.dockm.tech/`,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}
	// flags
	sdkCmd.PersistentFlags().StringVarP(
		&flags.Token, "token", "t", "",
		fmt.Sprintf("(Optional) allows passing an access_token to be used"),
	)
	//

	// commands

	sdkCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Get info about the version of SDK",
		Long:  `Get information about the version of SDK`,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(manager.Version())
		},
	})
	//

	return sdkCmd
}

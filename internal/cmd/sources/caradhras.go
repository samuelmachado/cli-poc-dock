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
		Short: "use to access the caradhas SDK",
		Long:  `command used to access caradhas SDK methods. You can also check the official documentation by going to https://lighthouse.dock.tech/`,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	caradhrasCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "init the caradhas SDK",
		Long:  `Initializes the caradhas SDK. The arg 0 is used for the token. e.g.: caradhas init eyJra...bEw`,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return caradhas.ErrTokenInvalid
			}
			return nil

			//TO DO:
			// pass this validation to domain?
		},
		Run: func(_ *cobra.Command, args []string) {
			caradhas.NormalizeFlags(flags)

			token := args[0]

			c := caradhas.NewCaradhasManager(token)
			fmt.Println(c.Ping())
		},
	})

	caradhrasCmd.PersistentFlags().IntVarP(
		&flags.Port, "port", "p", 80,
		fmt.Sprintf("specify the communication port that will be used (%d-%d): -p 8091", caradhas.MIN_PORT_VALUE, caradhas.MAX_PORT_VALUE),
	)

	caradhrasCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Get info about the version of Caradhas API",
		Long:  `Get information about the version of Caradhas API`,
		Run: func(_ *cobra.Command, _ []string) {

		},
	})

	return caradhrasCmd
}

package sources

import (
	"fmt"

	"github.com/samuelmachado/cli-poc-dock/internal/domains/caradhras"
	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"

	"github.com/spf13/cobra"
)

/*
type CaradhasI interface {
}
type Caradhras struct {
}
*/

//Caradhras is the command that handle this source
func Caradhras(flags *cmd.Flags) (caradhrasCmd *cobra.Command) {
	manager := caradhras.NewCaradhasManager()

	caradhrasCmd = &cobra.Command{
		Use:   "caradhras",
		Short: "use to access the Caradhras SDK",
		Long:  `command used to access Caradhras SDK methods. You can also check the official documentation by going to https://lighthouse.dock.tech/`,
		Run: func(cmd *cobra.Command, _ []string) {
			cmd.Help()
		},
	}

	caradhrasCmd.AddCommand(&cobra.Command{
		Use:   "init",
		Short: "init the Caradhras SDK",
		Long:  `Initializes the Caradhras SDK. The arg 0 is used for the token. e.g.: Caradhras init eyJra...bEw`,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) <= 0 {
				return caradhras.ErrTokenInvalid
			}

			/*
				token := args[0]
				err := caradhras.ValidateInit(token)
				if err != nil {
					return err
				}
			*/

			return nil

			//TO DO:
			// pass this validation to domain?
		},
		Run: func(_ *cobra.Command, args []string) {
			caradhras.NormalizeFlags(flags)
			token := args[0]
			manager.Auth(token)
			fmt.Println(manager.Ping())
		},
	})

	caradhrasCmd.PersistentFlags().IntVarP(
		&flags.Port, "profile", "p", 80,
		fmt.Sprintf("specify the communication port that will be used (%d-%d): -p 8091", caradhras.PORT_MIN, caradhras.PORT_MAX),
	)

	caradhrasCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Get info about the version of Caradhras SDK",
		Long:  `Get information about the version of Caradhras SDK`,
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Println(manager.Version())
		},
	})

	return caradhrasCmd
}

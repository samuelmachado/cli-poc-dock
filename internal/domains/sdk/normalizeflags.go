package sdk

import (
	"fmt"

	"github.com/samuelmachado/cli-poc-dock/internal/cmd/normalizers"
	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
)

const PORT_MIN = 1
const PORT_MAX = 65535

//NormalizeFlags normalize flags given
func NormalizeFlags(flags *cmd.Flags) error {
	err := normalizers.NormalizeGlobalFlags(flags)
	flags.Port = normalizePortFlag(flags.Port)
	return err
}

func normalizePortFlag(format int) int {
	if format < PORT_MIN {
		panic(
			fmt.Errorf(
				"Error port must to be greater than %v",
				PORT_MIN,
			),
		)
	}
	if format > PORT_MAX {
		panic(
			fmt.Errorf(
				"Error port must to be less than %v",
				PORT_MAX,
			),
		)
	}
	return format
}

package caradhas

import (
	"fmt"

	"github.com/samuelmachado/cli-poc-dock/internal/cmd/normalizers"
	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
)

//const MIN_TOKEN_LENGTH = 5
const MIN_PORT_VALUE = 1
const MAX_PORT_VALUE = 65535

//NormalizeFlags normalize flags given
func NormalizeFlags(flags *cmd.Flags) {
	normalizers.NormalizeGlobalFlags(flags)
	flags.Port = normalizePortFlag(flags.Port)
	//flags.Token = normalizeTokenFlag(flags.Token)
}

func normalizePortFlag(format int) int {
	if format < MIN_PORT_VALUE {
		panic(
			fmt.Errorf(
				"Error port must to be greater than %v",
				MIN_PORT_VALUE,
			),
		)
	}
	if format > MAX_PORT_VALUE {
		panic(
			fmt.Errorf(
				"Error port must to be less than %v",
				MAX_PORT_VALUE,
			),
		)
	}
	return format
}

//normalizeToken
/*
func normalizeTokenFlag(format string) string {

	format = strings.ToLower(format)

	if len(format) < MIN_TOKEN_LENGTH {
		panic(
			fmt.Errorf(
				"Error format type not suported for token: %v, suported token must be at least %v characters",
				format, MIN_TOKEN_LENGTH,
			),
		)
	}

	return format

}
*/

package normalizers

import (
	"fmt"
	"strings"

	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
)

//NormalizeFlags normalize flags given
func NormalizeGlobalFlags(flags *cmd.Flags) {
	flags.FormatType = normalizeFormatType(flags.FormatType)
}

//normalizeFormatType
func normalizeFormatType(format string) string {
	ok := false
	allowed := []string{"json", "xml"}

	format = strings.ToLower(format)

	for _, f := range allowed {
		if format == f {
			ok = true
			break
		}
	}

	if ok != true {
		panic(
			fmt.Errorf(
				"Error format type not suported: %v, suported list: %v",
				format, allowed,
			),
		)
	}

	return format

}

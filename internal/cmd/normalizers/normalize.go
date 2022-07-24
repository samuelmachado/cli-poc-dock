package normalizers

import (
	"fmt"
	"strings"

	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
)

//NormalizeFlags normalize flags given
func NormalizeGlobalFlags(flags *cmd.Flags) {
	flags.FormatType, _ = normalizeFormatType(flags.FormatType)
}

//normalizeFormatType
func normalizeFormatType(format string) (string, error) {
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
		return "", fmt.Errorf("Error format type not suported: %v, suported list: %v", format, allowed)
	}

	return format, nil

}

package normalizers

import (
	"strings"

	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
)

var SUPPORTED_FORMATS = []string{"json", "text"}

//NormalizeFlags normalize flags given
func NormalizeGlobalFlags(flags *cmd.Flags) error {
	var err error
	flags.FormatType, err = normalizeFormatType(flags.FormatType)
	return err
}

//normalizeFormatType
func normalizeFormatType(format string) (string, error) {
	ok := false

	format = strings.ToLower(format)

	for _, f := range SUPPORTED_FORMATS {
		if format == f {
			ok = true
			break
		}
	}

	if ok != true {
		return "", ErrInvalidOutputFormat
	}

	return format, nil

}

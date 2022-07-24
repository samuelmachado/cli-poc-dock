package normalizers

import "fmt"

var (
	ErrInvalidOutputFormat = fmt.Errorf("Error format type not suported. suported list: %v ", SUPPORTED_FORMATS)
)

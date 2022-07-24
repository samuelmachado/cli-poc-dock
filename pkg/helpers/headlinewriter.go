package helpers

import "fmt"

const DEFAULT_MARGIN = 2

func PrintHeadline(text string) {
	fmt.Println(CreateHeadline(text))
}

func CreateHeadline(text string) string {
	text = fmt.Sprintf("| %s |", text)
	line := "+"
	for i := 0; i < len(text)-DEFAULT_MARGIN; i++ {
		line += "-"
	}
	line = line + "+"

	output := fmt.Sprintf(`
%s
%s
%s`, line, text, line)
	return output
}

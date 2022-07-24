package helpers

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"syscall"

	"golang.org/x/term"
)

func ReadText(question string, stdin io.Reader) (string, error) {
	reader := bufio.NewReader(stdin)
	fmt.Print(question)
	answer, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	answer = strings.ReplaceAll(answer, "\n", "")
	return answer, nil
}

func ReadPassword(question string, stdin io.Reader) (string, error) {
	fmt.Print(question)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", err
	}
	fmt.Print("******")
	fmt.Println()
	return string(bytePassword), nil
}

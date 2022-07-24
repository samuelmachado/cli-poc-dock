package main

import (
	"fmt"

	"github.com/samuelmachado/cli-poc-dock/internal/cmd"
	v "github.com/samuelmachado/cli-poc-dock/pkg/structs/version"
)

var version string
var commit string
var date string

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Something went wrong:", r)
		}
	}()

	cmd.NewRootCmd(v.FullVersion{
		Version: version,
		Commit:  commit,
		Date:    date,
	})

}

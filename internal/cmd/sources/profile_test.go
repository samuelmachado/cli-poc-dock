package sources

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
	v "github.com/samuelmachado/cli-poc-dock/pkg/structs/version"

	"github.com/stretchr/testify/assert"
)

func Test_ProfileCommand(t *testing.T) {
	type args struct {
		cmdArgs []string
		buffer  *bytes.Buffer
		version v.FullVersion
	}

	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "success, the command should must create a profile",
			args: args{
				cmdArgs: []string{"create"},
				buffer:  bytes.NewBufferString(""),
				version: v.FullVersion{
					Version: "",
					Commit:  "",
					Date:    "",
				},
			},
			want: "",
			err:  nil,
		},
		{
			name: "success, the command must select the profile",
			args: args{
				cmdArgs: []string{"select", "jonhDoeProfile"},
				buffer:  bytes.NewBufferString(""),
				version: v.FullVersion{
					Version: "",
					Commit:  "",
					Date:    "",
				},
			},
			want: "jonhDoeProfile selected",
			err:  nil,
		},
		{
			name: "success, the command must delete the profile",
			args: args{
				cmdArgs: []string{"delete", "jonhDoeProfile"},
				buffer:  bytes.NewBufferString(""),
				version: v.FullVersion{
					Version: "",
					Commit:  "",
					Date:    "",
				},
			},
			want: "jonhDoeProfile deleted",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli := Profile(&cmd.Flags{})
			cli.SetOut(tt.args.buffer)
			cli.SetArgs(tt.args.cmdArgs)
			cli.Execute()

			output, err := ioutil.ReadAll(tt.args.buffer)
			got := string(output)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

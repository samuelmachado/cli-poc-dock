package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	v "github.com/samuelmachado/cli-poc-dock/pkg/structs/version"
	"github.com/stretchr/testify/assert"
)

func Test_VersionCommand(t *testing.T) {
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
			name: "success, the command should return the indicated version/commit/date",
			args: args{
				cmdArgs: []string{"version"},
				buffer:  bytes.NewBufferString(""),
				version: v.FullVersion{
					Version: "1.0.0",
					Commit:  "a4327627a9626b8693a015bdc9c2f8876099beed",
					Date:    "2022-01-01",
				},
			},
			want: "version: 1.0.0\nbuilded at: 2022-01-01\ncommit hash: a4327627a9626b8693a015bdc9c2f8876099beed\n",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewRootCmd(tt.args.version)
			cmd.SetOut(tt.args.buffer)
			cmd.SetArgs(tt.args.cmdArgs)
			cmd.Execute()

			output, err := ioutil.ReadAll(tt.args.buffer)
			got := string(output)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

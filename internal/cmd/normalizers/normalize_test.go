package normalizers

import (
	"testing"

	"github.com/samuelmachado/cli-poc-dock/pkg/structs/cmd"
	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	type args struct {
		format string
	}

	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "success, normalize the output flag for JSON",
			args: args{format: "json"},
			want: "json",
			err:  nil,
		},
		{
			name: "success, normalize the output flag for text",
			args: args{format: "text"},
			want: "text",
			err:  nil,
		},
		{
			name: "success, normalize the output with input in uppercase",
			args: args{format: "JSON"},
			want: "json",
			err:  nil,
		},
		{
			name: "fail, must not accept CSV format",
			args: args{format: "csv"},
			want: "",
			err:  ErrInvalidOutputFormat,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeFormatType(tt.args.format)
			assert.Equal(t, err, tt.err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestNormalizeGlobalFlags(t *testing.T) {
	type args struct {
		flags *cmd.Flags
	}

	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "success, normalize the output flag for JSON",
			args: args{flags: &cmd.Flags{FormatType: "json"}},
			want: "json",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NormalizeGlobalFlags(tt.args.flags)
			assert.Equal(t, err, tt.err)
		})
	}
}

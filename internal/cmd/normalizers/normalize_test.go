package normalizers

import (
	"testing"

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
			name: "success, normalize the output json",
			args: args{format: "JSON"},
			want: "json",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalizeFormatType(tt.args.format)
			assert.ErrorIs(t, err, tt.err)
			assert.Equal(t, got, tt.want)
		})
	}
}

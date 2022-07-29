package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateHeadlineLength(t *testing.T) {
	type args struct {
		input string
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success, the number of characters must be equal for a tiny text",
			args: args{input: "Hi this is a title test"},
			want: 84,
		},
		{
			name: "success, the number of characters must be equal for a big text",
			args: args{input: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum"},
			want: 1347,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateHeadline(tt.args.input)
			assert.Equal(t, len(got), tt.want)
		})
	}
}

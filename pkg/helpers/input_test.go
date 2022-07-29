package helpers

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInput(t *testing.T) {
	type args struct {
		question string
		stdin    bytes.Buffer
	}

	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "success, tests a word being received",
			args: args{question: "What color is Napoleon's white horse?", stdin: *bytes.NewBuffer([]byte("white\n"))},
			want: "white",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ReadText(tt.args.question, &tt.args.stdin)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}

/*
func TestReadPassword(t *testing.T) {
	type args struct {
		question string
		stdin    bytes.Buffer
	}

	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{
			name: "success, tests a password being received",
			args: args{question: "type a pass", stdin: *bytes.NewBuffer([]byte("123456\n"))},
			want: "123456",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ReadPassword(tt.args.question)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
		})
	}
}
*/

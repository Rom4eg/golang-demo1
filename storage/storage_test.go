package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_New(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		root string
		want error
	}{
		{
			name: "FAIL: empty root",
			root: "",
			want: ErrRootDoesntExists,
		},
		{
			name: "PASS: OK",
			root: "/tmp",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			_, err := New(tt.root)
			assert.ErrorIs(t, err, tt.want)
		})
	}
}

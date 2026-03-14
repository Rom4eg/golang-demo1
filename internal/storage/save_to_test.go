package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_SaveTo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		prepare func(*testing.T, string)
		expect  func(*testing.T, error, string)
	}{
		{
			name: "FAIL: file exists",
			prepare: func(t *testing.T, root string) {
				abs := filepath.Join(root, "test")
				f, e := os.Create(abs)
				assert.NoError(t, e)
				defer f.Close()
			},
			expect: func(t *testing.T, err error, root string) {
				assert.ErrorIs(t, err, ErrFileExists)
			},
		},
		{
			name: "PASS: OK",
			prepare: func(t *testing.T, _ string) {

			},
			expect: func(t *testing.T, err error, root string) {
				assert.NoError(t, err)
				abs := filepath.Join(root, "test")
				assert.FileExists(t, abs)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()
			s := &Storage{
				root: tmp,
			}
			tt.prepare(t, tmp)
			e := s.SaveTo("test")
			tt.expect(t, e, tmp)
		})
	}
}

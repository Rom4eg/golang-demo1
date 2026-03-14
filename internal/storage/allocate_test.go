package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_Allocate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		root    string
		size    int64
		prepare func(*testing.T, string, *Storage)
		expect  func(*testing.T, string, error)
	}{
		{
			name: "FAIL: already allocated",
			root: "",
			size: 10,
			prepare: func(t *testing.T, root string, s2 *Storage) {
				abs := filepath.Join(root, "test")
				f, e := os.OpenFile(abs, os.O_CREATE|os.O_WRONLY, 0777)
				assert.NoError(t, e)
				s2.fh = f
			},
			expect: func(t *testing.T, name string, err error) {
				assert.Empty(t, name)
				assert.ErrorIs(t, err, ErrStorageAlreadyAllocated)
			},
		},
		{
			name: "FAIL: zero size",
			root: "",
			size: 0,
			prepare: func(t *testing.T, root string, s2 *Storage) {

			},
			expect: func(t *testing.T, name string, err error) {
				assert.Empty(t, name)
				assert.ErrorIs(t, err, ErrCannotAllocateZeroBytes)
			},
		},
		{
			name: "PASS: OK",
			root: "",
			size: 10,
			prepare: func(t *testing.T, root string, s2 *Storage) {

			},
			expect: func(t *testing.T, name string, err error) {
				assert.NotEmpty(t, name)
				assert.NoError(t, err)

				stat, err := os.Stat(name)
				assert.NoError(t, err)
				assert.Greater(t, stat.Size(), int64(10))
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()
			root := filepath.Join(tmp, tt.root)
			err := os.MkdirAll(root, 0777)
			assert.NoError(t, err)

			s, e := New(root)
			assert.NoError(t, e)
			tt.prepare(t, root, s)
			if s.fh != nil {
				defer s.fh.Close()
			}

			name, err := s.Allocate(tt.size)
			tt.expect(t, name, err)
		})
	}
}

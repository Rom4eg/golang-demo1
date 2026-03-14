package storage

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_SetRoot(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		root    string
		prepare func(*testing.T, string, *Storage)
		expect  func(*testing.T, error)
	}{
		{
			name: "FAIL: root not exists",
			root: "not-exists",
			prepare: func(t *testing.T, s string, _ *Storage) {

			},
			expect: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, ErrRootDoesntExists)
			},
		},
		{
			name: "FAIL: not a dir",
			root: "file",
			prepare: func(t *testing.T, s string, _ *Storage) {
				_, e := os.Create(s)
				assert.NoError(t, e)
			},
			expect: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, ErrRootMustBeDir)
			},
		},
		{
			name: "FAIL: already allocated",
			root: "dir",
			prepare: func(t *testing.T, s string, st *Storage) {
				e := os.MkdirAll(s, 0777)
				assert.NoError(t, e)

				testFile := filepath.Join(s, "test")
				fh, e := os.OpenFile(testFile, os.O_CREATE|os.O_WRONLY, 0777)
				assert.NoError(t, e)
				st.fh = fh
			},
			expect: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, ErrStorageAlreadyAllocated)
			},
		},
		{
			name: "PASS: OK",
			root: "dir",
			prepare: func(t *testing.T, s string, _ *Storage) {
				e := os.MkdirAll(s, 0777)
				assert.NoError(t, e)
			},
			expect: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			tmp := t.TempDir()
			root := filepath.Join(tmp, tt.root)
			s := &Storage{}
			tt.prepare(t, root, s)
			defer func(s *Storage) {
				if s.fh != nil {
					s.fh.Close()
				}
			}(s)

			err := s.SetRoot(root)
			tt.expect(t, err)

		})
	}
}

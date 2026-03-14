package storage

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_Finalize(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		prepare func(*testing.T, string, *Storage)
		expect  func(*testing.T, error, string)
	}{
		{
			name: "FAIL: not allocated",
			prepare: func(t *testing.T, root string, s *Storage) {
				s.fh = nil
			},
			expect: func(t *testing.T, err error, _ string) {
				assert.ErrorIs(t, err, ErrStorageNotAllocated)
			},
		},
		{
			name: "FAIL: allocated with zero length",
			prepare: func(t *testing.T, root string, s *Storage) {
				abs := filepath.Join(root, "test")
				f, e := os.OpenFile(abs, os.O_CREATE|os.O_WRONLY, 0777)
				assert.NoError(t, e)
				s.fh = f
				s.length = 0
			},
			expect: func(t *testing.T, err error, _ string) {
				assert.ErrorIs(t, err, ErrStorageNotAllocated)
			},
		},
		{
			name: "FAIL: save file not set",
			prepare: func(t *testing.T, root string, s *Storage) {
				abs := filepath.Join(root, "test")
				f, e := os.OpenFile(abs, os.O_CREATE|os.O_WRONLY, 0777)
				assert.NoError(t, e)
				s.fh = f
				s.length = 5
			},
			expect: func(t *testing.T, err error, _ string) {
				assert.ErrorIs(t, err, ErrSaveFileNotSet)
			},
		},
		{
			name: "FAIL: file closed",
			prepare: func(t *testing.T, root string, s *Storage) {
				abs := filepath.Join(root, "test")
				f, e := os.OpenFile(abs, os.O_CREATE|os.O_WRONLY, 0777)
				assert.NoError(t, e)
				s.fh = f
				s.length = 5
				s.save = "test_new"
				e = f.Close()
				assert.NoError(t, e)
			},
			expect: func(t *testing.T, err error, _ string) {
				assert.Error(t, err)
			},
		},
		{
			name: "PASS: OK",
			prepare: func(t *testing.T, root string, s *Storage) {
				abs := filepath.Join(root, "test")
				f, e := os.OpenFile(abs, os.O_CREATE|os.O_WRONLY, 0777)
				assert.NoError(t, e)

				n, e := f.Write([]byte{'t', 'e', 's', 't', '0'})
				assert.NoError(t, e)
				assert.Equal(t, 5, n)

				s.fh = f
				s.length = 4
				s.save = "test_new"

			},
			expect: func(t *testing.T, err error, root string) {
				assert.NoError(t, err)

				abs := filepath.Join(root, "test_new")
				assert.FileExists(t, abs)

				f, e := os.Open(abs)
				assert.NoError(t, e)
				defer f.Close()

				var buf bytes.Buffer
				n, e := io.Copy(&buf, f)
				assert.NoError(t, e)
				assert.Equal(t, int64(4), n)
				assert.Equal(t, []byte{'t', 'e', 's', 't'}, buf.Bytes())
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

			tt.prepare(t, tmp, s)
			if s.fh != nil {
				defer s.fh.Close()
			}

			err := s.Finalize()
			tt.expect(t, err, tmp)
		})
	}
}

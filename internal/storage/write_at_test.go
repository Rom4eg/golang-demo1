package storage

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage_WriteAt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		prepare   func(*testing.T, string, *Storage)
		expect    func(*testing.T, string)
		expectErr func(*testing.T, error)
	}{
		{
			name: "FAIL: not allocated",
			prepare: func(t *testing.T, root string, s *Storage) {
				s.fh = nil
			},
			expect: func(t *testing.T, _ string) {

			},

			expectErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, ErrStorageNotAllocated)
			},
		},
		{
			name: "PASS: write zero bytes",
			prepare: func(t *testing.T, root string, s *Storage) {
				abs := filepath.Join(root, "test")
				f, e := os.OpenFile(abs, os.O_CREATE|os.O_WRONLY, 0777)
				assert.NoError(t, e)
				s.fh = f
			},
			expect: func(t *testing.T, _ string) {
			},

			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "PASS: OK",
			prepare: func(t *testing.T, root string, s *Storage) {
				abs := filepath.Join(root, "test")
				f, e := os.OpenFile(abs, os.O_CREATE|os.O_WRONLY, 0777)
				assert.NoError(t, e)
				s.fh = f
			},
			expect: func(t *testing.T, root string) {
				abs := filepath.Join(root, "test")
				f, e := os.Open(abs)
				assert.NoError(t, e)
				defer f.Close()

				var buf bytes.Buffer
				n, e := io.Copy(&buf, f)
				assert.Greater(t, n, int64(0))
				assert.NoError(t, e)

				expect := []byte{0, 0, 0, 0, 0, 't', 'e', 's', 't'}
				assert.Equal(t, buf.Bytes(), expect)
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			root := t.TempDir()
			s := &Storage{}
			tt.prepare(t, root, s)
			_, e := s.WriteAt([]byte("test"), 5)
			tt.expectErr(t, e)

			if s.fh != nil {
				e = s.fh.Close()
				assert.NoError(t, e)
			}

			tt.expect(t, root)
		})
	}
}

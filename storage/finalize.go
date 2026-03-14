package storage

import (
	"os"
	"path/filepath"
)

func (s *Storage) Finalize() error {
	if s.fh == nil || s.length < 1 {
		return ErrStorageNotAllocated
	}

	if s.save == "" {
		return ErrSaveFileNotSet
	}

	e := s.fh.Truncate(s.length)
	if e != nil {
		return e
	}

	e = s.fh.Close()
	if e != nil {
		return e
	}

	old := s.fh.Name()
	new := filepath.Join(s.root, s.save)
	return os.Rename(old, new)
}

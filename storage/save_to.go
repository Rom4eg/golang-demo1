package storage

import (
	"os"
	"path/filepath"
)

func (s *Storage) SaveTo(save string) error {
	abs := filepath.Join(s.root, save)

	if _, err := os.Stat(abs); err == nil {
		return ErrFileExists
	} else if !os.IsNotExist(err) {
		return err
	}

	_, e := os.OpenFile(abs, os.O_CREATE|os.O_EXCL, os.ModePerm)
	if e != nil {
		return e
	}

	s.save = save
	return nil
}

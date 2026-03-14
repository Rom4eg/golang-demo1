package storage

import (
	"os"
	"path/filepath"
)

func (s *Storage) Allocate(n int64) (string, error) {
	if s.fh != nil {
		return "", ErrStorageAlreadyAllocated
	}

	if n < 1 {
		return "", ErrCannotAllocateZeroBytes
	}
	s.length = n

	absName := ""
	for {
		name := s.GenerateTmpName()
		absName = filepath.Join(s.root, name)
		if _, err := os.Stat(absName); err == nil {
			continue
		} else if !os.IsNotExist(err) {
			return "", err
		} else if os.IsNotExist(err) {
			break
		}
	}

	fh, err := os.OpenFile(absName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return "", err
	}

	_, err = fh.Seek(s.length+1, 0)
	if err != nil {
		return "", err
	}

	_, err = fh.Write([]byte{0})
	if err != nil {
		return "", err
	}

	_, err = fh.Seek(0, 0)
	if err != nil {
		return "", err
	}

	s.fh = fh

	return absName, nil
}

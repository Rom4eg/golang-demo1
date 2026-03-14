package storage

import "os"

func (s *Storage) SetRoot(root string) error {
	if s.fh != nil {
		return ErrStorageAlreadyAllocated
	}

	if stat, err := os.Stat(root); err != nil {
		return ErrRootDoesntExists
	} else if !stat.IsDir() {
		return ErrRootMustBeDir
	}

	s.root = root
	return nil
}

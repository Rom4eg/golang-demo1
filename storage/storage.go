package storage

import (
	"os"
	"sync"
)

type Storage struct {
	root   string
	fh     *os.File
	save   string
	length int64

	rm sync.RWMutex
}

func New(root string) (*Storage, error) {
	s := &Storage{}
	e := s.SetRoot(root)
	if e != nil {
		return nil, e
	}

	s.rm = sync.RWMutex{}
	return s, nil
}

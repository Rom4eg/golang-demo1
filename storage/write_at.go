package storage

import "fmt"

func (s *Storage) WriteAt(data []byte, offset int64) (int, error) {
	if s.fh == nil {
		return 0, ErrStorageNotAllocated
	}

	if len(data) < 1 {
		return 0, nil
	}

	s.rm.Lock()
	defer s.rm.Unlock()
	n, err := s.fh.WriteAt(data, offset)
	if err != nil {
		return n, err
	}

	if n < len(data) {
		return n, fmt.Errorf("writen %d bytes, while recieved %d bytes: %w", n, len(data), err)
	}

	return n, nil
}

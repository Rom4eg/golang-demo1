package storage

import "errors"

var (
	ErrRootDoesntExists        = errors.New("root doesn't exist")
	ErrRootMustBeDir           = errors.New("root must be a directory")
	ErrStorageAlreadyAllocated = errors.New("storage already allocated")
	ErrWriteError              = errors.New("write error")
	ErrStorageNotAllocated     = errors.New("storage not allocated")
	ErrFileExists              = errors.New("save file exists")
	ErrSaveFileNotSet          = errors.New("save file not set")
	ErrCannotAllocateZeroBytes = errors.New("cannot allocate zero bytes")
)

package downloader

import "errors"

var (
	ErrStorageNotInitialized = errors.New("storage not initialized")
	ErrUnexpectedStatusCode  = errors.New("unexpected status code")
)

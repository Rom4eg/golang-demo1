package config

import "errors"

var (
	ErrUrlIsRequired    = errors.New("url is required")
	ErrOutputIsRequired = errors.New("output is required")
	ErrChunkSizeToSmall = errors.New("chunk size is too small")
)

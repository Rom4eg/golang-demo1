package target

import "errors"

var (
	ErrUnexpectedStatusCode = errors.New("unexpected status code")
	ErrIncorrectFile        = errors.New("incorrect file")
)

package exceptions

import "errors"

var (
	ErrNotFound     = errors.New("task not found")
	ErrInvalidInput = errors.New("invalid input")
)

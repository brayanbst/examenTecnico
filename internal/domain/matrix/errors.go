package matrix

import "errors"

var (
	ErrEmptyMatrix  = errors.New("matrix must be non-empty")
	ErrInvalidShape = errors.New("all rows must have the same length")
)

// Package local provides the features of the Local API.
package local

import "errors"

var (
	ErrUnsupportedFormat = errors.New("format must be either json or xml")
	ErrPageOutOfBound    = errors.New("page must be between 1 and 45")
	ErrEndPage           = errors.New("page reaches the end")
)

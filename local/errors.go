// Package local provides the features of the Local API.
package local

import "errors"

var (
	ErrUnsupportedFormat       = errors.New("format must be either json or xml")
	ErrPageOutOfBound          = errors.New("page must be between 1 and 45")
	ErrRadiusOutOfBound        = errors.New("radius must be between 0 and 20000")
	ErrUnsupportedSortingOrder = errors.New("sorting order must be either accuracy or distance")
	ErrEndPage                 = errors.New("page reaches the end")
)

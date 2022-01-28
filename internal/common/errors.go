package common

import "errors"

var (
	ErrUnsupportedFormat       = errors.New("unsupported format")
	ErrPageOutOfBound          = errors.New("page out of bound")
	ErrSizeOutOfBound          = errors.New("size out of bound")
	ErrEndPage                 = errors.New("page reaches the end")
	ErrUnsupportedSortingOrder = errors.New("unsupported sorting order")
)

package vision

import "errors"

var (
	ErrUnsupportedFormat = errors.New("file format must be either jpg or png")
)

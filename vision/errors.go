package vision

import "errors"

var (
	ErrUnsupportedFormat = errors.New("file format must be either jpg or png")
	ErrOverTheFileSize   = errors.New("file size must be 2 mb or less")
)

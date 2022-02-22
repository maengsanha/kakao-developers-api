package local

import (
	"errors"
	"internal/common"
)

var (
	Done                            = common.ErrEndPage
	ErrUnsupportedCategoryGroupCode = errors.New(
		`category group code must be one of the following options:
		MT1, CS2, PS3, SC4, AC5, PK6, OL7, SW8, CT1, AG2, PO3, AT4, FD6, CE7, HP8, PM9, BK9, AD5`)
	ErrRadiusOutOfBound = errors.New("radius must be between 0 and 20000")
)

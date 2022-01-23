// Package local provides the features of the Local API.
package local

import "errors"

var (
	ErrUnsupportedFormat            = errors.New("format must be either json or xml")
	ErrPageOutOfBound               = errors.New("page must be between 1 and 45")
	ErrRadiusOutOfBound             = errors.New("radius must be between 0 and 20000")
<<<<<<< Updated upstream
	ErrUnsupportedSortingOrder      = errors.New("sorting order must be either accuracy or distance")
	ErrEndPage                      = errors.New("page reaches the end")
	ErrUnsupportedCategoryGroupCode = errors.New("category group code must be one of the following options:\n MT1, CS2, PS3, SC4, AC5, PK6, OL7, SW8, CT1, AG2, PO3, AT4, FD6, CE7, HP8, PM9, BK9, AD5")
=======
	ErrUnsupportedCategoryGroupCode = errors.New("Category group code must be either MT1, CS2, PS3, SC4, AC5, PK6, OL7, SW8, CT1, AG2, PO3, AT4, FD6, CE7, HP8, PM9, BK9, AD5")
	ErrUnsupportedSortingOrder      = errors.New("sorting order must be either accuracy or distance")
	ErrEndPage                      = errors.New("page reaches the end")
>>>>>>> Stashed changes
)

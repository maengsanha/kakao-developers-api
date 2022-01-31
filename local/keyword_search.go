package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// KeywordSearchIterator is a lazy keyword search iterator.
type KeywordSearchIterator struct {
	Query             string
	CategoryGroupCode string
	Format            string
	AuthKey           string
	X                 string
	Y                 string
	Radius            int
	Rect              string
	Page              int
	Size              int
	Sort              string
	end               bool
}

// PlaceSearchByKeyword provides the search results for places that match @query
// in the specified sorting order.
//
// Details can be referred to
// https://developers.kakao.com/docs/latest/ko/local/dev-guide#search-by-keyword.
func PlaceSearchByKeyword(query string) *KeywordSearchIterator {
	return &KeywordSearchIterator{
		Query:             url.QueryEscape(strings.TrimSpace(query)),
		CategoryGroupCode: "",
		Format:            "json",
		AuthKey:           common.KeyPrefix,
		X:                 "",
		Y:                 "",
		Radius:            0,
		Rect:              "",
		Page:              1,
		Size:              15,
		Sort:              "accuracy",
	}
}

// FormatAs sets the request format to @format (json or xml).
func (ki *KeywordSearchIterator) FormatAs(format string) *KeywordSearchIterator {
	switch format {
	case "json", "xml":
		ki.Format = format
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ki
}

// AuthorizeWith sets the authorization key to @key.
func (ki *KeywordSearchIterator) AuthorizeWith(key string) *KeywordSearchIterator {
	ki.AuthKey = common.FormatKey(key)
	return ki
}

// Category sets the category group code of k.
// There are a few available category group codes:
//
// MT1: Large Supermarket
//
// CS2: Convenience Store
//
// PS3: Daycare Center, Kindergarten
//
// SC4: School
//
// AC5: Academic
//
// PK6: Parking
//
// OL7: Gas Station, Charging Station
//
// SW8: Subway Station
//
// CT1: Culture Facility
//
// AG2: Brokerage
//
// PO3: Public Institution
//
// AT4: Tourist Attractions
//
// FD6: Restaurant
//
// CE7: Cafe
//
// HP8: Hospital
//
// PM9: Pharmacy
//
// BK9: Bank
//
// AD5: Accommodation
func (ki *KeywordSearchIterator) Category(groupcode string) *KeywordSearchIterator {
	switch groupcode {
	case "MT1", "CS2", "PS3", "SC4", "AC5", "PK6", "OL7", "SW8", "CT1",
		"AG2", "PO3", "AT4", "FD6", "CE7", "HP8", "PM9", "BK9", "AD5", "":
		ki.CategoryGroupCode = groupcode
	default:
		panic(ErrUnsupportedCategoryGroupCode)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ki
}

// WithCoordinates sets the X and Y coordinates of k.
func (ki *KeywordSearchIterator) WithCoordinates(x, y float64) *KeywordSearchIterator {
	ki.X = strconv.FormatFloat(x, 'f', -1, 64)
	ki.Y = strconv.FormatFloat(y, 'f', -1, 64)
	return ki
}

// WithRadius searches places around a specific area along with @x and @y.
//
// @radius is the distance (a value between 0 and 20000) from the center coordinates to an axis of rotation in meters.
func (ki *KeywordSearchIterator) WithRadius(radius int) *KeywordSearchIterator {
	if 0 <= radius && radius <= 20000 {
		ki.Radius = radius
	} else {
		panic(ErrRadiusOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ki
}

// WithRect limits the search area, such as when searching places within the map screen.
func (ki *KeywordSearchIterator) WithRect(xMin, yMin, xMax, yMax float64) *KeywordSearchIterator {
	ki.Rect = strings.Join([]string{
		strconv.FormatFloat(xMin, 'f', -1, 64),
		strconv.FormatFloat(yMin, 'f', -1, 64),
		strconv.FormatFloat(xMax, 'f', -1, 64),
		strconv.FormatFloat(yMax, 'f', -1, 64)}, ",")
	return ki
}

// Result sets the result page number (a value between 1 and 45).
func (ki *KeywordSearchIterator) Result(page int) *KeywordSearchIterator {
	if 1 <= page && page <= 45 {
		ki.Page = page
	} else {
		panic(common.ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ki
}

// Display sets the number of documents displayed on a single page (a value between 1 and 45).
func (ki *KeywordSearchIterator) Display(size int) *KeywordSearchIterator {
	if 1 <= size && size <= 45 {
		ki.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ki
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or distance. (default is accuracy)
//
// In the case of distance, X and Y coordinates are required as a reference coordinates.
func (ki *KeywordSearchIterator) SortBy(order string) *KeywordSearchIterator {
	switch order {
	case "accuracy", "distance":
		ki.Sort = order
	default:
		panic(common.ErrUnsupportedSortingOrder)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ki
}

// Next returns the place search result and proceeds the iterator to the next page.
func (ki *KeywordSearchIterator) Next() (res PlaceSearchResult, err error) {
	if ki.end {
		return res, common.ErrEndPage
	}

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%ssearch/keyword.%s?query=%s&category_group_code=%s&x=%s&y=%s&radius=%d&rect=%s&page=%d&size=%d&sort=%s",
			prefix, ki.Format, ki.Query, ki.CategoryGroupCode, ki.X, ki.Y, ki.Radius, ki.Rect, ki.Page, ki.Size, ki.Sort), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(common.Authorization, ki.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if ki.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if ki.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	ki.end = res.Meta.IsEnd || 45 < ki.Page

	ki.Page++

	return
}

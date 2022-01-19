// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// KeywordSearchResult represents a keyword search result.
type KeywordSearchResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount    int        `json:"total_count" xml:"total_count"`
		PageableCount int        `json:"pageable_count" xml:"pageable_count"`
		IsEnd         bool       `json:"is_end" xml:"is_end"`
		SameName      RegionInfo `json:"same_name" xml:"same_name"`
	} `json:"meta" xml:"meta"`
	Documents []Place `json:"documents" xml:"documents"`
}

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
		AuthKey:           "KakaoAK ",
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
func (k *KeywordSearchIterator) FormatAs(format string) *KeywordSearchIterator {
	switch format {
	case "json", "xml":
		k.Format = format
	default:
		panic(ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return k
}

// AuthorizeWith sets the authorization key to @key.
func (k *KeywordSearchIterator) AuthorizeWith(key string) *KeywordSearchIterator {
	k.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return k
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
func (k *KeywordSearchIterator) Category(groupcode string) *KeywordSearchIterator {
	switch groupcode {
	case "MT1", "CS2", "PS3", "SC4", "AC5", "PK6", "OL7", "SW8", "CT1",
		"AG2", "PO3", "AT4", "FD6", "CE7", "HP8", "PM9", "BK9", "AD5":
		k.CategoryGroupCode = groupcode
	}
	return k
}

// WithCoordinates sets the X and Y coordinates of k.
func (k *KeywordSearchIterator) WithCoordinates(x, y float64) *KeywordSearchIterator {
	k.X = strconv.FormatFloat(x, 'f', -1, 64)
	k.Y = strconv.FormatFloat(y, 'f', -1, 64)
	return k
}

// WithRadius searches places around a specific area along with @x and @y.
//
// @radius is the distance (a value between 0 and 20000) from the center coordinates to an axis of rotation in meters.
func (k *KeywordSearchIterator) WithRadius(radius int) *KeywordSearchIterator {
	if 0 <= radius && radius <= 20000 {
		k.Radius = radius
	} else {
		panic(ErrRadiusOutOfBound)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return k
}

// WithRect limits the search area, such as when searching places within the map screen.
func (k *KeywordSearchIterator) WithRect(xMin, yMin, xMax, yMax float64) *KeywordSearchIterator {
	k.Rect = strings.Join([]string{
		strconv.FormatFloat(xMin, 'f', -1, 64),
		strconv.FormatFloat(yMin, 'f', -1, 64),
		strconv.FormatFloat(xMax, 'f', -1, 64),
		strconv.FormatFloat(yMax, 'f', -1, 64)}, ",")
	return k
}

// Result sets the result page number (a value between 1 and 45).
func (k *KeywordSearchIterator) Result(page int) *KeywordSearchIterator {
	if 1 <= page && page <= 45 {
		k.Page = page
	} else {
		panic(ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return k
}

// Display sets the number of documents displayed on a single page (a value between 1 and 45).
func (k *KeywordSearchIterator) Display(size int) *KeywordSearchIterator {
	if 1 <= size && size <= 45 {
		k.Size = size
	} else {
		panic(errors.New("size must be between 1 and 45"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return k
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or distance. (default is accuracy)
//
// In the case of distance, X and Y coordinates are required as a reference coordinates.
func (k *KeywordSearchIterator) SortBy(order string) *KeywordSearchIterator {
	switch order {
	case "accuracy", "distance":
		k.Sort = order
	}
	return k
}

// Next returns the keyword search result and proceeds the iterator to the next page.
func (k *KeywordSearchIterator) Next() (res KeywordSearchResult, err error) {
	// at first, send request to the API server
	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v2/local/search/keyword.%s?query=%s&category_group_code=%s&x=%s&y=%s&radius=%d&rect=%s&page=%d&size=%d&sort=%s",
			k.Format, k.Query, k.CategoryGroupCode, k.X, k.Y, k.Radius, k.Rect, k.Page, k.Size, k.Sort), nil)

	if err != nil {
		return
	}
	// don't forget to close the request for concurrent request
	req.Close = true

	// set authorization header
	req.Header.Set("Authorization", k.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	// don't forget to close the response body
	defer resp.Body.Close()

	if k.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if k.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	if res.Meta.IsEnd {
		return res, ErrEndPage
	}

	k.Page++

	return
}

// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// CategorySearchResult represents a category search result.
type CategorySearchResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount    int        `json:"total_count" xml:"total_count"`
		PageableCount int        `json:"pageable_count" xml:"pageable_count"`
		IsEnd         bool       `json:"is_end" xml:"is_end"`
		SameName      RegionInfo `json:"same_name" xml:"same_name"`
	} `json:"meta" xml:"meta"`
	Documents []Place `json:"documents" xml:"documents"`
}

type CategorySearchResults []CategorySearchResult

// String implements fmt.Stringer.
func (cr CategorySearchResult) String() string {
	bs, _ := json.MarshalIndent(cr, "", "  ")
	return string(bs)
}

// SaveAs saves crs to @filename.
//
// The file extension could be either .json or .xml.
func (crs CategorySearchResults) SaveAs(filename string) error {
	switch tokens := strings.Split(filename, "."); tokens[len(tokens)-1] {
	case "json":
		if bs, err := json.MarshalIndent(crs, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	case "xml":
		if bs, err := xml.MarshalIndent(crs, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	default:
		return ErrUnsupportedFormat
	}
}

// CategorySearchIterator is a lazy category search iterator.
type CategorySearchIterator struct {
	Query             string
	Format            string
	AuthKey           string
	CategoryGroupCode string
	X                 string
	Y                 string
	Radius            int
	Rect              string
	Page              int
	Size              int
	Sort              string
	end               bool
}

// PlaceSearchByCategory provides the search results for place by group code in the specified order.
//
// This api provides two search options.
//
// 1. Search with @x, @y coordinates and @radius that distance from @x and @y.
//
// 2. Search with @rect that coordinates of left X, left Y, right X, right Y.
//
// There are a few available category codes:
//
// MT1	Large supermarket
//
// CS2	Convenience store
//
// PS3	Kindergarten
//
// SC4	School
//
// AC5	Academic
//
// PK6	Parking
//
// OL7	Gas station, Charging Station
//
// SW8	Subway station
//
// BK9	Bank
//
// CT1	Cultural facility
//
// AG2	Brokerage
//
// PO3	Public institution
//
// AT4	Tourist attraction
//
// AD5	Accommodation
//
// FD6	Restaurant
//
// CE7	Cafe
//
// HP8	Hospital
//
// PM9	Pharmacys
//
// Details can be referred to
// https://developers.kakao.com/docs/latest/en/local/dev-guide#search-by-category.
func PlaceSearchByCategory(groupcode string) *CategorySearchIterator {
	switch groupcode {
	case "MT1", "CS2", "PS3", "SC4", "AC5", "PK6", "OL7", "SW8", "BK9", "CT1", "AG2", "PO3", "AT4", "AD5", "FD6", "CE7", "HP8", "PM9":
		return &CategorySearchIterator{
			Format:            "json",
			AuthKey:           "KakaoAK ",
			CategoryGroupCode: groupcode,
			X:                 "",
			Y:                 "",
			Radius:            0,
			Rect:              "",
			Page:              1,
			Size:              15,
			Sort:              "accuracy",
		}
	default:
		panic(ErrUnsupportedCategoryGroupCode)
	}
}

// FormatAs sets the request format to @format (json or xml).
func (ci *CategorySearchIterator) FormatAs(format string) *CategorySearchIterator {
	switch format {
	case "json", "xml":
		ci.Format = format
	default:
		panic(ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Authorization sets the authorization key to @key.
func (ci *CategorySearchIterator) AuthorizeWith(key string) *CategorySearchIterator {
	ci.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return ci
}

// WithRadius searches places around a specific area along with @x and @y.
//
// @radius is the distance (a value between 0 and 20000) from the center coordinates to an axis of rotation in meters.
func (ci *CategorySearchIterator) WithRadius(x, y float64, radius int) *CategorySearchIterator {
	if 0 <= ci.Radius && ci.Radius <= 20000 {
		ci.X = strconv.FormatFloat(x, 'f', -1, 64)
		ci.Y = strconv.FormatFloat(y, 'f', -1, 64)
		ci.Radius = radius
	} else {
		panic(ErrRadiusOutOfBound)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// WithRect limits the search area, such as when searching places within the map screen.
func (ci *CategorySearchIterator) WithRect(xMin, yMin, xMax, yMax float64) *CategorySearchIterator {
	ci.Rect = strings.Join([]string{
		strconv.FormatFloat(xMin, 'f', -1, 64),
		strconv.FormatFloat(yMin, 'f', -1, 64),
		strconv.FormatFloat(xMax, 'f', -1, 64),
		strconv.FormatFloat(yMax, 'f', -1, 64)}, ",")
	return ci
}

// Result sets the result page number (a value between 1 and 45).
func (ci *CategorySearchIterator) Result(page int) *CategorySearchIterator {
	if 1 <= page && page <= 45 {
		ci.Page = page
	} else {
		panic(ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Display sets the number of documents displayed on a single page (a value between 1 and 15).
func (ci *CategorySearchIterator) Display(size int) *CategorySearchIterator {
	if 1 <= size && size <= 15 {
		ci.Size = size
	} else {
		panic(errors.New("size must be between 1 and 15"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// SortBy sets the ordering type of c to @order.
//
// @order can be accuracy or distance. (default is accuracy)
func (ci *CategorySearchIterator) SortBy(order string) *CategorySearchIterator {
	switch order {
	case "accuracy", "distance":
		ci.Sort = order
	default:
		panic(ErrUnsupportedSortingOrder)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Next returns the category search result.
func (ci *CategorySearchIterator) Next() (res CategorySearchResult, err error) {
	if ci.end {
		return res, ErrEndPage
	}

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v2/local/search/category.%s?category_group_code=%s&page=%d&size=%d&sort=%s&x=%s&y=%s&radius=%d&rect=%s",
			ci.Format, ci.CategoryGroupCode, ci.Page, ci.Size, ci.Sort, ci.X, ci.Y, ci.Radius, ci.Rect), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", ci.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if ci.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if ci.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	ci.end = res.Meta.IsEnd

	ci.Page++

	return
}

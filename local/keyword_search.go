package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// KeywordSearchIterator initializes parameters that used in keyword search.
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

// KeywordSearch provides the search results for places that match @query
// in the specified sorting order.
//
// Details can be referred to
// https://developers.kakao.com/docs/latest/ko/local/dev-guide#search-by-keyword
func KeywordSearch(query string) *KeywordSearchIterator {
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

func (k *KeywordSearchIterator) FormatJSON() *KeywordSearchIterator {
	k.Format = "json"
	return k
}

func (k *KeywordSearchIterator) FormatXML() *KeywordSearchIterator {
	k.Format = "xml"
	return k
}

// AuthorizeWith sets the authorization key to @key.
func (k *KeywordSearchIterator) AuthorizeWith(key string) *KeywordSearchIterator {
	k.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return k
}

// Category group code is used if you want to filter results into categories.
//
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

// Coorinates set X coordinate (longitude) of the center and Y coordinate (latitude) of the center.
// Used to search places around a specific area along with radius.
func (k *KeywordSearchIterator) Coordinates(x, y float64) *KeywordSearchIterator {
	k.X = strconv.FormatFloat(x, 'f', -1, 64)
	k.Y = strconv.FormatFloat(y, 'f', -1, 64)
	return k
}

// RadiusDistance is used to search places around a specific area along with x and y (center coordinates).
//
// @radius : The distance from the center coordinates to an axis of rotation in meters. (between 0 and 20000)
func (k *KeywordSearchIterator) RadiusDistance(radius int) *KeywordSearchIterator {
	if 0 <= radius && radius <= 20000 {
		k.Radius = radius
	}
	return k
}

// Rectangle is used to limit search area, such as when searching places within the map screen.
//
// In the coordinates of left X(@xMin), left Y(@yMin), right X(@xMax), right Y(@yMax) format.
func (k *KeywordSearchIterator) Rectangle(xMin, yMin, xMax, yMax float64) *KeywordSearchIterator {
	k.Rect = strings.Join([]string{strconv.FormatFloat(xMin, 'f', -1, 64),
		strconv.FormatFloat(yMin, 'f', -1, 64),
		strconv.FormatFloat(xMax, 'f', -1, 64),
		strconv.FormatFloat(yMax, 'f', -1, 64)}, ",")
	return k
}

func (k *KeywordSearchIterator) Result(page int) *KeywordSearchIterator {
	if 1 <= page && page <= 45 {
		k.Page = page
	}
	return k
}

func (k *KeywordSearchIterator) Display(size int) *KeywordSearchIterator {
	if 1 <= size && size <= 45 {
		k.Size = size
	}
	return k
}

// Sorting specifies sorting order of the document results.
//
// Sorting order(@order) can be accuracy or distance. (Default: accuracy).
// In the case of distance, x and y values are required as a reference coordinates.
func (k *KeywordSearchIterator) Sorting(order string) *KeywordSearchIterator {
	switch order {
	case "accuracy", "distance":
		k.Sort = order
	}
	return k
}

// Next sends a GET request and
// returns the keyword search result and proceeds the iterator to the next page.
func (k *KeywordSearchIterator) Next() (res KeywordSearchResult, err error) {
	// at first, send request to the API server
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/search/keyword.%s?query=%s&category_group_code=%s&x=%s&y=%s&radius=%d&rect=%s&page=%d&size=%d&sort=%s", k.
		Format, k.Query, k.CategoryGroupCode, k.X, k.Y, k.Radius, k.Rect, k.Page, k.Size, k.Sort), nil)

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

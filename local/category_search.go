// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// CategorySearchResult ...
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

// CategorySearchIterator ...
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
}

// PlaceSearchByCategory ...
func PlaceSearchByCategory(groupcode string) *CategorySearchIterator {
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
}

func (c *CategorySearchIterator) FormatJSON() *CategorySearchIterator {
	c.Format = "json"
	return c
}

func (c *CategorySearchIterator) FormatXML() *CategorySearchIterator {
	c.Format = "xml"
	return c
}

// Authorization ...
func (c *CategorySearchIterator) AuthorizeWith(key string) *CategorySearchIterator {
	c.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return c
}

// WithRadius ...
func (c *CategorySearchIterator) WithRadius(x, y float64, radius int) *CategorySearchIterator {
	if 0 <= c.Radius && c.Radius <= 20000 {
		c.X = strconv.FormatFloat(x, 'f', -1, 64)
		c.Y = strconv.FormatFloat(y, 'f', -1, 64)
		c.Radius = radius
	}

	return c
}

// WithRect ...
func (c *CategorySearchIterator) WithRect(xMin, yMin, xMax, yMax float64) *CategorySearchIterator {
	c.Rect = strings.Join([]string{
		strconv.FormatFloat(xMin, 'f', -1, 64),
		strconv.FormatFloat(yMin, 'f', -1, 64),
		strconv.FormatFloat(xMax, 'f', -1, 64),
		strconv.FormatFloat(yMax, 'f', -1, 64)}, ",")
	return c
}

func (c *CategorySearchIterator) Result(page int) *CategorySearchIterator {
	if 1 <= page && page <= 45 {
		c.Page = page
	}
	return c
}

func (c *CategorySearchIterator) Display(size int) *CategorySearchIterator {
	if 1 <= size && size <= 15 {
		c.Size = size
	}
	return c
}

// SortBy ...
func (c *CategorySearchIterator) SortBy(typ string) *CategorySearchIterator {
	if typ == "accuracy" || typ == "distance" {
		c.Sort = typ
	}
	return c
}

// Next ...
func (c *CategorySearchIterator) Next() (res CategorySearchResult, err error) {
	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v2/local/search/category.%s?category_group_code=%s&page=%d&size=%d&sort=%s&x=%s&y=%s&radius=%d&rect=%s",
			c.Format, c.CategoryGroupCode, c.Page, c.Size, c.Sort, c.X, c.Y, c.Radius, c.Rect), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", c.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if c.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if c.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	if res.Meta.IsEnd {
		return res, ErrEndPage
	}

	c.Page++

	return
}

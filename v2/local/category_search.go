// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type Place struct {
	ID                string `json:"id" xml:"id"`
	PlaceName         string `json:"place_name" xml:"place_name"`
	CategoryName      string `json:"category_name" xml:"category_name"`
	CategoryGroupCode string `json:"category_group_code" xml:"category_group_code"`
	CategoryGroupName string `json:"category_group_name" xml:"category_group_name"`
	Phone             string `json:"phone" xml:"phone"`
	AddressName       string `json:"address_name" xml:"address_name"`
	RoadAddressName   string `json:"road_address_name" xml:"road_address_name"`
	X                 string `json:"x" xml:"x"`
	Y                 string `json:"y" xml:"y"`
	PlaceURL          string `json:"place_url" xml:"place_url"`
	Distance          string `json:"distance" xml:"distance"`
}

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

type RegionInfo struct {
	Region         []string `json:"region" xml:"region"`
	Keyword        string   `json:"keyword" xml:"keyword"`
	SelectedRegion string   `json:"selected_region" xml:"selected_region"`
}

type CategorySearchInitializer struct {
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

func CategorySearch(category_group_code string) *CategorySearchInitializer {
	return &CategorySearchInitializer{
		Format:            "json",
		AuthKey:           "KakaoAK ",
		CategoryGroupCode: category_group_code,
		X:                 "",
		Y:                 "",
		Radius:            0,
		Rect:              "",
		Page:              1,
		Size:              15,
		Sort:              "accuracy",
	}
}

func (d *CategorySearchInitializer) FormatJSON() *CategorySearchInitializer {
	d.Format = "json"
	return d
}

func (d *CategorySearchInitializer) FormatXML() *CategorySearchInitializer {
	d.Format = "xml"
	return d
}

func (d *CategorySearchInitializer) AuthorizeWith(key string) *CategorySearchInitializer {
	d.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return d
}

func (d *CategorySearchInitializer) SetLongitude(x string) *CategorySearchInitializer {
	d.X = x
	return d
}

func (d *CategorySearchInitializer) SetLatitude(y string) *CategorySearchInitializer {
	d.Y = y
	return d
}

func (d *CategorySearchInitializer) SetRadius(radius int) *CategorySearchInitializer {
	if 0 <= radius && radius <= 20000 {
		d.Radius = radius
	}
	return d
}

func (d *CategorySearchInitializer) SetRect(rect string) *CategorySearchInitializer {
	d.Rect = rect
	return d
}

func (d *CategorySearchInitializer) Result(page int) *CategorySearchInitializer {
	if 1 <= page && page <= 45 {
		d.Page = page
	}
	return d
}

func (d *CategorySearchInitializer) Display(size int) *CategorySearchInitializer {
	if 1 <= size && size <= 15 {
		d.Size = size
	}
	return d
}

func (d *CategorySearchInitializer) SortType(sort string) *CategorySearchInitializer {
	if sort == "accuracy" || sort == "distance" {
		d.Sort = sort
	}
	return d
}

func (d *CategorySearchInitializer) Collect() (res CategorySearchResult, err error) {
	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/search/category.%s?category_group_code=%s&page=%d&size=%d&sort=%s&x=%s&y=%s&radius=%d&rect=%s", d.Format, d.CategoryGroupCode, d.Page, d.Size, d.Sort, d.X, d.Y, d.Radius, d.Rect), nil)
	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", d.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if d.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if d.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	return
}

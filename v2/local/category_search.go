// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Place struct {
	Id                string `json:"id" xml:"id"`
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

type CategorySearchResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount    int        `json:"total_count" xml:"total_count"`
		PageableCount int        `json:"pageable_count" xml:"pageable_count"`
		IsEnd         bool       `json:"is_end" xml:"is_end"`
		SameName      RegionInfo `json:"same_name" xml:"same_name"`
	} `json:"meta" xml:"meta"`
	Documents []ComplexAddress `json:"documents" xml:"documents"`
}

type RegionInfo struct {
	Region         []string `json:"region" xml:"region"`
	Keyword        string   `json:"keyword" xml:"keyword"`
	SelectedRegion string   `json:"selected_region" xml:"selected_region"`
}

func CategorySearch(query string) *CategorySearchIterator {
	return &CategorySearchIterator{
		Query:             url.QueryEscape(strings.TrimSpace(query)),
		Format:            "json",
		AuthKey:           "KakaoAK",
		CategoryGroupCode: "",
		X:                 "",
		Y:                 "",
		Radius:            0,
		Rect:              "",
		Page:              1,
		Size:              15,
		Sort:              "accuracy",
	}
}

func (d *CategorySearchIterator) FormatJSON() *CategorySearchIterator {
	d.Format = "json"
	return d
}

func (d *CategorySearchIterator) FormatXML() *CategorySearchIterator {
	d.Format = "xml"
	return d
}

func (d *CategorySearchIterator) AuthorizeWith(key string) *CategorySearchIterator {
	d.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return d
}

// Category code list
// MT1	대형마트
// CS2	편의점
// PS3	어린이집, 유치원
// SC4	학교
// AC5	학원
// PK6	주차장
// OL7	주유소, 충전소
// SW8	지하철역
// BK9	은행
// CT1	문화시설
// AG2	중개업소
// PO3	공공기관
// AT4	관광명소
// AD5	숙박
// FD6	음식점
// CE7	카페
// HP8	병원
// PM9	약국

// Set as Large supermarket
func (d *CategorySearchIterator) SetCategoryMT1(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "MT1"
	return d
}

// Set as Convenience store
func (d *CategorySearchIterator) SetCategoryCS2(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "CS2"
	return d
}

//Set as Kindergarten(Preschool)
func (d *CategorySearchIterator) SetCategoryPS3(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "PS3"
	return d
}

//Set as School
func (d *CategorySearchIterator) SetCategorySC4(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "PK6"
	return d
}

//Set as Cram school
func (d *CategorySearchIterator) SetCategoryAC5(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "AC5"
	return d
}

//Set as Parking space
func (d *CategorySearchIterator) SetCategoryPK6(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "SC4"
	return d
}

//Set as Oil station
func (d *CategorySearchIterator) SetCategoryOL7(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "OL7"
	return d
}

//Set as Subway station
func (d *CategorySearchIterator) SetCategorySW8(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "SW8"
	return d
}

//Set as Bank
func (d *CategorySearchIterator) SetCategoryBK9(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "BK9"
	return d
}

//Set as Cultural facilities
func (d *CategorySearchIterator) SetCategoryCT1(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "CT1"
	return d
}

//Set as Brokage(Real estate)
func (d *CategorySearchIterator) SetCategoryAG2(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "AG2"
	return d
}

//Set as Public institution
func (d *CategorySearchIterator) SetCategoryPO3(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "PO3"
	return d
}

//Set as Sight spot
func (d *CategorySearchIterator) SetCategoryAT4(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "AT4"
	return d
}

//Set as Accommodation
func (d *CategorySearchIterator) SetCategoryAD5(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "AD5"
	return d
}

//Set as Restaurant
func (d *CategorySearchIterator) SetCategoryFD6(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "FD6"
	return d
}

//Set as Caffe
func (d *CategorySearchIterator) SetCategoryCE7(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "CE7"
	return d
}

//Set as Hospital
func (d *CategorySearchIterator) SetCategoryHP8(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "HP8"
	return d
}

//Set as Pharmacy
func (d *CategorySearchIterator) SetCategoryPM9(key string) *CategorySearchIterator {
	d.CategoryGroupCode = "PM9"
	return d
}

func (d *CategorySearchIterator) SetLongitude(x string) *CategorySearchIterator {
	d.X = x
	return d
}

func (d *CategorySearchIterator) SetLatitude(y string) *CategorySearchIterator {
	d.Y = y
	return d
}

func (d *CategorySearchIterator) SetRadius(radius int) *CategorySearchIterator {
	if 0 <= radius && radius <= 20000 {
		d.Radius = radius
	}
	return d
}

func (d *CategorySearchIterator) SetRect(rect string) *CategorySearchIterator {
	d.Rect = rect
	return d
}

func (d *CategorySearchIterator) Result(page int) *CategorySearchIterator {
	if 1 <= page && page <= 45 {
		d.Page = page
	}
	return d
}

func (d *CategorySearchIterator) Display(size int) *CategorySearchIterator {
	if 1 <= size && size <= 15 {
		d.Size = size
	}
	return d
}

func (d *CategorySearchIterator) SortType(sort string) *CategorySearchIterator {
	if sort == "accuracy" || sort == "distance" {
		d.Sort = sort
	}
	return d
}

func (d *CategorySearchIterator) Next() (res AddressSearchResult, err error) {
	client := new(http.Client)
	if d.X != "" && d.Y != "" && d.Radius != 0 {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/search/category.%s?category_group_code=%s&page=%s&size=%s&sort=%s&x=%s&y=%s&radius=%s", d.Format, d.CategoryGroupCode, d.Page, d.Size, d.Sort, d.X, d.Y, d.Radius), nil)
	} else if d.Rect != "" {
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/search/category.%s?category_group_code=%s&page=%s&size=%s&sort=%s&rect=%s", d.Format, d.CategoryGroupCode, d.Page, d.Size, d.Sort, d.Rect), nil)
	} else {
		return
	}
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

	if res.Meta.IsEnd {
		return res, ErrEndPage
	}

	d.Page++

	return
}

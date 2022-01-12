package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type RegionInfo struct {
	Region         []string `json:"region" xml:"region"`
	Keyword        string   `json:"keyword" xml:"keyword"`
	SelectedRegion string   `json:"selected_region" xml:"selected_region"`
}

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

func (k *KeywordSearchIterator) AuthorizeWith(key string) *KeywordSearchIterator {
	k.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return k
}

func (k *KeywordSearchIterator) SetCategoryGroupCode(code string) *KeywordSearchIterator {
	k.CategoryGroupCode = code
	return k
}

func (k *KeywordSearchIterator) SetX(x string) *KeywordSearchIterator {
	k.X = x
	return k
}

func (k *KeywordSearchIterator) SetY(y string) *KeywordSearchIterator {
	k.Y = y
	return k
}

func (k *KeywordSearchIterator) SetRadius(radius int) *KeywordSearchIterator {
	if 0 <= radius && radius <= 20000 {
		k.Radius = radius
	}
	return k
}

func (k *KeywordSearchIterator) SetRect(rect string) *KeywordSearchIterator {
	k.Rect = rect

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

func (k *KeywordSearchIterator) SortType(sort string) *KeywordSearchIterator {
	if sort == "accuracy" || sort == "distance" {
		k.Sort = sort
	}
	return k
}

// Next returns the search result and proceeds the iterator to the next page.
func (k *KeywordSearchIterator) Next() (res KeywordSearchResult, err error) {
	// at first, send request to the API server
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/search/keyword.%s?query=%s&category_group_code=%s&x=%s&y=%s&radius=%d&rect=%s&page=%d&size=%d&sort=%s", k.Format, k.Query, k.CategoryGroupCode, k.X, k.Y, k.Radius, k.Rect, k.Page, k.Size, k.Sort), nil)
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

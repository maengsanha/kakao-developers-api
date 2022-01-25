package daum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// CafeResult represents a document of a Daum cafe search result.
type CafeResult struct {
	WebResult
	CafeName  string `json:"cafename"`
	Thumbnail string `json:"thumbnail"`
}

// CafeSearchResult represents a Daum cafe search result.
type CafeSearchResult struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []CafeResult `json:"documents"`
}

// String implements fmt.Stringer.
func (cr CafeSearchResult) String() string {
	bs, _ := json.MarshalIndent(cr, "", "  ")
	return string(bs)
}

type CafeSearchResults []CafeSearchResult

// SaveAs saves crs to @filename.
func (crs CafeSearchResults) SaveAs(filename string) error {
	if bs, err := json.MarshalIndent(crs, "", "  "); err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, bs, 0644)
	}
}

// CafeSearchIterator is a lazy Cafe search iterator.
type CafeSearchIterator struct {
	Query   string
	AuthKey string
	Sort    string
	Page    int
	Size    int
	end     bool
}

// CafeSearch allows users to search posts by @query in the Daum Cafe service.
//
// See https://developers.kakao.com/docs/latest/en/daum-search/dev-guide#search-cafe for more details.
func CafeSearch(query string) *CafeSearchIterator {
	return &CafeSearchIterator{
		Query:   url.QueryEscape(strings.TrimSpace(query)),
		AuthKey: "KakaoAK ",
		Sort:    "accuracy",
		Page:    1,
		Size:    10,
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ci *CafeSearchIterator) AuthorizeWith(key string) *CafeSearchIterator {
	ci.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return ci
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (ci *CafeSearchIterator) SortBy(order string) *CafeSearchIterator {
	switch order {
	case "accuracy", "recency":
		ci.Sort = order
	default:
		panic("sorting order must be either accuracy or recency")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Result sets the result page number (a value between 1 and 50).
func (ci *CafeSearchIterator) Result(page int) *CafeSearchIterator {
	if 1 <= page && page <= 50 {
		ci.Page = page
	} else {
		panic("page must be between 1 and 50")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Display sets the number of documents displayed on a single page (a value between 1 and 50).
func (ci *CafeSearchIterator) Display(size int) *CafeSearchIterator {
	if 1 <= size && size <= 50 {
		ci.Size = size
	} else {
		panic("size musts be between 1 and 50")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Next returns the cafe search result and proceeds the iterator to the next page.
func (ci *CafeSearchIterator) Next() (res CafeSearchResult, err error) {
	if ci.end {
		return res, ErrEndPage
	}

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%scafe?query=%s&sort=%s&page=%d&size=%d",
			prefix, ci.Query, ci.Sort, ci.Page, ci.Size), nil)

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

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	ci.Page++

	ci.end = res.Meta.IsEnd || ci.Page > 50

	return
}

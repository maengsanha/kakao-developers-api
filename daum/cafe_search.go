package daum

import (
	"encoding/json"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// CafeResult represents a document of a Daum Cafe search result.
type CafeResult struct {
	WebResult
	CafeName  string `json:"cafename"`
	Thumbnail string `json:"thumbnail"`
}

// CafeSearchResult represents a Daum Cafe search result.
type CafeSearchResult struct {
	Meta      common.PageableMeta `json:"meta"`
	Documents []CafeResult        `json:"documents"`
}

// String implements fmt.Stringer.
func (cr CafeSearchResult) String() string { return common.String(cr) }

type CafeSearchResults []CafeSearchResult

// SaveAs saves crs to @filename.
func (crs CafeSearchResults) SaveAs(filename string) error { return common.SaveAsJSON(crs, filename) }

// CafeSearchIterator is a lazy cafe search iterator.
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
		AuthKey: common.KeyPrefix,
		Sort:    "accuracy",
		Page:    1,
		Size:    10,
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ci *CafeSearchIterator) AuthorizeWith(key string) *CafeSearchIterator {
	ci.AuthKey = common.FormatKey(key)
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
		panic(common.ErrUnsupportedSortingOrder)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ci
}

// Result sets the result page number (a value between 1 and 50).
func (ci *CafeSearchIterator) Result(page int) *CafeSearchIterator {
	if 1 <= page && page <= 50 {
		ci.Page = page
	} else {
		panic(common.ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ci
}

// Display sets the number of documents displayed on a single page (a value between 1 and 50).
func (ci *CafeSearchIterator) Display(size int) *CafeSearchIterator {
	if 1 <= size && size <= 50 {
		ci.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
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

	req.Header.Set(common.Authorization, ci.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	ci.Page++

	ci.end = res.Meta.IsEnd || 50 < ci.Page

	return
}

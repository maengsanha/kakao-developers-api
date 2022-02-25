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

// WebResult represents a document of a Daum search result.
type WebResult struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
	URL      string `json:"url"`
	Datetime string `json:"datetime"`
}

// DocumentSearchResult represents a Daum search result.
type DocumentSearchResult struct {
	Meta      common.PageableMeta `json:"meta"`
	Documents []WebResult         `json:"documents"`
}

// String implements fmt.Stringer.
func (dr DocumentSearchResult) String() string { return common.String(dr) }

type DocumentSearchResults []DocumentSearchResult

// SaveAs saves drs to @filename.
func (drs DocumentSearchResults) SaveAs(filename string) error {
	return common.SaveAsJSON(drs, filename)
}

// DocumentSearchIterator is a lazy document search iterator.
type DocumentSearchIterator struct {
	Query   string
	Sort    string
	Page    int
	Size    int
	AuthKey string
	end     bool
}

// DocumentSearch allows to search web documents by @query in the Daum Search service.
//
// See https://developers.kakao.com/docs/latest/ko/daum-search/dev-guide#search-doc for more details.
func DocumentSearch(query string) *DocumentSearchIterator {
	return &DocumentSearchIterator{
		Query:   url.QueryEscape(strings.TrimSpace(query)),
		Sort:    "accuracy",
		Page:    1,
		Size:    10,
		AuthKey: common.KeyPrefix,
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (di *DocumentSearchIterator) AuthorizeWith(key string) *DocumentSearchIterator {
	di.AuthKey = common.FormatKey(key)
	return di
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (di *DocumentSearchIterator) SortBy(order string) *DocumentSearchIterator {
	switch order {
	case "accuracy", "recency":
		di.Sort = order
	default:
		panic(common.ErrUnsupportedSortingOrder)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return di
}

// Result sets the result page number (a value between 1 and 50).
func (di *DocumentSearchIterator) Result(page int) *DocumentSearchIterator {
	if 1 <= page && page <= 50 {
		di.Page = page
	} else {
		panic(common.ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return di
}

// Display sets the number of documents displayed on a single page (a value between 1 and 50).
func (di *DocumentSearchIterator) Display(size int) *DocumentSearchIterator {
	if 1 <= size && size <= 50 {
		di.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return di
}

// Next returns the document search result and proceeds the iterator to the next page.
func (di *DocumentSearchIterator) Next() (res DocumentSearchResult, err error) {
	if di.end {
		return res, ErrEndPage
	}

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%sweb?query=%s&sort=%s&page=%d&size=%d",
			prefix, di.Query, di.Sort, di.Page, di.Size), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(common.Authorization, di.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	di.end = res.Meta.IsEnd || 50 < di.Page

	di.Page++

	return
}

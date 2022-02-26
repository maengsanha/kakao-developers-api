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

// BlogResult represents a document of a blog search result.
type BlogResult struct {
	WebResult
	Blogname  string `json:"blogname"`
	Thumbnail string `json:"thumbnail"`
}

// BlogSearchResult represents a blog search result.
type BlogSearchResult struct {
	Meta      common.PageableMeta `json:"meta"`
	Documents []BlogResult        `json:"documents"`
}

// String implements fmt.Stringer.
func (br BlogSearchResult) String() string { return common.String(br) }

type BlogSearchResults []BlogSearchResult

// SaveAs saves brs to @filename.
func (brs BlogSearchResults) SaveAs(filename string) error { return common.SaveAsJSON(brs, filename) }

// BlogSearchIterator is a lazy blog search iterator.
type BlogSearchIterator struct {
	Query   string
	Sort    string
	Page    int
	Size    int
	AuthKey string
	end     bool
}

// BlogSearch allows to search blog posts by @query in the Daum Blog service.
//
// See https://developers.kakao.com/docs/latest/ko/daum-search/dev-guide#search-blog for more details.
func BlogSearch(query string) *BlogSearchIterator {
	return &BlogSearchIterator{
		Query:   url.QueryEscape(strings.TrimSpace(query)),
		Sort:    "accuracy",
		Page:    1,
		Size:    10,
		AuthKey: common.KeyPrefix,
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (it *BlogSearchIterator) AuthorizeWith(key string) *BlogSearchIterator {
	it.AuthKey = common.FormatKey(key)
	return it
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (it *BlogSearchIterator) SortBy(order string) *BlogSearchIterator {
	switch order {
	case "accuracy", "recency":
		it.Sort = order
	default:
		panic(common.ErrUnsupportedSortingOrder)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Result sets the result page number (a value between 1 and 50).
func (it *BlogSearchIterator) Result(page int) *BlogSearchIterator {
	if 1 <= page && page <= 50 {
		it.Page = page
	} else {
		panic(common.ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Display sets the number of documents displayed on a single page (a value between 1 and 50).
func (it *BlogSearchIterator) Display(size int) *BlogSearchIterator {
	if 1 <= size && size <= 50 {
		it.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Next returns the blog search result and proceeds the iterator to the next page.
func (it *BlogSearchIterator) Next() (res BlogSearchResult, err error) {
	if it.end {
		return res, Done
	}

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%sblog?query=%s&sort=%s&page=%d&size=%d",
			prefix, it.Query, it.Sort, it.Page, it.Size), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(common.Authorization, it.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	it.end = res.Meta.IsEnd || 50 < it.Page

	it.Page++

	return
}

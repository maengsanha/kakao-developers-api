package daum

import (
	"encoding/json"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// VClipResult represents a document of a video search result.
type VClipResult struct {
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Datetime  time.Time `json:"datetime"`
	PlayTime  int       `json:"play_time"`
	Thumbnail string    `json:"thumbnail"`
	Author    string    `json:"author"`
}

// VideoSearchResult represents a video search result.
type VideoSearchResult struct {
	Meta      common.PageableMeta `json:"meta"`
	Documents []VClipResult       `json:"documents"`
}

// String implements fmt.Stringer.
func (vr VideoSearchResult) String() string { return common.String(vr) }

type VideoSearchResults []VideoSearchResult

// SaveAs saves vrs to @filename.
func (vrs VideoSearchResults) SaveAs(filename string) error { return common.SaveAsJSON(vrs, filename) }

// VideoSearchIterator is a lazy video search iterator.
type VideoSearchIterator struct {
	Query   string
	Sort    string
	Page    int
	Size    int
	AuthKey string
	end     bool
}

// VideoSearch allows users to search videos by @query on the video platforms such as Youtube or Kakao TV.
//
// For more details visit https://developers.kakao.com/docs/latest/en/daum-search/dev-guide#search-video.
func VideoSearch(query string) *VideoSearchIterator {
	return &VideoSearchIterator{
		Query:   url.QueryEscape(strings.TrimSpace(query)),
		Sort:    "accuracy",
		Page:    1,
		Size:    15,
		AuthKey: common.KeyPrefix,
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (vi *VideoSearchIterator) AuthorizeWith(key string) *VideoSearchIterator {
	vi.AuthKey = common.FormatKey(key)
	return vi
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (vi *VideoSearchIterator) SortBy(order string) *VideoSearchIterator {
	switch order {
	case "accuracy", "recency":
		vi.Sort = order
	default:
		panic(common.ErrUnsupportedSortingOrder)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return vi
}

// Result sets the result page number (a value between 1 and 15).
func (vi *VideoSearchIterator) Result(page int) *VideoSearchIterator {
	if 1 <= page && page <= 15 {
		vi.Page = page
	} else {
		panic(common.ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return vi
}

// Display sets the number of documents displayed on a single page (a value between 1 and 30).
func (vi *VideoSearchIterator) Display(size int) *VideoSearchIterator {
	if 1 <= size && size <= 30 {
		vi.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return vi
}

// Next returns the video search result and proceeds the iterator to the next page.
func (vi *VideoSearchIterator) Next() (res VideoSearchResult, err error) {
	if vi.end {
		return res, ErrEndPage
	}

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%svclip?query=%s&sort=%s&page=%d&size=%d",
			prefix, vi.Query, vi.Sort, vi.Page, vi.Size), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(common.Authorization, vi.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	vi.end = res.Meta.IsEnd || 15 < vi.Page

	vi.Page++

	return
}

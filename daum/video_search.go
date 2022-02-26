package daum

import (
	"encoding/json"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
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
func (it *VideoSearchIterator) AuthorizeWith(key string) *VideoSearchIterator {
	it.AuthKey = common.FormatKey(key)
	return it
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (it *VideoSearchIterator) SortBy(order string) *VideoSearchIterator {
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

// Result sets the result page number (a value between 1 and 15).
func (it *VideoSearchIterator) Result(page int) *VideoSearchIterator {
	if 1 <= page && page <= 15 {
		it.Page = page
	} else {
		panic(common.ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Display sets the number of documents displayed on a single page (a value between 1 and 30).
func (it *VideoSearchIterator) Display(size int) *VideoSearchIterator {
	if 1 <= size && size <= 30 {
		it.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Next returns the video search result and proceeds the iterator to the next page.
func (it *VideoSearchIterator) Next() (res VideoSearchResult, err error) {
	if it.end {
		return res, Done
	}

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%svclip?query=%s&sort=%s&page=%d&size=%d",
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

	it.end = res.Meta.IsEnd || 15 < it.Page

	it.Page++

	return
}

// CollectAll collects all the remaining video search results.
func (it *VideoSearchIterator) CollectAll() (results VideoSearchResults) {

	result, err := it.Next()
	if err == nil {
		results = append(results, result)
	}

	n := common.RemainingPages(result.Meta.PageableCount, it.Size, it.Page, 15)

	var (
		items  = make(VideoSearchResults, n)
		errors = make([]error, n)
		wg     sync.WaitGroup
	)

	for page := it.Page; page < it.Page+n; page++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			worker := *it
			items[page-it.Page], errors[page-it.Page] = worker.Result(page).Next()
		}(page)
	}
	wg.Wait()

	for idx, err := range errors {
		if err == nil {
			results = append(results, items[idx])
		}
	}

	it.end = true

	return
}

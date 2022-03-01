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
func (it *DocumentSearchIterator) AuthorizeWith(key string) *DocumentSearchIterator {
	it.AuthKey = common.FormatKey(key)
	return it
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (it *DocumentSearchIterator) SortBy(order string) *DocumentSearchIterator {
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
func (it *DocumentSearchIterator) Result(page int) *DocumentSearchIterator {
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
func (it *DocumentSearchIterator) Display(size int) *DocumentSearchIterator {
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

// Next returns the document search result and proceeds the iterator to the next page.
func (it *DocumentSearchIterator) Next() (res DocumentSearchResult, err error) {
	if it.end {
		return res, Done
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%sweb?query=%s&sort=%s&page=%d&size=%d",
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

// CollectAll collects all the remaining document search results.
func (it *DocumentSearchIterator) CollectAll() (results DocumentSearchResults) {
	result, err := it.Next()
	if err == nil {
		results = append(results, result)
	}

	n := common.RemainingPages(result.Meta.PageableCount, it.Size, it.Page, 50)

	var (
		items  = make(DocumentSearchResults, n)
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

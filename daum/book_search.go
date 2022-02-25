package daum

import (
	"encoding/json"
	"errors"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// BookResult represents a document of a Daum Book search result.
type BookResult struct {
	WebResult
	ISBN        string   `json:"isbn"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	Translators []string `json:"translators"`
	Price       int      `json:"price"`
	SalePrice   int      `json:"sale_price"`
	Thumbnail   string   `json:"thumbnail"`
	Status      string   `json:"status"`
}

// BookSearchResult represents a Daum Book search result.
type BookSearchResult struct {
	Meta      common.PageableMeta `json:"meta"`
	Documents []BookResult        `json:"documents"`
}

// String implements fmt.Stringer.
func (br BookSearchResult) String() string { return common.String(br) }

type BookSearchResults []BookSearchResult

// SaveAs saves brs to @filename.
func (brs BookSearchResults) SaveAs(filename string) error { return common.SaveAsJSON(brs, filename) }

// BookSearchIterator is a lazy book search iterator.
type BookSearchIterator struct {
	Query   string
	AuthKey string
	Sort    string
	Page    int
	Size    int
	Target  string
	end     bool
}

// BookSearch allows to search books by @query in the Daum Book service.
//
// See https://developers.kakao.com/docs/latest/ko/daum-search/dev-guide#search-book for more details.
func BookSearch(query string) *BookSearchIterator {
	return &BookSearchIterator{
		Query:   url.QueryEscape(strings.TrimSpace(query)),
		AuthKey: common.KeyPrefix,
		Sort:    "accuracy",
		Page:    1,
		Size:    10,
		Target:  "",
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (it *BookSearchIterator) AuthorizeWith(key string) *BookSearchIterator {
	it.AuthKey = common.FormatKey(key)
	return it
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or latest (default is accuracy).
func (it *BookSearchIterator) SortBy(order string) *BookSearchIterator {
	switch order {
	case "accuracy", "latest":
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
func (it *BookSearchIterator) Result(page int) *BookSearchIterator {
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
func (it *BookSearchIterator) Display(size int) *BookSearchIterator {
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

// Filter limits the search field.
//
// @target can be one of the following options:
// title, isbn, publisher, person
func (it *BookSearchIterator) Filter(target string) *BookSearchIterator {
	switch target {
	case "title", "isbn", "publisher", "person", "":
		it.Target = target
	default:
		panic(errors.New(
			`target must be one of the following options:
			title, isbn, publisher, person`))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Next returns the book search result and proceeds the iterator to the next page.
func (it *BookSearchIterator) Next() (res BookSearchResult, err error) {
	if it.end {
		return res, ErrEndPage
	}

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v3/search/book?query=%s&sort=%s&page=%d&size=%d&target=%s",
			it.Query, it.Sort, it.Page, it.Size, it.Target), nil)

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
	it.Page++
	it.end = res.Meta.IsEnd || 50 < it.Page

	return
}

// CollectAll collects all the remaining book search results.
func (it *BookSearchIterator) CollectAll() (results BookSearchResults) {
	result, err := it.Next()
	if err == nil {
		results = append(results, result)
	}
	n := common.RemainingPages(result.Meta.PageableCount, it.Size, it.Page, 50)

	var (
		items  = make(BookSearchResults, n)
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

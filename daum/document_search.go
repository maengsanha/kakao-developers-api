package daum

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// WebResult represents a document of a Daum search result.
type WebResult struct {
	Title    string    `json:"title"`
	Contents string    `json:"contents"`
	URL      string    `json:"url"`
	Datetime time.Time `json:"datetime"`
}

// DocumentSearchResult represents a Daum search result.
type DocumentSearchResult struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []WebResult `json:"documents"`
}

// String implements fmt.Stringer.
func (dr DocumentSearchResult) String() string {
	bs, _ := json.MarshalIndent(dr, "", "  ")
	return string(bs)
}

type DocumentSearchResults []DocumentSearchResult

// SaveAs saves drs to @filename.
func (drs DocumentSearchResults) SaveAs(filename string) error {
	if bs, err := json.MarshalIndent(drs, "", "  "); err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, bs, 0644)
	}
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
		AuthKey: "KakaoAK ",
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (di *DocumentSearchIterator) AuthorizeWith(key string) *DocumentSearchIterator {
	di.AuthKey = "KakaoAK " + strings.TrimSpace(key)
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
		panic("order must be either accuracy or recency")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return di
}

// Result sets the result page number (a value between 1 and 50).
func (di *DocumentSearchIterator) Result(page int) *DocumentSearchIterator {
	if 1 <= page && page <= 50 {
		di.Page = page
	} else {
		panic("page must be between 1 and 50")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return di
}

// Display sets the number of documents displayed on a single page (a value between 1 and 50).
func (di *DocumentSearchIterator) Display(size int) *DocumentSearchIterator {
	if 1 <= size && size <= 50 {
		di.Size = size
	} else {
		panic("size must be between 1 and 50")
	}
	if r := recover(); r != nil {
		log.Println(r)
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

	req.Header.Set("Authorization", di.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	di.end = res.Meta.IsEnd

	di.Page++

	return
}

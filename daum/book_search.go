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

// BookResult represents a document of a Daum book search result.
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

// BookSearchResult represents a Daum book search result.
type BookSearchResult struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []BookResult `json:"documents"`
}

// String implements fmt.Stringer.
func (br BookSearchResult) String() string {
	bs, _ := json.MarshalIndent(br, "", "  ")
	return string(bs)
}

type BookSearchResults []BookSearchResult

// SaveAs saves brs to @filename.
func (brs BookSearchResults) SaveAs(filename string) error {
	if bs, err := json.MarshalIndent(brs, "", "  "); err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, bs, 0644)
	}
}

// BookSearchIterator is a lazy blog search iterator.
type BookSearchIterator struct {
	Query   string
	AuthKey string
	Sort    string
	Page    int
	Size    int
	Target  string
	end     bool
}

// BlogSearch allows to search books by @query in the Daum Book service.
//
// See https://developers.kakao.com/docs/latest/ko/daum-search/dev-guide#search-book for more details.
func BookSearch(query string) *BookSearchIterator {
	return &BookSearchIterator{
		Query:   url.QueryEscape(strings.TrimSpace(query)),
		AuthKey: "KakaoAK ",
		Sort:    "accuracy",
		Page:    1,
		Size:    10,
		Target:  "",
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (bi *BookSearchIterator) AuthorizeWith(key string) *BookSearchIterator {
	bi.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return bi
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or latest. (default is accuracy)
func (bi *BookSearchIterator) SortBy(order string) *BookSearchIterator {
	switch order {
	case "accuracy", "latest":
		bi.Sort = order
	default:
		panic("sorting order must be either accuracy or latest")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return bi
}

// Result sets the result page number (a value between 1 and 50).
func (bi *BookSearchIterator) Result(page int) *BookSearchIterator {
	if 1 <= page && page <= 50 {
		bi.Page = page
	} else {
		panic("page must be between 1 and 50")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return bi
}

// Display sets the number of documents displayed on a single page (a value between 1 and 50).
func (bi *BookSearchIterator) Display(size int) *BookSearchIterator {
	if 1 <= size && size <= 50 {
		bi.Size = size
	} else {
		panic("size must be between 1 and 50")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return bi
}

// Filter limits the search field.
//
// @target can be one of the title, isbn, publisher, person.
func (bi *BookSearchIterator) Filter(target string) *BookSearchIterator {
	switch target {
	case "title", "isbn", "publisher", "person", "":
		bi.Target = target
	default:
		panic("filter target must be one of the title, isbn, publisher, person")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return bi
}

// Next returns the book search result and proceeds the iterator to the next page.
func (bi *BookSearchIterator) Next() (res BookSearchResult, err error) {
	if bi.end {
		return res, ErrEndPage
	}

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v3/search/book?query=%s&sort=%s&page=%d&size=%d&target=%s",
			bi.Query, bi.Sort, bi.Page, bi.Size, bi.Target), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", bi.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	bi.Page++

	bi.end = res.Meta.IsEnd || bi.Page > 50

	return
}

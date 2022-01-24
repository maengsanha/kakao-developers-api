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

// BlogResult represents a document of a blog search result.
type BlogResult struct {
	WebResult
	Blogname  string `json:"blogname"`
	Thumbnail string `json:"thumbnail"`
}

// BlogSearchResult represents a blog search result.
type BlogSearchResult struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []BlogResult `json:"documents"`
}

// String implements fmt.Stringer.
func (br BlogSearchResult) String() string {
	bs, _ := json.MarshalIndent(br, "", "  ")
	return string(bs)
}

type BlogSearchResults []BlogSearchResult

// SaveAs saves brs to @filename.
func (brs BlogSearchResults) SaveAs(filename string) error {
	if bs, err := json.MarshalIndent(brs, "", "  "); err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, bs, 0644)
	}
}

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
		AuthKey: "KakaoAK ",
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (bi *BlogSearchIterator) AuthorizeWith(key string) *BlogSearchIterator {
	bi.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return bi
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (bi *BlogSearchIterator) SortBy(order string) *BlogSearchIterator {
	switch order {
	case "accuracy", "recency":
		bi.Sort = order
	default:
		panic("order must be either accuracy or recency")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return bi
}

// Result sets the result page number (a value between 1 and 50).
func (bi *BlogSearchIterator) Result(page int) *BlogSearchIterator {
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
func (bi *BlogSearchIterator) Display(size int) *BlogSearchIterator {
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

// Next returns the blog search result and proceeds the iterator to the next page.
func (bi *BlogSearchIterator) Next() (res BlogSearchResult, err error) {
	if bi.end {
		return res, ErrEndPage
	}

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%sblog?query=%s&sort=%s&page=%d&size=%d",
			prefix, bi.Query, bi.Sort, bi.Page, bi.Size), nil)

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

	bi.end = res.Meta.IsEnd

	bi.Page++

	return
}

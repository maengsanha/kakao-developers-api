package daum

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CafeResult represents a document of cafe search result.
type CafeResult struct {
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	URL       string    `json:"url"`
	CafeName  string    `json:"cafename"`
	Thumbnail string    `json:"thumbnail"`
	Datetime  time.Time `json:"datetime"`
}

// CafeSearchResult represents a cafe search result.
type CafeSearchResult struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []CafeResult `json:"documents"`
}

// String implements fmt.Stringer.
func (ci CafeSearchResult) String() string {
	bs, _ := Indent(ci, "", "  ")
	return string(bs)
}

// JSONEncode encodes JSON format to []byte format.
func JSONEncode(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

// Indent appends to an indented form of the JSON-encoded src.
func Indent(v interface{}, prefix, indent string) ([]byte, error) {
	b, err := JSONEncode(v)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, prefix, indent)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type CafeSearchResults []CafeSearchResult

// SaveAs saves crs to @filename.
func (crs CafeSearchResults) SaveAs(filename string) error {
	switch tokens := strings.Split(filename, "."); tokens[len(tokens)-1] {
	case "json":
		if bs, err := json.MarshalIndent(crs, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	default:
		return errors.New("file format must be json")
	}
}

// CafeSearchIterator is a lazy Cafe search iterator.
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
		AuthKey: "KakaoAK ",
		Sort:    "accuracy",
		Page:    1,
		Size:    10,
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ci *CafeSearchIterator) AuthorizeWith(key string) *CafeSearchIterator {
	ci.AuthKey = "KakaoAK " + strings.TrimSpace(key)
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
		panic(errors.New("sorting order must be either accuracy or recency"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Result sets the result page number (a value between 1 and 50).
func (ci *CafeSearchIterator) Result(page int) *CafeSearchIterator {
	if 1 <= page && page <= 50 {
		ci.Page = page
	} else {
		panic(errors.New("page must be between 1 and 50"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Display sets the number of documents displayed on a single page (a value between 1 and 50).
func (ci *CafeSearchIterator) Display(size int) *CafeSearchIterator {
	if 1 <= size && size <= 50 {
		ci.Size = size
	} else {
		panic(errors.New("size musts be between 1 and 50"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// Next returns the cafe search result and proceeds the iterator to the next page.
func (ci *CafeSearchIterator) Next() (res CafeSearchResult, err error) {
	if ci.end {
		return res, errors.New("page reaches the end")
	}

	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%scafe?query=%s&sort=%s&page=%d&size=%d",
			prefix, ci.Query, ci.Sort, ci.Page, ci.Size), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", ci.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	ci.end = res.Meta.IsEnd

	ci.Page++

	return
}

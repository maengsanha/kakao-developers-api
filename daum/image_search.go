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

// ImageResult represents a result of a image search result.
type ImageResult struct {
	Collection      string    `json:"collection"`
	ThumbnailURL    string    `json:"thumbnail_url"`
	ImageURL        string    `json:"image_url"`
	Width           int       `json:"width"`
	Height          int       `json:"height"`
	DisplaySitename string    `json:"display_sitename"`
	DocURL          string    `json:"doc_url"`
	Datetime        time.Time `json:"datetime"`
}

// ImageSearchResult represents a image search result.
type ImageSearchResult struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []ImageResult `json:"documents"`
}

// String implements fmt.Stringer.
func (ir ImageSearchResult) String() string {
	bs, _ := json.MarshalIndent(ir, "", "  ")
	return string(bs)
}

type ImageSearchResults []ImageSearchResult

// SaveAs saves irs to @filename.
func (irs ImageSearchResults) SaveAs(filename string) error {
	if bs, err := json.MarshalIndent(irs, "", "  "); err != nil {
		return err
	} else {
		return ioutil.WriteFile(filename, bs, 0644)
	}
}

// ImageSearchIterator is a lazy image search iterator.
type ImageSearchIterator struct {
	Query   string
	Sort    string
	Page    int
	Size    int
	AuthKey string
	end     bool
}

// ImageSearch allows users to search images by @query in the Daum Search service.
//
// For more details visit https://developers.kakao.com/docs/latest/en/daum-search/dev-guide#search-image.
func ImageSearch(query string) *ImageSearchIterator {
	return &ImageSearchIterator{
		Query:   url.QueryEscape(strings.TrimSpace(query)),
		Sort:    "accuracy",
		Page:    1,
		Size:    80,
		AuthKey: "KakaoAK ",
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ii *ImageSearchIterator) AuthorizeWith(key string) *ImageSearchIterator {
	ii.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return ii
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (ii *ImageSearchIterator) SortBy(order string) *ImageSearchIterator {
	switch order {
	case "accuracy", "recency":
		ii.Sort = order
	default:
		panic("order must be either accuracy or recency")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ii
}

// Result sets the result page number (a value between 1 and 50).
func (ii *ImageSearchIterator) Result(page int) *ImageSearchIterator {
	if 1 <= page && page <= 50 {
		ii.Page = page
	} else {
		panic("page must be between 1 and 50")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ii
}

// Display sets the number of documents displayed on a single page (a value between 1 and 80).
func (ii *ImageSearchIterator) Display(size int) *ImageSearchIterator {
	if 1 <= size && size <= 80 {
		ii.Size = size
	} else {
		panic("size must be between 1 and 80")
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ii
}

// Next returns the image search result and proceeds the iterator to the next page.
func (ii *ImageSearchIterator) Next() (res ImageSearchResult, err error) {
	if ii.end {
		return res, ErrEndPage
	}

	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%simage?query=%s&sort=%s&page=%d&size=%d",
			prefix, ii.Query, ii.Sort, ii.Page, ii.Size), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", ii.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	ii.end = res.Meta.IsEnd || ii.Page > 50

	ii.Page++

	return
}

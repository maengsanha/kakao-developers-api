// Copyright 2022 Sanha Maeng, Soyang Baek, Jinmyeong Kim
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package daum

import (
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/goccy/go-json"
)

// ImageResult represents a document of an image search result.
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

// ImageSearchResult represents an image search result.
type ImageSearchResult struct {
	Meta      common.PageableMeta `json:"meta"`
	Documents []ImageResult       `json:"documents"`
}

// String implements fmt.Stringer.
func (ir ImageSearchResult) String() string { return common.String(ir) }

type ImageSearchResults []ImageSearchResult

// SaveAs saves irs to @filename.
func (irs ImageSearchResults) SaveAs(filename string) error { return common.SaveAsJSON(irs, filename) }

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
		AuthKey: common.KeyPrefix,
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (it *ImageSearchIterator) AuthorizeWith(key string) *ImageSearchIterator {
	it.AuthKey = common.FormatKey(key)
	return it
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (it *ImageSearchIterator) SortBy(order string) *ImageSearchIterator {
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
func (it *ImageSearchIterator) Result(page int) *ImageSearchIterator {
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

// Display sets the number of documents displayed on a single page (a value between 1 and 80).
func (it *ImageSearchIterator) Display(size int) *ImageSearchIterator {
	if 1 <= size && size <= 80 {
		it.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Next returns the image search result and proceeds the iterator to the next page.
func (it *ImageSearchIterator) Next() (res ImageSearchResult, err error) {
	if it.end {
		return res, Done
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%simage?query=%s&sort=%s&page=%d&size=%d",
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

// CollectAll collects all the remaining image search results.
func (it *ImageSearchIterator) CollectAll() (results ImageSearchResults) {

	result, err := it.Next()
	if err == nil {
		results = append(results, result)
	}

	n := common.RemainingPages(result.Meta.PageableCount, it.Size, it.Page, 50)

	var (
		items  = make(ImageSearchResults, n)
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

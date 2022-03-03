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
	"encoding/json"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// BlogResult represents a document of a blog search result.
type BlogResult struct {
	WebResult
	Blogname  string `json:"blogname"`
	Thumbnail string `json:"thumbnail"`
}

// BlogSearchResult represents a blog search result.
type BlogSearchResult struct {
	Meta      common.PageableMeta `json:"meta"`
	Documents []BlogResult        `json:"documents"`
}

// String implements fmt.Stringer.
func (br BlogSearchResult) String() string { return common.String(br) }

type BlogSearchResults []BlogSearchResult

// SaveAs saves brs to @filename.
func (brs BlogSearchResults) SaveAs(filename string) error { return common.SaveAsJSON(brs, filename) }

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
		AuthKey: common.KeyPrefix,
		end:     false,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (it *BlogSearchIterator) AuthorizeWith(key string) *BlogSearchIterator {
	it.AuthKey = common.FormatKey(key)
	return it
}

// SortBy sets the sorting order of the document results to @order.
//
// @order can be accuracy or recency. (default is accuracy)
func (it *BlogSearchIterator) SortBy(order string) *BlogSearchIterator {
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
func (it *BlogSearchIterator) Result(page int) *BlogSearchIterator {
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
func (it *BlogSearchIterator) Display(size int) *BlogSearchIterator {
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

// Next returns the blog search result and proceeds the iterator to the next page.
func (it *BlogSearchIterator) Next() (res BlogSearchResult, err error) {
	if it.end {
		return res, Done
	}

	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%sblog?query=%s&sort=%s&page=%d&size=%d",
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

// CollectAll collects all the remaining blog search results.
func (it *BlogSearchIterator) CollectAll() (results BlogSearchResults) {
	result, err := it.Next()
	if err == nil {
		results = append(results, result)
	}

	n := common.RemainingPages(result.Meta.PageableCount, it.Size, it.Page, 50)

	var (
		items  = make(BlogSearchResults, n)
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

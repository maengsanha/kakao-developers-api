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

package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// CategorySearchIterator is a lazy category search iterator.
type CategorySearchIterator struct {
	Query             string
	Format            string
	AuthKey           string
	CategoryGroupCode string
	X                 string
	Y                 string
	Radius            int
	Rect              string
	Page              int
	Size              int
	Sort              string
	end               bool
}

// PlaceSearchByCategory provides the search results for place by group code in the specified order.
//
// This api provides two search options.
//
// 1. Search with @x, @y coordinates and @radius that distance from @x and @y.
//
// 2. Search with @rect that coordinates of left X, left Y, right X, right Y.
//
// There are a few available category codes:
//
// MT1	Large supermarket
//
// CS2	Convenience store
//
// PS3	Kindergarten
//
// SC4	School
//
// AC5	Academic
//
// PK6	Parking
//
// OL7	Gas station, Charging Station
//
// SW8	Subway station
//
// BK9	Bank
//
// CT1	Cultural facility
//
// AG2	Brokerage
//
// PO3	Public institution
//
// AT4	Tourist attraction
//
// AD5	Accommodation
//
// FD6	Restaurant
//
// CE7	Cafe
//
// HP8	Hospital
//
// PM9	Pharmacys
//
// Details can be referred to
// https://developers.kakao.com/docs/latest/en/local/dev-guide#search-by-category.
func PlaceSearchByCategory(groupcode string) *CategorySearchIterator {
	switch groupcode {
	case "MT1", "CS2", "PS3", "SC4", "AC5", "PK6", "OL7", "SW8", "BK9",
		"CT1", "AG2", "PO3", "AT4", "AD5", "FD6", "CE7", "HP8", "PM9":
	default:
		panic(ErrUnsupportedCategoryGroupCode)
	}

	if r := recover(); r != nil {
		log.Panicln(r)
		return nil
	}

	return &CategorySearchIterator{
		Format:            "json",
		AuthKey:           common.KeyPrefix,
		CategoryGroupCode: groupcode,
		X:                 "",
		Y:                 "",
		Radius:            0,
		Rect:              "",
		Page:              1,
		Size:              15,
		Sort:              "accuracy",
	}
}

// FormatAs sets the request format to @format (json or xml).
func (it *CategorySearchIterator) FormatAs(format string) *CategorySearchIterator {
	switch format {
	case "json", "xml":
		it.Format = format
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Authorization sets the authorization key to @key.
func (it *CategorySearchIterator) AuthorizeWith(key string) *CategorySearchIterator {
	it.AuthKey = common.FormatKey(key)
	return it
}

// WithRadius searches places around a specific area along with @x and @y.
//
// @radius is the distance (a value between 0 and 20000) from the center coordinates to an axis of rotation in meters.
func (it *CategorySearchIterator) WithRadius(x, y float64, radius int) *CategorySearchIterator {
	if 0 <= it.Radius && it.Radius <= 20000 {
		it.X = strconv.FormatFloat(x, 'f', -1, 64)
		it.Y = strconv.FormatFloat(y, 'f', -1, 64)
		it.Radius = radius
	} else {
		panic(ErrRadiusOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// WithRect limits the search area, such as when searching places within the map screen.
func (it *CategorySearchIterator) WithRect(xMin, yMin, xMax, yMax float64) *CategorySearchIterator {
	it.Rect = strings.Join([]string{
		strconv.FormatFloat(xMin, 'f', -1, 64),
		strconv.FormatFloat(yMin, 'f', -1, 64),
		strconv.FormatFloat(xMax, 'f', -1, 64),
		strconv.FormatFloat(yMax, 'f', -1, 64)}, ",")
	return it
}

// Result sets the result page number (a value between 1 and 45).
func (it *CategorySearchIterator) Result(page int) *CategorySearchIterator {
	if 1 <= page && page <= 45 {
		it.Page = page
	} else {
		panic(common.ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Display sets the number of documents displayed on a single page (a value between 1 and 15).
func (it *CategorySearchIterator) Display(size int) *CategorySearchIterator {
	if 1 <= size && size <= 15 {
		it.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// SortBy sets the ordering type of c to @order.
//
// @order can be accuracy or distance. (default is accuracy)
func (it *CategorySearchIterator) SortBy(order string) *CategorySearchIterator {
	switch order {
	case "accuracy", "distance":
		it.Sort = order
	default:
		panic(common.ErrUnsupportedSortingOrder)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Next returns the place search result.
func (it *CategorySearchIterator) Next() (res PlaceSearchResult, err error) {
	if it.end {
		return res, Done
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%ssearch/category.%s?category_group_code=%s&page=%d&size=%d&sort=%s&x=%s&y=%s&radius=%d&rect=%s",
			prefix, it.Format, it.CategoryGroupCode, it.Page, it.Size, it.Sort, it.X, it.Y, it.Radius, it.Rect), nil)

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

	if it.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if it.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	it.end = res.Meta.IsEnd || 45 < it.Page

	it.Page++

	return
}

// CollectAll collects all the remaining category search results.
func (it *CategorySearchIterator) CollectAll() (results PlaceSearchResults) {

	result, err := it.Next()
	if err == nil {
		results = append(results, result)
	}

	n := common.RemainingPages(result.Meta.PageableCount, it.Size, it.Page, 45)

	var (
		items  = make(PlaceSearchResults, n)
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

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

package local_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestKeywordSearchWithJSON(t *testing.T) {
	query := "카카오"
	groupcode := "PK6"
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 10000
	order := "accuracy"

	it := local.PlaceSearchByKeyword(query).
		FormatAs("json").
		AuthorizeWith(common.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order)

	for {
		item, err := it.Next()
		if err == local.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		t.Log(item)
	}
}

func TestKeywordSearchWithSaveAsJSON(t *testing.T) {
	query := "카카오"
	groupcode := "PK6"
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 10000
	order := "accuracy"

	it := local.PlaceSearchByKeyword(query).
		FormatAs("json").
		AuthorizeWith(common.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order)

	items := local.PlaceSearchResults{}

	for {
		item, err := it.Next()
		if err == local.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		items = append(items, item)
	}
	if err := items.SaveAs("keyword_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestKeywordSearchWithXML(t *testing.T) {
	query := "카카오"
	groupcode := ""
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 15000
	order := "distance"
	xMin := 126.92839423213
	yMin := 37.412341512321
	xMax := 126.943241321321
	yMax := 37.5904321012312

	it := local.PlaceSearchByKeyword(query).
		FormatAs("xml").
		AuthorizeWith(common.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		WithRect(xMin, yMin, xMax, yMax).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order)

	for {
		item, err := it.Next()
		if err == local.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		t.Log(item)
	}
}

func TestKeywordSearchWithSaveAsXML(t *testing.T) {
	query := "카카오"
	groupcode := ""
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 15000
	order := "distance"
	xMin := 126.92839423213
	yMin := 37.412341512321
	xMax := 126.943241321321
	yMax := 37.5904321012312

	iter := local.PlaceSearchByKeyword(query).
		FormatAs("xml").
		AuthorizeWith(common.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		WithRect(xMin, yMin, xMax, yMax).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order)

	items := local.PlaceSearchResults{}

	for {
		item, err := iter.Next()
		if err == local.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		items = append(items, item)
	}

	if err := items.SaveAs("keyword_search_test.xml"); err != nil {
		t.Error(err)
	}

}

func TestKeywordSearchCollectAll(t *testing.T) {
	query := "카카오"
	groupcode := "PK6"
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 10000
	order := "accuracy"

	items := local.PlaceSearchByKeyword(query).
		FormatAs("json").
		AuthorizeWith(common.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order).
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

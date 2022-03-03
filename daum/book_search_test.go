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

package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestBookSearchWithJSON(t *testing.T) {
	query := "히가시노 게이고"

	it := daum.BookSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("latest").
		Result(1).
		Display(10).
		Filter("person")

	for {
		item, err := it.Next()
		if err == daum.Done {
			break
		}
		if err != nil {
			t.Error(err)

		}
		t.Log(item)
	}

}

func TestBookSearchWithSaveAsJSON(t *testing.T) {
	query := "히가시노 게이고"

	it := daum.BookSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("latest").
		Result(1).
		Display(10).
		Filter("person")

	items := daum.BookSearchResults{}

	for {
		item, err := it.Next()
		if err == daum.Done {
			break
		}

		if err != nil {
			t.Error(err)
			break
		}
		items = append(items, item)
	}
	if err := items.SaveAs("book_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestBookSearchCollectAll(t *testing.T) {
	query := "히가시노 게이고"

	items := daum.BookSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("latest").
		Result(1).
		Display(10).
		Filter("person").
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

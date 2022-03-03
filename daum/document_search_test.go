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

func TestDocumentSearchWithJSON(t *testing.T) {
	query := "Alan Turing"

	it := daum.DocumentSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Result(10).
		Display(50)

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

func TestDocumentSearchWithSaveAsJSON(t *testing.T) {
	query := "Alan Turing"

	it := daum.DocumentSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Result(1).
		Display(30)

	items := daum.DocumentSearchResults{}

	for {
		item, err := it.Next()
		if err == daum.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		items = append(items, item)
	}

	if err := items.SaveAs("document_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestDocumentSearchCollectAll(t *testing.T) {
	query := "Alan Turing"

	items := daum.DocumentSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Result(1).
		Display(50).
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

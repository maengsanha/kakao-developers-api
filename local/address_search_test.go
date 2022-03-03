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

func TestAddressSearchWithJSON(t *testing.T) {
	query := "을지로"

	it := local.AddressSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		Analyze("similar").
		FormatAs("json").
		Display(20).
		Result(1)

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

func TestAddressSearchWithSaveAsJSON(t *testing.T) {
	query := "을지로"

	it := local.AddressSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		Analyze("similar").
		FormatAs("json").
		Display(20).
		Result(1)

	items := local.AddressSearchResults{}

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

	if err := items.SaveAs("address_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestAddressSearchWithXML(t *testing.T) {
	query := "을지로"

	it := local.AddressSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		Analyze("similar").
		FormatAs("xml").
		Display(30).
		Result(1)

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

func TestAddressSearchWithSaveAsXML(t *testing.T) {
	query := "을지로"

	it := local.AddressSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		Analyze("similar").
		FormatAs("xml").
		Display(30).
		Result(1)

	items := local.AddressSearchResults{}

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

	if err := items.SaveAs("address_search_test.xml"); err != nil {
		t.Error(err)
	}
}

func TestAddressSearchCollectAll(t *testing.T) {
	query := "을지로"

	items := local.AddressSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		Analyze("similar").
		FormatAs("json").
		Display(30).
		Result(1).
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

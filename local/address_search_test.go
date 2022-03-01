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

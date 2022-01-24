package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestAddressSearchWithJSON(t *testing.T) {
	query := "을지로"

	iter := local.AddressSearch(query).
		AuthorizeWith(local.REST_API_KEY).
		Analyze("similar").
		FormatAs("json").
		Display(20).
		Result(1)

	for ar, err := iter.Next(); err == nil; ar, err = iter.Next() {
		t.Log(ar)
	}
}

func TestAddressSearchWithSaveAsJSON(t *testing.T) {
	query := "을지로"

	iter := local.AddressSearch(query).
		AuthorizeWith(local.REST_API_KEY).
		Analyze("similar").
		FormatAs("json").
		Display(20).
		Result(1)

	ars := local.AddressSearchResults{}

	for ar, err := iter.Next(); err == nil; ar, err = iter.Next() {
		ars = append(ars, ar)
	}

	if err := ars.SaveAs("address_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestAddressSearchWithXML(t *testing.T) {
	query := "을지로"

	iter := local.AddressSearch(query).
		AuthorizeWith(local.REST_API_KEY).
		Analyze("similar").
		FormatAs("xml").
		Display(30).
		Result(1)

	for ar, err := iter.Next(); err == nil; ar, err = iter.Next() {
		t.Log(ar)
	}
}

func TestAddressSearchWithSaveAsXML(t *testing.T) {
	query := "을지로"

	iter := local.AddressSearch(query).
		AuthorizeWith(local.REST_API_KEY).
		Analyze("similar").
		FormatAs("xml").
		Display(30).
		Result(1)

	ars := local.AddressSearchResults{}

	for ar, err := iter.Next(); err == nil; ar, err = iter.Next() {
		ars = append(ars, ar)
	}

	if err := ars.SaveAs("address_search_test.xml"); err != nil {
		t.Error(err)
	}
}

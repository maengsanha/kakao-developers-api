package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/local"
)

func TestAddressSearchWithJSON(t *testing.T) {
	query := "을지로"
	key := ""

	iter := local.AddressSearch(query).
		AuthorizeWith(key).
		Analyze("similar").
		FormatAs("json").
		Display(20).
		Result(1)

	for res, err := iter.Next(); err == nil; res, err = iter.Next() {
		t.Log(res)
	}
}

func TestAddressSearchWithXML(t *testing.T) {
	query := "을지로"
	key := ""

	iter := local.AddressSearch(query).
		AuthorizeWith(key).
		Analyze("similar").
		FormatAs("xml").
		Display(30).
		Result(1)

	for res, err := iter.Next(); err == nil; res, err = iter.Next() {
		t.Log(res)
	}
}

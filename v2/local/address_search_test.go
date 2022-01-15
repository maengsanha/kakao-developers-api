package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestAddressSearchWithJSON(t *testing.T) {
	query := "성북구 정릉동"
	key := ""

	iter := local.AddressSearch(query).
		AuthorizeWith(key).
		Analyze("similar").
		FormatJSON().
		Display(20).
		Result(1)

	for res, err := iter.Next(); ; res, err = iter.Next() {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

func TestAddressSearchWithXML(t *testing.T) {
	query := "익선동"
	key := ""

	iter := local.AddressSearch(query).
		AuthorizeWith(key).
		Analyze("similar").
		FormatXML().
		Display(30).
		Result(1)

	for res, err := iter.Next(); ; res, err = iter.Next() {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

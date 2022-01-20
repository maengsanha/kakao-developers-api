package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestCoord2AddressWithJSON(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if res, err := local.CoordToAddress(x, y).
		AuthorizeWith(local.REST_API_KEY).
		Input(coord).
		FormatJSON().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}

}

func TestCoord2AddressWithXML(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if res, err := local.CoordToAddress(x, y).
		AuthorizeWith(local.REST_API_KEY).
		Input(coord).
		FormatXML().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}

}

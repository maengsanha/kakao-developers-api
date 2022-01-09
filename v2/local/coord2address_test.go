package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestCoord2AddressWithJSON(t *testing.T) {
	x := "127.1086228"
	y := "37.4012191"
	key := ""

	if res, err := local.CoordToAddress(x, y).
		AuthorizeWith(key).
		RequestWGS84().
		FormatJSON().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}

}

func TestCoord2AddressWithXML(t *testing.T) {
	x := "127.1086228"
	y := "37.4012191"
	key := ""

	if res, err := local.CoordToAddress(x, y).
		AuthorizeWith(key).
		RequestWGS84().
		FormatXML().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}

}

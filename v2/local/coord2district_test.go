package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestCoord2DistrictWithJSON(t *testing.T) {
	x := "127.1086228"
	y := "37.4012191"
	key := ""

	if res, err := local.Coord2District(x, y).
		AuthorizeWith(key).
		RequestWGS84().
		DisplayWGS84().
		FormatJSON().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestCoord2DistrictWithXML(t *testing.T) {
	x := "127.1086228"
	y := "37.4012191"
	key := ""

	if res, err := local.Coord2District(x, y).
		AuthorizeWith(key).
		RequestWGS84().
		DisplayCONGNAMUL().
		FormatXML().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

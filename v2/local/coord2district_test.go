package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestCoordToDistrictWithJSON(t *testing.T) {
	x := "127.1086228"
	y := "37.4012191"
	key := ""

	if res, err := local.CoordToDistrict(x, y).
		AuthorizeWith(key).
		Input("WGS84").
		Output("WGS84").
		FormatJSON().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestCoordToDistrictWithXML(t *testing.T) {
	x := "127.1086228"
	y := "37.4012191"
	key := ""

	if res, err := local.CoordToDistrict(x, y).
		AuthorizeWith(key).
		Input("WGS84").
		Output("CONGNAMUL").
		FormatXML().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestCoordToDistrictWithJSON(t *testing.T) {
	x := 127.1086228
	y := 37.4012191
	key := ""

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(key).
		Input("WGS84").
		Output("WGS84").
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2district_test.json"); err != nil {
		t.Error(err)
	}
}

func TestCoordToDistrictWithXML(t *testing.T) {
	x := 127.1086228
	y := 37.4012191
	key := ""

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(key).
		Input("WGS84").
		Output("CONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2district_test.xml"); err != nil {
		t.Error(err)
	}
}

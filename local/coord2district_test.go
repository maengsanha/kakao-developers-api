package local_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestCoordToDistrictWithJSON(t *testing.T) {
	x := 127.1086228
	y := 37.4012191

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WGS84").
		Output("WGS84").
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}
}

func TestCoordToDistrictWithSaveAsJSON(t *testing.T) {
	x := 127.1086228
	y := 37.4012191

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(common.REST_API_KEY).
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

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WGS84").
		Output("CONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}
}

func TestCoordToDistrictWithSaveAsXML(t *testing.T) {
	x := 127.1086228
	y := 37.4012191

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WGS84").
		Output("CONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2district_test.xml"); err != nil {
		t.Error(err)
	}
}

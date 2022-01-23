package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestCoord2AddressWithJSON(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if cr, err := local.CoordToAddress(x, y).
		AuthorizeWith(local.REST_API_KEY).
		Input(coord).
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}

}

func TestCoord2AddressWithSaveAsJSON(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if cr, err := local.CoordToAddress(x, y).
		AuthorizeWith(local.REST_API_KEY).
		Input(coord).
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2address_test.json"); err != nil {
		t.Error(err)
	}

}

func TestCoord2AddressWithXML(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if cr, err := local.CoordToAddress(x, y).
		AuthorizeWith(local.REST_API_KEY).
		Input(coord).
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}

}

func TestCoord2AddressWithSaveAsXML(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if cr, err := local.CoordToAddress(x, y).
		AuthorizeWith(local.REST_API_KEY).
		Input(coord).
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2address_test.xml"); err != nil {
		t.Error(cr)
	}

}

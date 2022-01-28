package local_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestTransCoordWithJSON(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if tr, err := local.TransCoord(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestTransCoordWithSaveAsJSON(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if tr, err := local.TransCoord(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else if err = tr.SaveAs("transcoord_test.json"); err != nil {
		t.Error(err)
	}
}

func TestTransCoordWithXML(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if tr, err := local.TransCoord(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestTransCoordWithSaveAsXML(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if tr, err := local.TransCoord(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else if err = tr.SaveAs("transcoord_test.xml"); err != nil {
		t.Error(err)
	}
}

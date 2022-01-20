package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestTransCoordWithJSON(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if res, err := local.TransCoord(x, y).
		AuthorizeWith(local.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatJSON().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestTransCoordWithXML(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if res, err := local.TransCoord(x, y).
		AuthorizeWith(local.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatXML().
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

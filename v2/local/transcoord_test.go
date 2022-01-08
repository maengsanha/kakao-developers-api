package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestTransCoord(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299
	key := ""

	if res, err := local.TransCoord(x, y).
		AuthorizeWith(key).
		Request("WTM").
		Display("WGS84").
		As("JSON").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

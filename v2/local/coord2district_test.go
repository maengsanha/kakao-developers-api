package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestCoord2District(t *testing.T) {
	x := "127.1086228"
	y := "37.4012191"
	key := ""

	if res, err := local.Coord2District(x, y).
		AuthorizeWith(key).
		Request(local.WGS84).
		Display(local.WGS84).
		As(local.XML).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

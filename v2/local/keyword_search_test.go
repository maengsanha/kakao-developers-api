package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestKeyWordSearch(t *testing.T) {
	query := "카카오프렌즈"
	key := ""
	category_group_code := ""
	x := "127.06283102249932"
	y := "37.514322572335935"
	radius := 20000
	sort := "accuracy"
	rect := ""
	if res, err := local.KeyWordSearch(query, x, y, category_group_code, rect).
		As(local.JSON).
		AuthorizeWith(key).
		SetRadius(radius).
		Result(local.DefaultPage).
		Display(local.DefaultSize).
		SortType(sort).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}

}

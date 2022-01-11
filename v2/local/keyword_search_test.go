package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestKeywordSearchWithJSON(t *testing.T) {
	query := "카카오프렌즈"
	key := ""
	category_group_code := ""
	x := "127.06283102249932"
	y := "37.514322572335935"
	radius := 20000
	sort := "accuracy"
	rect := ""
	if res, err := local.KeywordSearch(query).
		As("json").
		AuthorizeWith(key).
		SetRadius(radius).
		Result(1).
		Display(15).
		SetCategoryGroupCode(category_group_code).
		SetX(x).
		SetY(y).
		SetRect(rect).
		SortType(sort).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}

}

func TestKeywordSearchWithXML(t *testing.T) {
	query := "카카오프렌즈"
	key := ""
	category_group_code := ""
	x := "127.06283102249932"
	y := "37.514322572335935"
	radius := 20000
	sort := "accuracy"
	rect := ""
	if res, err := local.KeywordSearch(query).
		As("xml").
		AuthorizeWith(key).
		SetRadius(radius).
		Result(1).
		Display(15).
		SetCategoryGroupCode(category_group_code).
		SetX(x).
		SetY(y).
		SetRect(rect).
		SortType(sort).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}

}

package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestKeywordRadiusSearchWithJSON(t *testing.T) {
	query := "카카오프렌즈"
	key := ""
	groupcode := ""
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 20000
	typ := "accuracy"

	iter := local.KeywordSearch(query).
		FormatJSON().
		AuthorizeWith(key).
		WithRadius(x, y, radius).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(typ)

	for res, err := iter.Next(); ; res, err = iter.Next() {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

func TestKeywordRadiusSearchWithXML(t *testing.T) {
	query := "카카오프렌즈"
	key := ""
	groupcode := ""
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 20000
	typ := "accuracy"

	iter := local.KeywordSearch(query).
		FormatXML().
		AuthorizeWith(key).
		WithRadius(x, y, radius).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(typ)

	for res, err := iter.Next(); ; res, err = iter.Next() {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

func TestKeywordRectSearchWithJSON(t *testing.T) {
	query := "카카오"
	key := ""
	groupcode := ""
	xMin := 126.92839423213
	yMin := 37.412341512321
	xMax := 126.943241321321
	yMax := 37.5904321012312
	typ := "accuracy"

	iter := local.KeywordSearch(query).
		FormatJSON().
		AuthorizeWith(key).
		WithRect(xMin, yMin, xMax, yMax).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(typ)

	for res, err := iter.Next(); ; res, err = iter.Next() {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

func TestKeywordRectSearchWithXML(t *testing.T) {
	query := "카카오"
	key := ""
	groupcode := ""
	xMin := 126.92839423213
	yMin := 37.412341512321
	xMax := 126.943241321321
	yMax := 37.5904321012312
	typ := "accuracy"

	iter := local.KeywordSearch(query).
		FormatXML().
		AuthorizeWith(key).
		WithRect(xMin, yMin, xMax, yMax).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(typ)

	for res, err := iter.Next(); ; res, err = iter.Next() {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

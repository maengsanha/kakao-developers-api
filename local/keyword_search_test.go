package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/local"
)

func TestKeywordSearchWithJSON(t *testing.T) {
	query := "카카오"
	key := ""
	groupcode := "PK6"
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 10000
	typ := "accuracy"

	iter := local.PlaceSearchByKeyword(query).
		FormatJSON().
		AuthorizeWith(key).
		WithCoordinates(x, y).
		WithRadius(radius).
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

func TestKeywordSearchWithXML(t *testing.T) {
	query := "카카오"
	key := ""
	groupcode := ""
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 15000
	typ := "distance"
	xMin := 126.92839423213
	yMin := 37.412341512321
	xMax := 126.943241321321
	yMax := 37.5904321012312

	iter := local.PlaceSearchByKeyword(query).
		FormatXML().
		AuthorizeWith(key).
		WithCoordinates(x, y).
		WithRadius(radius).
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

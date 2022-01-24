package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestKeywordSearchWithJSON(t *testing.T) {
	query := "카카오"
	groupcode := "PK6"
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 10000
	order := "accuracy"

	iter := local.PlaceSearchByKeyword(query).
		FormatAs("json").
		AuthorizeWith(local.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order)

	for pr, err := iter.Next(); err == nil; pr, err = iter.Next() {
		t.Log(pr)
	}

}

func TestKeywordSearchWithSaveAsJSON(t *testing.T) {
	query := "카카오"
	groupcode := "PK6"
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 10000
	order := "accuracy"

	iter := local.PlaceSearchByKeyword(query).
		FormatAs("json").
		AuthorizeWith(local.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order)

	prs := local.PlaceSearchResults{}

	for pr, err := iter.Next(); err == nil; pr, err = iter.Next() {
		prs = append(prs, pr)
	}

	if err := prs.SaveAs("keyword_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestKeywordSearchWithXML(t *testing.T) {
	query := "카카오"
	groupcode := ""
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 15000
	order := "distance"
	xMin := 126.92839423213
	yMin := 37.412341512321
	xMax := 126.943241321321
	yMax := 37.5904321012312

	iter := local.PlaceSearchByKeyword(query).
		FormatAs("xml").
		AuthorizeWith(local.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		WithRect(xMin, yMin, xMax, yMax).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order)

	for pr, err := iter.Next(); err == nil; pr, err = iter.Next() {
		t.Log(pr)
	}

}
func TestKeywordSearchWithSaveAsXML(t *testing.T) {
	query := "카카오"
	groupcode := ""
	x := 127.06283102249932
	y := 37.514322572335935
	radius := 15000
	order := "distance"
	xMin := 126.92839423213
	yMin := 37.412341512321
	xMax := 126.943241321321
	yMax := 37.5904321012312

	iter := local.PlaceSearchByKeyword(query).
		FormatAs("xml").
		AuthorizeWith(local.REST_API_KEY).
		WithCoordinates(x, y).
		WithRadius(radius).
		WithRect(xMin, yMin, xMax, yMax).
		Result(1).
		Display(15).
		Category(groupcode).
		SortBy(order)

	prs := local.PlaceSearchResults{}

	for pr, err := iter.Next(); err == nil; pr, err = iter.Next() {
		prs = append(prs, pr)
	}

	if err := prs.SaveAs("keyword_search_test.xml"); err != nil {
		t.Error(err)
	}

}

package local_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestCategorySearchWithJSON(t *testing.T) {
	var x float64 = 127.06283102249932
	var y float64 = 37.514322572335935
	radius := 2000
	groupcode := "MT1"

	iter := local.PlaceSearchByCategory(groupcode).
		FormatAs("json").
		AuthorizeWith(common.REST_API_KEY).
		WithRadius(x, y, radius).
		Display(15).
		Result(1)

	for pr, err := iter.Next(); err == nil; pr, err = iter.Next() {
		t.Log(pr)
	}
}

func TestCategorySearchWithSaveAsJSON(t *testing.T) {
	var x float64 = 127.06283102249932
	var y float64 = 37.514322572335935
	radius := 2000
	groupcode := "MT1"

	iter := local.PlaceSearchByCategory(groupcode).
		FormatAs("json").
		AuthorizeWith(common.REST_API_KEY).
		WithRadius(x, y, radius).
		Display(15).
		Result(1)

	prs := local.PlaceSearchResults{}

	for pr, err := iter.Next(); err == nil; pr, err = iter.Next() {
		prs = append(prs, pr)
	}

	if err := prs.SaveAs("category_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestCategorySearchWithXML(t *testing.T) {
	groupcode := "CS2"
	xmin := 127.05897078335246
	ymin := 37.506051888130386
	xmax := 128.05897078335276
	ymax := 38.506051888130406

	iter := local.PlaceSearchByCategory(groupcode).
		FormatAs("xml").
		AuthorizeWith(common.REST_API_KEY).
		WithRect(xmin, ymin, xmax, ymax).
		Display(15).
		Result(1)

	for pr, err := iter.Next(); err == nil; pr, err = iter.Next() {
		t.Log(pr)
	}
}

func TestCategorySearchWithSaveAsXML(t *testing.T) {
	var x float64 = 127.06283102249932
	var y float64 = 37.514322572335935
	radius := 2000
	groupcode := "MT1"

	iter := local.PlaceSearchByCategory(groupcode).
		FormatAs("xml").
		AuthorizeWith(common.REST_API_KEY).
		WithRadius(x, y, radius).
		Display(15).
		Result(1)

	prs := local.PlaceSearchResults{}

	for pr, err := iter.Next(); err == nil; pr, err = iter.Next() {
		prs = append(prs, pr)
	}

	if err := prs.SaveAs("category_search_test.xml"); err != nil {
		t.Error(err)
	}
}

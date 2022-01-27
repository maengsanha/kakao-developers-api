package daum_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestCafeSearchWithJSON(t *testing.T) {
	query := "손흥민"
	iter := daum.CafeSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("accuracy").
		Display(10).
		Result(1)

	for cr, err := iter.Next(); err == nil; cr, err = iter.Next() {
		t.Log(cr)
	}

}

func TestCafeSearchWithSaveAsJSON(t *testing.T) {
	query := "손흥민"
	iter := daum.CafeSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("recency").
		Display(5).
		Result(1)

	crs := daum.CafeSearchResults{}

	for cr, err := iter.Next(); err == nil; cr, err = iter.Next() {
		crs = append(crs, cr)
	}

	if err := crs.SaveAs("cafe_search_test.json"); err != nil {
		t.Error(err)
	}
}

package daum_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestCafeSearchSort(t *testing.T) {
	query := "손흥민"
	iter := daum.CafeSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("accuracy").
		Display(10).
		Result(10)

	cr, err := iter.Next()
	if err == nil {
		t.Log(cr)
	}

}

func TestCafeSearchSaveAsJSON(t *testing.T) {
	query := "손흥민"
	iter := daum.CafeSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("recency").
		Display(5).
		Result(1)

	crs := daum.CafeSearchResults{}
	cr, err := iter.Next()
	if err == nil {
		crs = append(crs, cr)
	}

	if err := crs.SaveAs("cafe_search_test.json"); err != nil {
		t.Error(err)
	}
}

package daum_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestImageSearchWithJSON(t *testing.T) {
	query := "g2"

	iter := daum.ImageSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("accuracy").
		Display(1).
		Result(1)

	for ir, err := iter.Next(); err == nil; ir, err = iter.Next() {
		t.Log(ir)
	}
}

func TestImageSearchWithSaveAsJSON(t *testing.T) {
	query := "g2"

	iter := daum.ImageSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("recency")

	irs := daum.ImageSearchResults{}

	for ir, err := iter.Next(); err == nil; ir, err = iter.Next() {
		irs = append(irs, ir)
	}

	if err := irs.SaveAs("image_search_test.json"); err != nil {
		t.Error(err)
	}
}

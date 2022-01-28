package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestDocumentSearchWithJSON(t *testing.T) {
	query := "Alan Turing"

	iter := daum.DocumentSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Result(10).
		Display(50)

	for dr, err := iter.Next(); err == nil; dr, err = iter.Next() {
		t.Log(dr)
	}
}

func TestDocumentSearchWithSaveAsJSON(t *testing.T) {
	query := "Alan Turing"

	iter := daum.DocumentSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Result(1).
		Display(30)

	drs := daum.DocumentSearchResults{}

	for dr, err := iter.Next(); err == nil; dr, err = iter.Next() {
		drs = append(drs, dr)
	}

	if err := drs.SaveAs("document_search_test.json"); err != nil {
		t.Error(err)
	}
}

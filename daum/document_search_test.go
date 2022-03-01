package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestDocumentSearchWithJSON(t *testing.T) {
	query := "Alan Turing"

	it := daum.DocumentSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Result(10).
		Display(50)

	for {
		item, err := it.Next()
		if err == daum.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		t.Log(item)
	}
}

func TestDocumentSearchWithSaveAsJSON(t *testing.T) {
	query := "Alan Turing"

	it := daum.DocumentSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Result(1).
		Display(30)

	items := daum.DocumentSearchResults{}

	for {
		item, err := it.Next()
		if err == daum.Done {
			break
		}
		if err != nil {
			t.Error(err)
		}
		items = append(items, item)
	}

	if err := items.SaveAs("document_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestDocumentSearchCollectAll(t *testing.T) {
	query := "Alan Turing"

	items := daum.DocumentSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Result(1).
		Display(50).
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

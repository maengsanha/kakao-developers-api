package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestImageSearchWithJSON(t *testing.T) {
	query := "g2"

	it := daum.ImageSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(1).
		Result(1)

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

func TestImageSearchWithSaveAsJSON(t *testing.T) {
	query := "g2"

	it := daum.ImageSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency")

	items := daum.ImageSearchResults{}

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

	if err := items.SaveAs("image_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestImageSearchCollectAll(t *testing.T) {
	query := "g2"

	items := daum.ImageSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(30).
		Result(1).
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

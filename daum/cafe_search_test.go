package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestCafeSearchWithJSON(t *testing.T) {
	query := "손흥민"
	it := daum.CafeSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(10).
		Result(1)

	for {
		item, err := it.Next()
		if err == daum.ErrEndPage {
			break
		}
		if err != nil {
			t.Error(err)
		}
		t.Log(item)
	}

}

func TestCafeSearchWithSaveAsJSON(t *testing.T) {
	query := "손흥민"
	it := daum.CafeSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Display(5).
		Result(1)

	items := daum.CafeSearchResults{}

	for {
		item, err := it.Next()
		if err == daum.ErrEndPage {
			break
		}
		if err != nil {
			t.Error(err)
		}
		items = append(items, item)
	}
	if err := items.SaveAs("cafe_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestCafeSearchCollectAll(t *testing.T) {
	query := "손흥민"
	items := daum.CafeSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(10).
		Result(1).
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

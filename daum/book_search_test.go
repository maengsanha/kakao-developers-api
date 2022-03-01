package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestBookSearchWithJSON(t *testing.T) {
	query := "히가시노 게이고"

	it := daum.BookSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("latest").
		Result(1).
		Display(10).
		Filter("person")

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

func TestBookSearchWithSaveAsJSON(t *testing.T) {
	query := "히가시노 게이고"

	it := daum.BookSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("latest").
		Result(1).
		Display(10).
		Filter("person")

	items := daum.BookSearchResults{}

	for {
		item, err := it.Next()
		if err == daum.Done {
			break
		}

		if err != nil {
			t.Error(err)
			break
		}
		items = append(items, item)
	}
	if err := items.SaveAs("book_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestBookSearchCollectAll(t *testing.T) {
	query := "히가시노 게이고"

	items := daum.BookSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("latest").
		Result(1).
		Display(10).
		Filter("person").
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

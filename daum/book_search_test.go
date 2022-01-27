package daum_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestBookSearchWithJSON(t *testing.T) {
	query := "밤은 짧아 걸어 아가씨야"

	iter := daum.BookSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("accuracy").
		Result(1).
		Display(10).
		Filter("title")

	for br, err := iter.Next(); err == nil; br, err = iter.Next() {
		t.Log(br)
	}
}

func TestBookSearchWithSaveAsJSON(t *testing.T) {
	query := "히가시노 게이고"

	iter := daum.BookSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("latest").
		Result(1).
		Display(10).
		Filter("person")

	brs := daum.BookSearchResults{}

	for br, err := iter.Next(); err == nil; br, err = iter.Next() {
		brs = append(brs, br)
	}

	if err := brs.SaveAs("book_search_test.json"); err != nil {
		t.Error(err)
	}
}

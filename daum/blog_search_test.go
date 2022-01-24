package daum_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestBlogSearchWithJSON(t *testing.T) {
	query := "Imitation Game"

	iter := daum.BlogSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("accuracy").
		Result(10).
		Display(50)

	for br, err := iter.Next(); err == nil; br, err = iter.Next() {
		t.Log(br)
	}
}

func TestBlogSearchWithSaveAsJSON(t *testing.T) {
	query := "Imitation Game"

	iter := daum.BlogSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("recency").
		Result(1).
		Display(30)

	brs := daum.BlogSearchResults{}

	for br, err := iter.Next(); err == nil; br, err = iter.Next() {
		brs = append(brs, br)
	}

	if err := brs.SaveAs("blog_search_test.json"); err != nil {
		t.Error(err)
	}
}

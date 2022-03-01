package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestBlogSearchWithJSON(t *testing.T) {
	query := "Imitation Game"

	it := daum.BlogSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(50).
		Result(10)

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

func TestBlogSearchWithSaveAsJSON(t *testing.T) {
	query := "Imitation Game"

	it := daum.BlogSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Display(30).
		Result(1)

	items := daum.BlogSearchResults{}

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

	if err := items.SaveAs("blog_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestBlogSearchCollectAll(t *testing.T) {
	query := "Imitation Game"

	items := daum.BlogSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Display(50).
		Result(1).
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

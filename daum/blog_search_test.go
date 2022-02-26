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

func TestBlogSearchWithSaveAsJSON(t *testing.T) {
	query := "Imitation Game"

	it := daum.BlogSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("recency").
		Result(1).
		Display(30)

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

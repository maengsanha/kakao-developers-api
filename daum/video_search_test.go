package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestVideoSearchWithJSON(t *testing.T) {
	query := "major scale"

	it := daum.VideoSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(30).
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

func TestVideoSearchWithSaveAsJSON(t *testing.T) {
	query := "minor scale"

	it := daum.VideoSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(30).
		Result(1)

	items := daum.VideoSearchResults{}

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

	if err := items.SaveAs("video_search_test.json"); err != nil {
		t.Error(err)
	}
}

func TestVideoSearchCollectAll(t *testing.T) {
	query := "minor scale"

	items := daum.VideoSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		CollectAll()

	for _, item := range items {
		t.Log(item)
	}
}

package daum_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestVideoSearchWithJSON(t *testing.T) {
	query := "major scale"

	iter := daum.VideoSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(30).
		Result(1)

	for vr, err := iter.Next(); err == nil; vr, err = iter.Next() {
		t.Log(vr)
	}
}

func TestVideoSearchWithSaveAsJSON(t *testing.T) {
	query := "minor scale"

	iter := daum.VideoSearch(query).
		AuthorizeWith(common.REST_API_KEY).
		SortBy("accuracy").
		Display(30).
		Result(1)

	vrs := daum.VideoSearchResults{}

	for vr, err := iter.Next(); err == nil; vr, err = iter.Next() {
		vrs = append(vrs, vr)
	}

	if err := vrs.SaveAs("video_search_test.json"); err != nil {
		t.Error(err)
	}
}

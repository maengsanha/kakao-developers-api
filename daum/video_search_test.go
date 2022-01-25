package daum_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-client/daum"
)

func TestVideoSearchWithJSON(t *testing.T) {
	query := "major scale code"

	iter := daum.VideoSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
		SortBy("accuracy").
		Display(30).
		Result(1)

	for vr, err := iter.Next(); err == nil; vr, err = iter.Next() {
		t.Log(vr)
	}
}

func TestVideoSearchWithSaveAsJSON(t *testing.T) {
	query := "minor code"

	iter := daum.VideoSearch(query).
		AuthorizeWith(daum.REST_API_KEY).
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

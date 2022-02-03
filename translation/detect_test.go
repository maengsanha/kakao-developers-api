package translation_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/translation"
)

func TestDetectWithJSON(t *testing.T) {
	query := "안녕하세요"

	if dr, err := translation.Detect(query).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(dr)
	}
}

func TestDetectWithSaveAsJSON(t *testing.T) {
	query := "안녕하세요"

	if dr, err := translation.Detect(query).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = dr.SaveAs("detect_test.json"); err != nil {
		t.Error(err)
	}
}

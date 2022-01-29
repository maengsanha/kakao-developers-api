package translation_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/translation"
)

func TestDetectLanguageWithJSONByGET(t *testing.T) {
	query := "안녕하세요"
	method := "GET"

	if dr, err := translation.DetectLanguage(query).
		AuthorizeWith(common.REST_API_KEY).
		RequestBy(method); err != nil {
		t.Error(err)

	} else {
		t.Log(dr)
	}
}

func TestDetectLanguageWithSaveAsJSONByGET(t *testing.T) {
	query := "안녕하세요"
	method := "GET"
	if dr, err := translation.DetectLanguage(query).
		AuthorizeWith(common.REST_API_KEY).
		RequestBy(method); err != nil {
		t.Error(err)
	} else if err = dr.SaveAs("detect_test_by_get.json"); err != nil {
		t.Error(err)
	}
}

func TestDetectLanguageWithJSONByPOST(t *testing.T) {
	query := "안녕하세요"
	method := "POST"
	if dr, err := translation.DetectLanguage(query).
		AuthorizeWith(common.REST_API_KEY).
		RequestBy(method); err != nil {
		t.Error(err)
	} else {
		t.Log(dr)
	}
}

func TestDetectLanguageWithSaveAsJSONByPOST(t *testing.T) {
	query := "안녕하세요"
	method := "POST"
	if dr, err := translation.DetectLanguage(query).
		AuthorizeWith(common.REST_API_KEY).
		RequestBy(method); err != nil {
		t.Error(err)
	} else if dr.SaveAs("detect_test_by_post.json"); err != nil {
		t.Error(err)
	}
}

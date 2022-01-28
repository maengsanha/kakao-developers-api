package translation_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/translation"
)

func TestDetectLanguageByGET(t *testing.T) {
	query := "안녕하세요"
	method := "GET"
	if dl, err := translation.DetectLanguage(query).
		AuthorizeWith(common.REST_API_KEY).
		RequestBy(method); err != nil {
		t.Error(err)
	} else {
		t.Log(dl)
	}
}

func TestDetectLanguageByGETSaveAsJSON(t *testing.T) {
	query := "안녕하세요"
	method := "GET"
	if dl, err := translation.DetectLanguage(query).
		AuthorizeWith(common.REST_API_KEY).
		RequestBy(method); err != nil {
		t.Error(err)
	} else if err = dl.SaveAs("detect_test_by_get.json"); err != nil {
		t.Error(err)
	}
}

func TestDetectLanguageByPOST(t *testing.T) {
	query := "안녕하세요"
	method := "POST"
	if dl, err := translation.DetectLanguage(query).
		AuthorizeWith(common.REST_API_KEY).
		RequestBy(method); err != nil {
		t.Error(err)
	} else {
		t.Log(dl)
	}
}

func TestDetectLanguageByPOSTSaveAsJSON(t *testing.T) {
	query := "안녕하세요"
	method := "POST"
	if dl, err := translation.DetectLanguage(query).
		AuthorizeWith(common.REST_API_KEY).
		RequestBy(method); err != nil {
		t.Error(err)
	} else if dl.SaveAs("detect_test_by_post.json"); err != nil {
		t.Error(err)
	}
}

package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestOCR(t *testing.T) {
	filepath := "/home/js/test6.jpeg"

	if or, err := vision.OCR(filepath).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(or)
	}
}
func TestOCRSaveAsJSON(t *testing.T) {
	filepath := "/home/js/test6.jpeg"

	if or, err := vision.OCR(filepath).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = or.SaveAs("ocr_test.json"); err != nil {
		t.Error(err)
	}
}

package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestProductDetectWithUrl(t *testing.T) {
	source := "https://topguide.kr/wp-content/uploads/2020/03/image-689-1024x828.jpg"
	if pr, err := vision.ProductDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.7).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(pr)
	}
}

func TestProductDetectWithUrlSaveAsJSON(t *testing.T) {
	source := "https://topguide.kr/wp-content/uploads/2020/03/image-689-1024x828.jpg"
	if pr, err := vision.ProductDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.7).
		Collect(); err != nil {
		t.Error(err)
	} else if err = pr.SaveAs("product_detect_url_test.json"); err != nil {
		t.Error(pr)
	}
}

func TestProductDetectWithFile(t *testing.T) {
	source := "/home/js/test2.jpg"
	if pr, err := vision.ProductDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.7).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(pr)
	}
}

func TestProductDetectWithFileSaveAsJSON(t *testing.T) {
	source := "/home/js/test2.jpg"
	if pr, err := vision.ProductDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.7).
		Collect(); err != nil {
		t.Error(err)
	} else if err = pr.SaveAs("product_detect_file_test.json"); err != nil {
		t.Error(pr)
	}
}

package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestThumbnailDetectWithUrl(t *testing.T) {
	source := "https://img.khan.co.kr/news/2021/09/30/l_2021093001003585000310901.jpg"

	if tr, err := vision.ThumbnailDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestThumbnailDetectWithUrlSaveAsJSON(t *testing.T) {
	source := "https://img.khan.co.kr/news/2021/09/30/l_2021093001003585000310901.jpg"

	if tr, err := vision.ThumbnailDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else if err = tr.SaveAs("thumbnail_detect_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestThumbnailDetectWithFile(t *testing.T) {
	source := "/home/js/test4.jpg"

	if tr, err := vision.ThumbnailDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestThumbnailDetectWithFileSaveAsJSON(t *testing.T) {
	source := "/home/js/test4.jpg"

	if tr, err := vision.ThumbnailDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else if err = tr.SaveAs("thumbnail_detect_file_test.json"); err != nil {
		t.Error(err)
	}
}

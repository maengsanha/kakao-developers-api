package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestThumbnailCreateWithUrl(t *testing.T) {
	source := "https://img.khan.co.kr/news/2021/09/30/l_2021093001003585000310901.jpg"

	if tr, err := vision.ThumbnailCreate(source).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestThumbnailCreateWithUrlSaveAsJSON(t *testing.T) {
	source := "https://img.khan.co.kr/news/2021/09/30/l_2021093001003585000310901.jpg"

	if tr, err := vision.ThumbnailCreate(source).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else if tr.SaveAs("thumbnail_create_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestThumbnailCreateWithFile(t *testing.T) {
	source := "/home/js/test4.jpg"

	if tr, err := vision.ThumbnailCreate(source).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestThumbnailCreateWithFileSaveAsJSON(t *testing.T) {
	source := "/home/js/test4.jpg"

	if tr, err := vision.ThumbnailCreate(source).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else if tr.SaveAs("thumbnail_create_file_test.json"); err != nil {
		t.Error(err)
	}
}

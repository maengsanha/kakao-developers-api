package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestThumbnailCreateWithURL(t *testing.T) {
	url := "https://img.khan.co.kr/news/2021/09/30/l_2021093001003585000310901.jpg"

	if tr, err := vision.ThumbnailCreate().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestThumbnailCreateWithURLSaveAsJSON(t *testing.T) {
	url := "https://img.khan.co.kr/news/2021/09/30/l_2021093001003585000310901.jpg"

	if tr, err := vision.ThumbnailCreate().
		WithURL(url).
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
	filepath := "/home/js/test4.jpg"

	if tr, err := vision.ThumbnailCreate().
		WithFile(filepath).
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
	filepath := "/home/js/test4.jpg"

	if tr, err := vision.ThumbnailCreate().
		WithFile(filepath).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(200).
		HeightTo(200).
		Collect(); err != nil {
		t.Error(err)
	} else if tr.SaveAs("thumbnail_create_file_test.json"); err != nil {
		t.Error(err)
	}
}

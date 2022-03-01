package pose_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/pose"
)

func TestAnalyzeImageWithURL(t *testing.T) {
	imageurl := "https://pbs.twimg.com/media/EiqWMtcWkAEDgZh.jpg"

	if ir, err := pose.AnalyzeImage().
		WithURL(imageurl).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ir)
	}
}

func TestAnalyzeImageWithURLSaveAsJSON(t *testing.T) {
	imageurl := "https://pbs.twimg.com/media/EiqWMtcWkAEDgZh.jpg"

	if ir, err := pose.AnalyzeImage().
		WithURL(imageurl).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = ir.SaveAs("analyze_image_test_url.json"); err != nil {
		t.Log(ir)
	}
}

func TestAnalyzeImageWithFile(t *testing.T) {
	imagepath := "test.jpeg"

	if ir, err := pose.AnalyzeImage().
		WithFile(imagepath).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ir)
	}
}

func TestAnalyzeImageWithFileSaveAsJSON(t *testing.T) {
	imagepath := "testimage.jpg"

	if ir, err := pose.AnalyzeImage().
		WithFile(imagepath).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = ir.SaveAs("analyze_image_test_file.json"); err != nil {
		t.Log(ir)
	}
}

package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestAdultImageDetectWithURL(t *testing.T) {
	url := "https://dimg.donga.com/wps/NEWS/IMAGE/2021/11/12/110211591.2.jpg"

	if ar, err := vision.AdultImageDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ar)
	}
}

func TestAdultImageDetectWithURLSaveAsJson(t *testing.T) {
	url := "https://dimg.donga.com/wps/NEWS/IMAGE/2021/11/12/110211591.2.jpg"

	if ar, err := vision.AdultImageDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = ar.SaveAs("adult_image_detect_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestAdultImageDetectWithFile(t *testing.T) {
	filepath := "/home/js/test3.jpg"

	if ar, err := vision.AdultImageDetect().
		WithFile(filepath).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ar)
	}
}

func TestAdultImageDetectWithFileSaveAsJson(t *testing.T) {
	filepath := "/home/js/test3.jpg"

	if ar, err := vision.AdultImageDetect().
		WithFile(filepath).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = ar.SaveAs("adult_image_detect_file_test.json"); err != nil {
		t.Error(err)
	}
}

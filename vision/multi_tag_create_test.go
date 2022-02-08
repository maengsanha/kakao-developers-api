package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestMultiTagCreateWithUrl(t *testing.T) {
	source := "https://cdn-asia.heykorean.com/community/uploads/images/2019/06/1561461763.png"

	if mr, err := vision.MultiTagCreate(source).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(mr)
	}
}

func TestMultiTagCreateWithUrlSaveAsJson(t *testing.T) {
	source := "https://cdn-asia.heykorean.com/community/uploads/images/2019/06/1561461763.png"

	if mr, err := vision.MultiTagCreate(source).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = mr.SaveAs("multi_tag_create_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestMultiTagCreateWithFile(t *testing.T) {
	source := "/home/js/test2.jpg"

	if mr, err := vision.MultiTagCreate(source).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(mr)
	}
}

func TestMultiTagCreateWithFileSaveAsJson(t *testing.T) {
	source := "/home/js/test2.jpg"

	if mr, err := vision.MultiTagCreate(source).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = mr.SaveAs("multi_tag_create_file_test.json"); err != nil {
		t.Error(err)
	}
}

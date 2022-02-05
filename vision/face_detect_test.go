package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestFaceDetectWithUrl(t *testing.T) {
	source := "https://resources.premierleague.com/premierleague/photos/players/250x250/p85971.png"

	if fr, err := vision.FaceDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(fr)
	}
}

func TestFaceDetectWithUrlSaveAsJSON(t *testing.T) {
	source := "https://resources.premierleague.com/premierleague/photos/players/250x250/p85971.png"

	if fr, err := vision.FaceDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = fr.SaveAs("face_detection_Url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestFaceDetectWithFile(t *testing.T) {
	source := "/home/js/test.jpg"

	if fr, err := vision.FaceDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(fr)
	}
}

func TestFaceDetectWithFileSaveAsJSON(t *testing.T) {
	source := "/home/js/test.jpg"

	if fr, err := vision.FaceDetect(source).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = fr.SaveAs("face_detection_File_test.json"); err != nil {
		t.Error(fr)
	}
}

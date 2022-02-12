package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestFaceDetectWithURL(t *testing.T) {
	url := "https://resources.premierleague.com/premierleague/photos/players/250x250/p85971.png"

	if fr, err := vision.FaceDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(fr)
	}
}

func TestFaceDetectWithURLSaveAsJSON(t *testing.T) {
	url := "https://resources.premierleague.com/premierleague/photos/players/250x250/p85971.png"

	if fr, err := vision.FaceDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = fr.SaveAs("face_detect_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestFaceDetectWithFile(t *testing.T) {
	filepath := "/home/js/test.jpg"

	if fr, err := vision.FaceDetect().
		WithFile(filepath).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.9).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(fr)
	}
}

func TestFaceDetectWithFileSaveAsJSON(t *testing.T) {
	filepath := "/home/js/test.jpg"

	if fr, err := vision.FaceDetect().
		WithFile(filepath).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.9).
		Collect(); err != nil {
		t.Error(err)
	} else if err = fr.SaveAs("face_detect_file_test.json"); err != nil {
		t.Error(fr)
	}
}

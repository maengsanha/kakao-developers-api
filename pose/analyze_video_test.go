package pose_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/pose"
)

func TestVideoAnalyzeWithURL(t *testing.T) {
	video_url := "https://raw.githubusercontent.com/intel-iot-devkit/sample-videos/master/face-demographics-walking.mp4"

	if vr, err := pose.AnalyzeVideo().
		WithURL(video_url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(vr)
	}
}

func TestVideoAnalyzeWithFile(t *testing.T) {
	file_path := "/Users/goryne/Downloads/testvideo.mp4"

	if vr, err := pose.AnalyzeVideo().
		WithFile(file_path).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(vr)
	}
}

package pose_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/pose"
)

func TestVideoAnalyzeResult(t *testing.T) {
	id := "9524567f-887b-474f-9e33-a3d480b400c1"

	if cr, err := pose.CheckVideo(id).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}
}

func TestVideoAnalyzeResultSaveAsJSON(t *testing.T) {
	id := "9524567f-887b-474f-9e33-a3d480b400c1"

	if cr, err := pose.CheckVideo(id).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err := cr.SaveAs("video_analysis_result_test.json"); err != nil {
		t.Error(cr)
	}
}

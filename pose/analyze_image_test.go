package pose

import (
	"internal/common"
	"testing"
)

func TestAnalyzeImageWithJSON(t *testing.T) {
	imageurl := "https://pbs.twimg.com/media/EiqWMtcWkAEDgZh.jpg"

	if ir, err := ImageAnalyze(imageurl).
		AuthorizeWith(common.KeyPrefix).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ir)
	}
}

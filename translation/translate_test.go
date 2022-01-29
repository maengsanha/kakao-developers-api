package translation_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/translation"
)

func TestTranslateWithJSON(t *testing.T) {
	query := "화성학 5도 스케일"

	if tr, err := translation.Translation(query).
		Source("kr").
		Target("en").
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

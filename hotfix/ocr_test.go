package hotfix_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/hotfix"
)

func TestOCR(t *testing.T) {
	image_path := "test-image.jpeg"

	result := hotfix.OCR(image_path, common.REST_API_KEY)

	t.Logf("[OCR] output:\n%v\n", result)
}

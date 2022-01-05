package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestAddressSearch(t *testing.T) {
	query := "전북 삼성동 100"
	key := "0xdeadbeef"

	iter := local.AddressSearch(query).
		AuthorizeWith(key).
		Result(local.DefaultPage).
		Display(local.DefaultSize).
		Analyze(local.Similar).
		As(local.JSON)

	for page, err := iter.Next(); ; {
		if err != nil {
			if err != local.ErrEndPage {
				t.Errorf("error: %v\n", err)
			}
			break
		}
		t.Logf("result: %v\n", page)
	}
}

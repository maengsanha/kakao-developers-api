package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestAddressSearch(t *testing.T) {
	query := "성북구 정릉동"
	key := ""

	iter := local.AddressSearch(query).
		AuthorizeWith(key).
		Result(local.DefaultPage).
		Display(local.DefaultSize).
		Analyze(local.Similar).
		As(local.JSON)

	for res, err := iter.Next(); ; {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

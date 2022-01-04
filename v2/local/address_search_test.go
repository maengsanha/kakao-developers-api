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
		Display(local.MaxSize).
		Result(local.MinPage).
		Analyze(local.Similar).
		As(local.JSON)

	for resp, err := iter.Next(); err != nil; {
		t.Logf("response: %v", resp)
	}
}

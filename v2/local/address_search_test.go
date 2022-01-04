package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestAddressSearch(t *testing.T) {
	iter := local.AddressSearch("전북 삼성동 100").
		Analyze("similar").
		Display(10).
		Result(1)

	for resp, err := iter.Next(); iter != nil; {
		t.Log(resp, err)
	}
}

package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

// Category code list
// MT1	대형마트
// CS2	편의점
// PS3	어린이집, 유치원
// SC4	학교
// AC5	학원
// PK6	주차장
// OL7	주유소, 충전소
// SW8	지하철역
// BK9	은행
// CT1	문화시설
// AG2	중개업소
// PO3	공공기관
// AT4	관광명소
// AD5	숙박
// FD6	음식점
// CE7	카페
// HP8	병원
// PM9	약국

func TestCategorySearchWithJSON(t *testing.T) {
	key := ""
	x := "127.06283102249932"
	y := "37.514322572335935"
	radius := 2000
	categorygroupcode := "MT1"

	if res, err := local.CategorySearch(categorygroupcode).
		FormatJSON().
		AuthorizeWith(key).
		SetLongitude(x).
		SetLatitude(y).
		SetRadius(radius).
		Display(15).
		Result(1).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestCategorySearchWithXML(t *testing.T) {
	key := ""
	rect := "1"
	categorygroupcode := "MT1"

	if res, err := local.CategorySearch(categorygroupcode).
		FormatXML().
		AuthorizeWith(key).
		SetRect(rect).
		Display(15).
		Result(1).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

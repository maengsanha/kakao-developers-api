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
	var x float64 = 127.06283102249932
	var y float64 = 37.514322572335935
	radius := 2000
	groupcode := "MT1"

	iter := local.PlaceSearchByCategory(groupcode).
		FormatJSON().
		AuthorizeWith(key).
		WithRadius(x, y, radius).
		Display(15).
		Result(1)

	for res, err := iter.Next(); ; res, err = iter.Next() {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

func TestCategorySearchWithXML(t *testing.T) {
	key := ""
	groupcode := "CS2"
	xmin := 127.05897078335246
	ymin := 37.506051888130386
	xmax := 128.05897078335276
	ymax := 38.506051888130406

	iter := local.PlaceSearchByCategory(groupcode).
		FormatXML().
		AuthorizeWith(key).
		WithRect(xmin, ymin, xmax, ymax).
		Display(15).
		Result(1)

	for res, err := iter.Next(); ; res, err = iter.Next() {
		t.Log(res)
		if err != nil {
			if err != local.ErrEndPage {
				t.Error(err)
			}
			break
		}
	}
}

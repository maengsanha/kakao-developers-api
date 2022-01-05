package local_test

import (
	"testing"

	"github.com/maengsanha/kakao-developers-api/v2/local"
)

func TestKeyWordSearch(t *testing.T) {
	query := "카카오프렌즈"
	key := ""

	if res, err := local.KeyWordSearch(query).
		As(local.JSON).
		AuthorizeWith(key).
		SetCategoryGroupCode(local.DefaultCategoryGroupCode).
		SetX(local.DefaultX).
		SetY(local.DefaultY).
		SetRadius(local.DefaultRadius).
		SetRect(local.DefaultRect).
		Result(local.DefaultPage).
		Display(local.DefaultSize).
		SortType(local.DefaultSort).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}

}

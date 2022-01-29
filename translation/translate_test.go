package translation_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/translation"
)

func TestTranslateWithJSONByGET(t *testing.T) {
	query := "이"

	if tr, err := translation.Translate(query).
		Source("kr").
		Target("en").
		AuthorizeWith(common.REST_API_KEY).
		CollectByGET(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestTranslateWithSaveAsJSONByGET(t *testing.T) {
	query := "이 대성당이라는 작품은 아주 짧은 시간 내에서의 한정된 공간의 사건을 다루고 있지만 작품의 의미에 대한 무게는 장편 소설 못지않게 강렬하다. 또한 단편 소설만의 간략한 서술의 특징으로 독자의 행동반경을 더욱 더 자유롭게 하여주었다.이 작품은 기본적으로 성장 소설의 흐름과 유사점을 보여준다.다만 그 대상이 이미 주체화된 어른이라는 점을 주목해 볼 필요가 있다.성장이란 단어가 아직 완성되지 않은 아이들에게 한정되는 단어로서 오인할 수 있지만 기존의 삶에 지치고 고착된 어른들의 삶에도 아이 못지않게 성장이란 단어가 절실하게 다가올 수 있다. 마찬가지로 작품에서 화자의 아내가 시를 쓰는 것을 자신의 유일한 탈출구로 삼은 것은 어떤 의미에서는 지겨운 현실에서의 삶의 안주와 극복되지 못하는 현실에 염증을 느끼고 새로운 ‘성장’을 희망하는 욕망의 표출이다. 그리고 아내의 시를 새로운 성장을 희망하는 시가 있고 또한 그렇지 못한 시로 분류할 수 있다. 전자의 경우는 맹인 친구가 아내 얼굴의 모든 부분부터 목까지 그의 손가락으로 만졌을 때 생겼던 느낌을 표현한 시로 대표된다. 또 후자는 아내가 공군 행정관의 아내로서 가졌던 느낌을 표현한 미완성의 시로 대표된다. 이 시는 긍정적인 성장의 모습을 도저히 끄집어 낼 수 없었기 때문에 화자의 아내는 아직 완성할 수 없었다."

	if tr, err := translation.Translate(query).
		Source("kr").
		Target("en").
		AuthorizeWith(common.REST_API_KEY).
		CollectByGET(); err != nil {
		t.Error(err)
	} else if err := tr.SaveAs("translate_test_get.json"); err != nil {
		t.Log(tr)
	}
}

func TestTranslateWithJSONByPOST(t *testing.T) {
	query := "이 대성당이라는 작품은 아주 짧은 시간 내에서의 한정된 공간의 사건을 다루고 있지만 작품의 의미에 대한 무게는 장편 소설 못지않게 강렬하다."

	if tr, err := translation.Translate(query).
		Source("kr").
		Target("en").
		AuthorizeWith(common.REST_API_KEY).
		CollectByPOST(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestTranslateWithSaveAsJSONByPOST(t *testing.T) {
	query := "이 대성당이라는 작품은 아주 짧은 시간 내에서의 한정된 공간의 사건을 다루고 있지만 작품의 의미에 대한 무게는 장편 소설 못지않게 강렬하다. 또한 단편 소설만의 간략한 서술의 특징으로 독자의 행동반경을 더욱 더 자유롭게 하여주었다.이 작품은 기본적으로 성장 소설의 흐름과 유사점을 보여준다.다만 그 대상이 이미 주체화된 어른이라는 점을 주목해 볼 필요가 있다.성장이란 단어가 아직 완성되지 않은 아이들에게 한정되는 단어로서 오인할 수 있지만 기존의 삶에 지치고 고착된 어른들의 삶에도 아이 못지않게 성장이란 단어가 절실하게 다가올 수 있다. 마찬가지로 작품에서 화자의 아내가 시를 쓰는 것을 자신의 유일한 탈출구로 삼은 것은 어떤 의미에서는 지겨운 현실에서의 삶의 안주와 극복되지 못하는 현실에 염증을 느끼고 새로운 ‘성장’을 희망하는 욕망의 표출이다. 그리고 아내의 시를 새로운 성장을 희망하는 시가 있고 또한 그렇지 못한 시로 분류할 수 있다. 전자의 경우는 맹인 친구가 아내 얼굴의 모든 부분부터 목까지 그의 손가락으로 만졌을 때 생겼던 느낌을 표현한 시로 대표된다. 또 후자는 아내가 공군 행정관의 아내로서 가졌던 느낌을 표현한 미완성의 시로 대표된다. 이 시는 긍정적인 성장의 모습을 도저히 끄집어 낼 수 없었기 때문에 화자의 아내는 아직 완성할 수 없었다."

	if tr, err := translation.Translate(query).
		Source("kr").
		Target("en").
		AuthorizeWith(common.REST_API_KEY).
		CollectByPOST(); err != nil {
		t.Error(err)
	} else if err := tr.SaveAs("translate_test_post.json"); err != nil {
		t.Log(tr)
	}
}

// Copyright 2022 Sanha Maeng, Soyang Baek, Jinmyeong Kim
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vision_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/vision"
)

func TestThumbnailDetectWithURL(t *testing.T) {
	url := "https://img.khan.co.kr/news/2021/09/30/l_2021093001003585000310901.jpg"

	if tr, err := vision.ThumbnailDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(400).
		HeightTo(400).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestThumbnailDetectWithURLSaveAsJSON(t *testing.T) {
	url := "https://img.khan.co.kr/news/2021/09/30/l_2021093001003585000310901.jpg"

	if tr, err := vision.ThumbnailDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(400).
		HeightTo(400).
		Collect(); err != nil {
		t.Error(err)
	} else if err = tr.SaveAs("thumbnail_detect_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestThumbnailDetectWithFile(t *testing.T) {
	filename := "test4.jpg"

	if tr, err := vision.ThumbnailDetect().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(500).
		HeightTo(500).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestThumbnailDetectWithFileSaveAsJSON(t *testing.T) {
	filename := "test4.jpg"

	if tr, err := vision.ThumbnailDetect().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		WidthTo(500).
		HeightTo(500).
		Collect(); err != nil {
		t.Error(err)
	} else if err = tr.SaveAs("thumbnail_detect_file_test.json"); err != nil {
		t.Error(err)
	}
}

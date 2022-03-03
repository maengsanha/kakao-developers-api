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

func TestProductDetectWithURL(t *testing.T) {
	url := "https://topguide.kr/wp-content/uploads/2020/03/image-689-1024x828.jpg"
	if pr, err := vision.ProductDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.7).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(pr)
	}
}

func TestProductDetectWithURLSaveAsJSON(t *testing.T) {
	url := "https://topguide.kr/wp-content/uploads/2020/03/image-689-1024x828.jpg"
	if pr, err := vision.ProductDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.7).
		Collect(); err != nil {
		t.Error(err)
	} else if err = pr.SaveAs("product_detect_url_test.json"); err != nil {
		t.Error(pr)
	}
}

func TestProductDetectWithFile(t *testing.T) {
	filename := "test2.jpg"
	if pr, err := vision.ProductDetect().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.7).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(pr)
	}
}

func TestProductDetectWithFileSaveAsJSON(t *testing.T) {
	filename := "test2.jpg"
	if pr, err := vision.ProductDetect().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.7).
		Collect(); err != nil {
		t.Error(err)
	} else if err = pr.SaveAs("product_detect_file_test.json"); err != nil {
		t.Error(pr)
	}
}

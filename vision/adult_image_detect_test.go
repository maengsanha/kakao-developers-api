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

func TestAdultImageDetectWithURL(t *testing.T) {
	url := "https://dimg.donga.com/wps/NEWS/IMAGE/2021/11/12/110211591.2.jpg"

	if ar, err := vision.AdultImageDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ar)
	}
}

func TestAdultImageDetectWithURLSaveAsJson(t *testing.T) {
	url := "https://dimg.donga.com/wps/NEWS/IMAGE/2021/11/12/110211591.2.jpg"

	if ar, err := vision.AdultImageDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = ar.SaveAs("adult_image_detect_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestAdultImageDetectWithFile(t *testing.T) {
	filename := "test3.jpg"

	if ar, err := vision.AdultImageDetect().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ar)
	}
}

func TestAdultImageDetectWithFileSaveAsJson(t *testing.T) {
	filename := "test3.jpg"

	if ar, err := vision.AdultImageDetect().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = ar.SaveAs("adult_image_detect_file_test.json"); err != nil {
		t.Error(err)
	}
}

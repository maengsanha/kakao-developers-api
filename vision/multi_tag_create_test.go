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

func TestMultiTagCreateWithURL(t *testing.T) {
	url := "https://cdn-asia.heykorean.com/community/uploads/images/2019/06/1561461763.png"

	if mr, err := vision.MultiTagCreate().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(mr)
	}
}

func TestMultiTagCreateWithURLSaveAsJson(t *testing.T) {
	url := "https://cdn-asia.heykorean.com/community/uploads/images/2019/06/1561461763.png"

	if mr, err := vision.MultiTagCreate().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = mr.SaveAs("multi_tag_create_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestMultiTagCreateWithFile(t *testing.T) {
	filename := "test2.jpg"

	if mr, err := vision.MultiTagCreate().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(mr)
	}
}

func TestMultiTagCreateWithFileSaveAsJson(t *testing.T) {
	filename := "test2.jpg"

	if mr, err := vision.MultiTagCreate().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = mr.SaveAs("multi_tag_create_file_test.json"); err != nil {
		t.Error(err)
	}
}

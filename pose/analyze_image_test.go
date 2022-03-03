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

package pose_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/pose"
)

func TestAnalyzeImageWithURL(t *testing.T) {
	imageurl := "https://pbs.twimg.com/media/EiqWMtcWkAEDgZh.jpg"

	if ir, err := pose.AnalyzeImage().
		WithURL(imageurl).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ir)
	}
}

func TestAnalyzeImageWithURLSaveAsJSON(t *testing.T) {
	imageurl := "https://pbs.twimg.com/media/EiqWMtcWkAEDgZh.jpg"

	if ir, err := pose.AnalyzeImage().
		WithURL(imageurl).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = ir.SaveAs("analyze_image_test_url.json"); err != nil {
		t.Log(ir)
	}
}

func TestAnalyzeImageWithFile(t *testing.T) {
	imagepath := "test.jpeg"

	if ir, err := pose.AnalyzeImage().
		WithFile(imagepath).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(ir)
	}
}

func TestAnalyzeImageWithFileSaveAsJSON(t *testing.T) {
	imagepath := "testimage.jpg"

	if ir, err := pose.AnalyzeImage().
		WithFile(imagepath).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = ir.SaveAs("analyze_image_test_file.json"); err != nil {
		t.Log(ir)
	}
}

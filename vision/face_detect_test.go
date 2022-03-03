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

func TestFaceDetectWithURL(t *testing.T) {
	url := "https://resources.premierleague.com/premierleague/photos/players/250x250/p85971.png"

	if fr, err := vision.FaceDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(fr)
	}
}

func TestFaceDetectWithURLSaveAsJSON(t *testing.T) {
	url := "https://resources.premierleague.com/premierleague/photos/players/250x250/p85971.png"

	if fr, err := vision.FaceDetect().
		WithURL(url).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err = fr.SaveAs("face_detect_url_test.json"); err != nil {
		t.Error(err)
	}
}

func TestFaceDetectWithFile(t *testing.T) {
	filename := "test.jpg"

	if fr, err := vision.FaceDetect().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.9).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(fr)
	}
}

func TestFaceDetectWithFileSaveAsJSON(t *testing.T) {
	filename := "test.jpg"

	if fr, err := vision.FaceDetect().
		WithFile(filename).
		AuthorizeWith(common.REST_API_KEY).
		ThresholdAt(0.9).
		Collect(); err != nil {
		t.Error(err)
	} else if err = fr.SaveAs("face_detect_file_test.json"); err != nil {
		t.Error(fr)
	}
}

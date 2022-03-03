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

func TestVideoAnalyzeResult(t *testing.T) {
	id := "9524567f-887b-474f-9e33-a3d480b400c1"

	if cr, err := pose.CheckVideo(id).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}
}

func TestVideoAnalyzeResultSaveAsJSON(t *testing.T) {
	id := "9524567f-887b-474f-9e33-a3d480b400c1"

	if cr, err := pose.CheckVideo(id).
		AuthorizeWith(common.REST_API_KEY).
		Collect(); err != nil {
		t.Error(err)
	} else if err := cr.SaveAs("video_analysis_result_test.json"); err != nil {
		t.Error(cr)
	}
}

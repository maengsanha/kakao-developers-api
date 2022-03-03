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

package local_test

import (
	"internal/common"
	"testing"

	"github.com/maengsanha/kakao-developers-client/local"
)

func TestTransCoordWithJSON(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if tr, err := local.TransCoord(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestTransCoordWithSaveAsJSON(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if tr, err := local.TransCoord(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else if err = tr.SaveAs("transcoord_test.json"); err != nil {
		t.Error(err)
	}
}

func TestTransCoordWithXML(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if tr, err := local.TransCoord(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(tr)
	}
}

func TestTransCoordWithSaveAsXML(t *testing.T) {
	x := 160710.37729270622
	y := -4388.879299157299

	if tr, err := local.TransCoord(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WTM").
		Output("WCONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else if err = tr.SaveAs("transcoord_test.xml"); err != nil {
		t.Error(err)
	}
}

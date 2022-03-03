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

func TestCoordToDistrictWithJSON(t *testing.T) {
	x := 127.1086228
	y := 37.4012191

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WGS84").
		Output("WGS84").
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}
}

func TestCoordToDistrictWithSaveAsJSON(t *testing.T) {
	x := 127.1086228
	y := 37.4012191

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WGS84").
		Output("WGS84").
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2district_test.json"); err != nil {
		t.Error(err)
	}
}

func TestCoordToDistrictWithXML(t *testing.T) {
	x := 127.1086228
	y := 37.4012191

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WGS84").
		Output("CONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}
}

func TestCoordToDistrictWithSaveAsXML(t *testing.T) {
	x := 127.1086228
	y := 37.4012191

	if cr, err := local.CoordToDistrict(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input("WGS84").
		Output("CONGNAMUL").
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2district_test.xml"); err != nil {
		t.Error(err)
	}
}

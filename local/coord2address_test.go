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

func TestCoord2AddressWithJSON(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if cr, err := local.CoordToAddress(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input(coord).
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}

}

func TestCoord2AddressWithSaveAsJSON(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if cr, err := local.CoordToAddress(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input(coord).
		FormatAs("json").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2address_test.json"); err != nil {
		t.Error(err)
	}

}

func TestCoord2AddressWithXML(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if cr, err := local.CoordToAddress(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input(coord).
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else {
		t.Log(cr)
	}

}

func TestCoord2AddressWithSaveAsXML(t *testing.T) {
	x := "127.423084873712"
	y := "37.0789561558879"
	coord := "WGS84"

	if cr, err := local.CoordToAddress(x, y).
		AuthorizeWith(common.REST_API_KEY).
		Input(coord).
		FormatAs("xml").
		Collect(); err != nil {
		t.Error(err)
	} else if err = cr.SaveAs("coord2address_test.xml"); err != nil {
		t.Error(cr)
	}

}

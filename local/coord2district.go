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

package local

import (
	"encoding/xml"
	"errors"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"strconv"

	"github.com/goccy/go-json"
)

// Region represents a document of a coordinate conversion result.
type Region struct {
	RegionType       string  `json:"region_type" xml:"region_type"`
	AddressName      string  `json:"address_name" xml:"address_name"`
	Region1depthName string  `json:"region_1depth_name" xml:"region_1depth_name"`
	Region2depthName string  `json:"region_2depth_name" xml:"region_2depth_name"`
	Region3depthName string  `json:"region_3depth_name" xml:"region_3depth_name"`
	Region4depthName string  `json:"region_4depth_name" xml:"region_4depth_name"`
	Code             string  `json:"code" xml:"code"`
	X                float64 `json:"x" xml:"x"`
	Y                float64 `json:"y" xml:"y"`
}

// CoordToDistrictResult represents a coordinate conversion result.
type CoordToDistrictResult struct {
	XMLName   xml.Name    `json:"-" xml:"result"`
	Meta      common.Meta `json:"meta" xml:"meta"`
	Documents []Region    `json:"documents" xml:"documents"`
}

// String implements fmt.Stringer.
func (cr CoordToDistrictResult) String() string { return common.String(cr) }

// SaveAs saves cr to @filename.
//
// The file extension could be either .json or .xml.
func (cr CoordToDistrictResult) SaveAs(filename string) error {
	return common.SaveAsJSONorXML(cr, filename)
}

// CoordToDistrictInitializer is a lazy coordinate converter.
type CoordToDistrictInitializer struct {
	X           string
	Y           string
	Format      string
	AuthKey     string
	InputCoord  string
	OutputCoord string
}

// CoordToDistrict converts the coordinates of @x and @y in the selected coordinate system
// into the administrative and legal-status area information.
//
// See https://developers.kakao.com/docs/latest/ko/local/dev-guide#coord-to-district for more details.
func CoordToDistrict(x, y float64) *CoordToDistrictInitializer {
	return &CoordToDistrictInitializer{
		X:           strconv.FormatFloat(x, 'f', -1, 64),
		Y:           strconv.FormatFloat(y, 'f', -1, 64),
		Format:      "json",
		AuthKey:     common.KeyPrefix,
		InputCoord:  "WGS84",
		OutputCoord: "WGS84",
	}
}

// FormatAs sets the request format to @format (json or xml).
func (ci *CoordToDistrictInitializer) FormatAs(format string) *CoordToDistrictInitializer {
	switch format {
	case "json", "xml":
		ci.Format = format
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ci
}

// AuthorizeWith sets the authorization key to @key.
func (ci *CoordToDistrictInitializer) AuthorizeWith(key string) *CoordToDistrictInitializer {
	ci.AuthKey = common.FormatKey(key)
	return ci
}

// Input sets the input coordinate system of ci to @coord.
//
// There are a few supported coordinate systems:
//
// WGS84
//
// WCONGNAMUL
//
// CONGNAMUL
//
// WTM
//
// TM
func (ci *CoordToDistrictInitializer) Input(coord string) *CoordToDistrictInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM":
		ci.InputCoord = coord
	default:
		panic(errors.New(
			`input coordinate system must be one of the following options:
			WGS84, WCONGNAMUL, CONGNAMUL, WTM, TM`))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ci
}

// Output sets the output coordinate system of ci to @coord.
//
// There are a few supported coordinate systems:
//
// WGS84
//
// WCONGNAMUL
//
// CONGNAMUL
//
// WTM
//
// TM
func (ci *CoordToDistrictInitializer) Output(coord string) *CoordToDistrictInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM":
		ci.OutputCoord = coord
	default:
		panic(errors.New(
			`output coordinate system must be one of the following options:
			WGS84, WCONGNAMUL, CONGNAMUL, WTM, TM`))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ci
}

// Collect returns the coordinate conversion result.
func (ci *CoordToDistrictInitializer) Collect() (res CoordToDistrictResult, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%sgeo/coord2regioncode.%s?x=%s&y=%s&input_coord=%s&output_coord=%s",
			prefix, ci.Format, ci.X, ci.Y, ci.InputCoord, ci.OutputCoord), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(common.Authorization, ci.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if ci.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if ci.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	return
}

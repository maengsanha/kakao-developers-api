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
	"internal/common"
)

// Place represents a place information of Local APIs.
type Place struct {
	Id                string `json:"id" xml:"id"`
	PlaceName         string `json:"place_name" xml:"place_name"`
	CategoryName      string `json:"category_name" xml:"category_name"`
	CategoryGroupCode string `json:"category_group_code" xml:"category_group_code"`
	CategoryGroupName string `json:"category_group_name" xml:"category_group_name"`
	Phone             string `json:"phone" xml:"phone"`
	AddressName       string `json:"address_name" xml:"address_name"`
	RoadAddressName   string `json:"road_address_name" xml:"road_address_name"`
	X                 string `json:"x" xml:"x"`
	Y                 string `json:"y" xml:"y"`
	PlaceURL          string `json:"place_url" xml:"place_url"`
	Distance          string `json:"distance" xml:"distance"`
}

// PlaceSearchResult represents a place search result.
type PlaceSearchResult struct {
	XMLName   xml.Name            `json:"-" xml:"result"`
	Meta      common.PageableMeta `json:"meta" xml:"meta"`
	Documents []Place             `json:"documents" xml:"documents"`
}

// String implements fmt.Stringer.
func (pr PlaceSearchResult) String() string { return common.String(pr) }

type PlaceSearchResults []PlaceSearchResult

// SaveAs saves prs to @filename.
//
// The file extension could be either .json or .xml.
func (prs PlaceSearchResults) SaveAs(filename string) error {
	return common.SaveAsJSONorXML(prs, filename)
}

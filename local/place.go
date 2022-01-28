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

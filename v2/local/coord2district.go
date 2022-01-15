// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
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
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []Region `json:"documents" xml:"documents"`
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
func CoordToDistrict(x, y string) *CoordToDistrictInitializer {
	return &CoordToDistrictInitializer{
		X:           x,
		Y:           y,
		Format:      "json",
		AuthKey:     "KakaoAK ",
		InputCoord:  "WGS84",
		OutputCoord: "WGS84",
	}
}

func (c *CoordToDistrictInitializer) FormatJSON() *CoordToDistrictInitializer {
	c.Format = "json"
	return c
}

func (c *CoordToDistrictInitializer) FormatXML() *CoordToDistrictInitializer {
	c.Format = "xml"
	return c
}

// AuthorizeWith sets the authorization key to @key.
func (c *CoordToDistrictInitializer) AuthorizeWith(key string) *CoordToDistrictInitializer {
	c.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return c
}

// Input sets the input coordinate system of c to @coord.
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
func (c *CoordToDistrictInitializer) Input(coord string) *CoordToDistrictInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM":
		c.InputCoord = coord
	}
	return c
}

// Output sets the output coordinate system of c to @coord.
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
func (c *CoordToDistrictInitializer) Output(coord string) *CoordToDistrictInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM":
		c.OutputCoord = coord
	}
	return c
}

// Collect returns the coordinate conversion result.
func (c *CoordToDistrictInitializer) Collect() (res CoordToDistrictResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/coord2regioncode.%s?x=%s&y=%s&input_coord=%s&output_coord=%s",
			c.Format, c.X, c.Y, c.InputCoord, c.OutputCoord), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", c.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if c.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if c.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	return
}

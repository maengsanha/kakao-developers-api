// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

// TotalAddress ...
type TotalAddress struct {
	Address struct {
		AddressName      string `json:"address_name" xml:"address_name"`
		Region1depthName string `json:"region_1depth_name" xml:"region_1depth_name"`
		Region2depthName string `json:"region_2depth_name" xml:"region_2depth_name"`
		Region3depthName string `json:"region_3depth_name" xml:"region_3depth_name"`
		MountainYN       string `json:"mountain_yn" xml:"mountain_yn"`
		MainAddressNo    string `json:"main_address_no" xml:"main_address_no"`
		SubAddressNo     string `json:"sub_address_no" xml:"sub_address_no"`
		ZipCode          string `json:"zip_code" xml:"zip_code"`
	} `json:"address" xml:"address"`
	RoadAddress struct {
		AddressName      string `json:"address_name" xml:"address_name"`
		Region1depthName string `json:"region_1depth_name" xml:"region_1depth_name"`
		Region2depthName string `json:"region_2depth_name" xml:"region_2depth_name"`
		Region3depthName string `json:"region_3depth_name" xml:"region_3depth_name"`
		RoadName         string `json:"road_name" xml:"road_name"`
		UndergroundYN    string `json:"underground_yn" xml:"underground_yn"`
		MainBuildingNo   string `json:"main_building_no" xml:"main_building_no"`
		SubBuildingNo    string `json:"sub_building_no" xml:"sub_building_no"`
		BuildingName     string `json:"building_name" xml:"building_name"`
		ZoneNo           string `json:"zone_no" xml:"zone_no"`
	} `json:"road_address" xml:"road_address"`
}

// CoordToAddressResult ...
type CoordToAddressResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []TotalAddress `json:"documents" xml:"documents"`
}

// CoordToAddressInitializer ...
type CoordToAddressInitializer struct {
	X          string
	Y          string
	Format     string
	AuthKey    string
	InputCoord string
}

// CoordToAddress ...
func CoordToAddress(x, y string) *CoordToAddressInitializer {
	return &CoordToAddressInitializer{
		X:          x,
		Y:          y,
		Format:     "json",
		AuthKey:    "KakaoAK ",
		InputCoord: "WGS84",
	}
}

func (c *CoordToAddressInitializer) FormatJSON() *CoordToAddressInitializer {
	c.Format = "json"
	return c
}

func (c *CoordToAddressInitializer) FormatXML() *CoordToAddressInitializer {
	c.Format = "xml"
	return c
}

// AuthorizeWith ...
func (c *CoordToAddressInitializer) AuthorizeWith(key string) *CoordToAddressInitializer {
	c.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return c
}

// Input ...
func (c *CoordToAddressInitializer) Input(coord string) *CoordToAddressInitializer {
	if coord == "WGS84" || coord == "WCONGNAMUL" || coord == "CONGNAMUL" || coord == "WTM" || coord == "TM" {
		c.InputCoord = coord
	}
	return c
}

// Collect ...
func (c *CoordToAddressInitializer) Collect() (res CoordToAddressResult, err error) {
	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/coord2address.%s?x=%s&y=%s&input_coord=%s", c.Format, c.X, c.Y, c.InputCoord), nil)
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

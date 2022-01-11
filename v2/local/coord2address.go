// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type address struct {
	AddressName      string `json:"address_name" xml:"address_name"`
	Region1depthName string `json:"region_1depth_name" xml:"region_1depth_name"`
	Region2depthName string `json:"region_2depth_name" xml:"region_2depth_name"`
	Region3depthName string `json:"region_3depth_name" xml:"region_3depth_name"`
	MountainYN       string `json:"mountain_yn" xml:"mountain_yn"`
	MainAddressNo    string `json:"main_address_no" xml:"main_address_no"`
	SubAddressNo     string `json:"sub_address_no" xml:"sub_address_no"`
	ZipCode          string `json:"zip_code" xml:"zip_code"`
	X                string `json:"x" xml:"x"`
	Y                string `json:"y" xml:"y"`
}

type roadAddress struct {
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
}

type TotalAddress struct {
	Address     address     `json:"address" xml:"address"`
	RoadAddress roadAddress `json:"road_address" xml:"road_address"`
}

type CoordToAddressResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []TotalAddress `json:"documents" xml:"documents"`
}

type CoordToAddressInitializer struct {
	X          string
	Y          string
	Format     string
	AuthKey    string
	InputCoord string
}

func CoordToAddress(x, y string) *CoordToAddressInitializer {
	return &CoordToAddressInitializer{
		X:          x,
		Y:          y,
		Format:     "json",
		AuthKey:    "KakaoAK ",
		InputCoord: "WGS84",
	}
}

func (b *CoordToAddressInitializer) FormatJSON() *CoordToAddressInitializer {
	b.Format = "json"
	return b
}

func (b *CoordToAddressInitializer) FormatXML() *CoordToAddressInitializer {
	b.Format = "xml"
	return b
}

func (b *CoordToAddressInitializer) AuthorizeWith(key string) *CoordToAddressInitializer {
	b.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return b
}

func (b *CoordToAddressInitializer) RequestWGS84() *CoordToAddressInitializer {
	b.InputCoord = "WGS84"
	return b
}

func (b *CoordToAddressInitializer) RequestWCONGNAMUL() *CoordToAddressInitializer {
	b.InputCoord = "WCONGNAMUL"
	return b
}

func (b *CoordToAddressInitializer) RequestCONGNAMUL() *CoordToAddressInitializer {
	b.InputCoord = "CONGNAMUL"
	return b
}

func (b *CoordToAddressInitializer) RequestWTM() *CoordToAddressInitializer {
	b.InputCoord = "WTM"
	return b
}

func (b *CoordToAddressInitializer) RequestTM() *CoordToAddressInitializer {
	b.InputCoord = "TM"
	return b
}

func (b *CoordToAddressInitializer) Collect() (res CoordToAddressResult, err error) {
	client := new(http.Client)

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/coord2address.%s?x=%s&y=%s&input_coord=%s", b.Format, b.X, b.Y, b.InputCoord), nil)
	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", b.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if b.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if b.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	return

}

// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// TotalAddress represents a document of CoordToAddressResult.
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

// CoordToAddressResult represents a CoordToAddress result.
type CoordToAddressResult struct {
	XMLName xml.Name `json:"-" xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []TotalAddress `json:"documents" xml:"documents"`
}

// String implements fmt.Stringer.
func (cr CoordToAddressResult) String() string {
	bs, _ := json.MarshalIndent(cr, "", "  ")
	return string(bs)
}

// SaveAs saves cr to @filename.
//
// The file extension could be either .json or .xml.
func (cr CoordToAddressResult) SaveAs(filename string) error {
	switch tokens := strings.Split(filename, "."); tokens[len(tokens)-1] {
	case "json":
		if bs, err := json.MarshalIndent(cr, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	case "xml":
		if bs, err := xml.MarshalIndent(cr, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	default:
		return ErrUnsupportedFormat
	}
}

// CoordToAddressInitializer is a lazy coord to address converter.
type CoordToAddressInitializer struct {
	X          string
	Y          string
	Format     string
	AuthKey    string
	InputCoord string
}

// CoordToAddress converts the @x and @y coordinates of location in the selected coordinate system
// to land-lot number address(with post number) and road name address.
//
// Details can be referred to
// https://developers.kakao.com/docs/latest/ko/local/dev-guide#coord-to-address.
func CoordToAddress(x, y string) *CoordToAddressInitializer {
	return &CoordToAddressInitializer{
		X:          x,
		Y:          y,
		Format:     "json",
		AuthKey:    "KakaoAK ",
		InputCoord: "WGS84",
	}
}

// FormatAs sets the request format to @format (json or xml).
func (ci *CoordToAddressInitializer) FormatAs(format string) *CoordToAddressInitializer {
	switch format {
	case "json", "xml":
		ci.Format = format
	default:
		panic(ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ci
}

// AuthorizeWith sets the authorization key to @key.
func (ci *CoordToAddressInitializer) AuthorizeWith(key string) *CoordToAddressInitializer {
	ci.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return ci
}

// Input sets the coordinate system of request.
//
// There are following coordinate system exist:
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
func (ci *CoordToAddressInitializer) Input(coord string) *CoordToAddressInitializer {
	switch coord {
	case "WGS84", "WCONAMUL", "CONGNAMUL", "WTM", "TM":
		ci.InputCoord = coord
	default:
		panic(errors.New("input coordinate system must be one of following options:\n WGS84, WCONGNAMUL, CONGNAMUL, WTM, TM"))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return ci
}

// Collect returns the land-lot number address(with post number) and road name address.
func (ci *CoordToAddressInitializer) Collect() (res CoordToAddressResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%sgeo/coord2address.%s?x=%s&y=%s&input_coord=%s",
			prefix, ci.Format, ci.X, ci.Y, ci.InputCoord), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", ci.AuthKey)

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

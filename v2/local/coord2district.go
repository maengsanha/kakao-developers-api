// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

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

type Coord2DistrictResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []Region `json:"documents" xml:"documents"`
}

type Coord2DistrictInitializer struct {
	X           string
	Y           string
	Format      string
	AuthKey     string
	InputCoord  string
	OutputCoord string
}

func Coord2District(x, y string) *Coord2DistrictInitializer {
	return &Coord2DistrictInitializer{
		X:           x,
		Y:           y,
		Format:      "json",
		AuthKey:     "KakaoAK ",
		InputCoord:  "WGS84",
		OutputCoord: "WGS84",
	}
}

func (c *Coord2DistrictInitializer) FormatJSON() *Coord2DistrictInitializer {
	c.Format = "json"
	return c
}

func (c *Coord2DistrictInitializer) FormatXML() *Coord2DistrictInitializer {
	c.Format = "xml"
	return c
}

func (c *Coord2DistrictInitializer) AuthorizeWith(key string) *Coord2DistrictInitializer {
	c.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return c
}

func (c *Coord2DistrictInitializer) RequestWGS84() *Coord2DistrictInitializer {
	c.InputCoord = "WGS84"
	return c
}

func (c *Coord2DistrictInitializer) RequestWCONGNAMUL() *Coord2DistrictInitializer {
	c.InputCoord = "WCONGNAMUL"
	return c
}

func (c *Coord2DistrictInitializer) RequestCONGNAMUL() *Coord2DistrictInitializer {
	c.InputCoord = "CONGNAMUL"
	return c
}

func (c *Coord2DistrictInitializer) RequestWTM() *Coord2DistrictInitializer {
	c.InputCoord = "WTM"
	return c
}

func (c *Coord2DistrictInitializer) RequestTM() *Coord2DistrictInitializer {
	c.InputCoord = "TM"
	return c
}

func (c *Coord2DistrictInitializer) DisplayWGS84() *Coord2DistrictInitializer {
	c.OutputCoord = "WGS84"
	return c
}

func (c *Coord2DistrictInitializer) DisplayWCONGNAMUL() *Coord2DistrictInitializer {
	c.OutputCoord = "WCONGNAMUL"
	return c
}

func (c *Coord2DistrictInitializer) DisplayCONGNAMUL() *Coord2DistrictInitializer {
	c.OutputCoord = "CONGNAMUL"
	return c
}

func (c *Coord2DistrictInitializer) DisplayWTM() *Coord2DistrictInitializer {
	c.OutputCoord = "WTM"
	return c
}

func (c *Coord2DistrictInitializer) DisplayTM() *Coord2DistrictInitializer {
	c.OutputCoord = "TM"
	return c
}

func (c *Coord2DistrictInitializer) Collect() (res Coord2DistrictResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/coord2regioncode.%s?x=%s&y=%s&input_coord=%s&output_coord=%s", c.Format, c.X, c.Y, c.InputCoord, c.OutputCoord), nil)
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

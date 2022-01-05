// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

const (
	WGS84      = "WGS84"
	WCONGNAMUL = "WCONGNAMUL"
	CONGNAMUL  = "CONGNAMUL"
	WTM        = "WTM"
	TM         = "TM"
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
		Format:      JSON,
		AuthKey:     keyPrefix,
		InputCoord:  WGS84,
		OutputCoord: WGS84,
	}
}

func (c *Coord2DistrictInitializer) As(format string) *Coord2DistrictInitializer {
	if format == JSON || format == XML {
		c.Format = format
	}
	return c
}

func (c *Coord2DistrictInitializer) AuthorizeWith(key string) *Coord2DistrictInitializer {
	c.AuthKey = keyPrefix + strings.TrimSpace(key)
	return c
}

func (c *Coord2DistrictInitializer) Request(coord string) *Coord2DistrictInitializer {
	if coord == WGS84 || coord == WCONGNAMUL || coord == CONGNAMUL || coord == WTM || coord == TM {
		c.InputCoord = coord
	}
	return c
}

func (c *Coord2DistrictInitializer) Display(coord string) *Coord2DistrictInitializer {
	if coord == WGS84 || coord == WCONGNAMUL || coord == CONGNAMUL || coord == WTM || coord == TM {
		c.OutputCoord = coord
	}
	return c
}

func (c *Coord2DistrictInitializer) Collect() (res Coord2DistrictResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/coord2regioncode.%s?x=%s&y=%s&input_coord=%s&output_coord=%s", c.Format, c.X, c.Y, c.InputCoord, c.OutputCoord), nil)
	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(authorization, c.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if c.Format == JSON {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if c.Format == XML {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	return
}

// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type Coord struct {
	X float64 `json:"x" xml:"x"`
	Y float64 `json:"y" xml:"y"`
}

type TransCoordInitializer struct {
	X           float64
	Y           float64
	Format      string
	AuthKey     string
	InputCoord  string
	OutputCoord string
}

type TransCoordResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []Coord `json:"documents" xml:"documents"`
}

func TransCoord(x, y float64) *TransCoordInitializer {
	return &TransCoordInitializer{
		X:           x,
		Y:           y,
		Format:      JSON,
		AuthKey:     keyPrefix,
		InputCoord:  "WGS84",
		OutputCoord: "WGS84",
	}
}

func (t *TransCoordInitializer) As(format string) *TransCoordInitializer {
	if format == JSON || format == XML {
		t.Format = format
	}
	return t
}

func (t *TransCoordInitializer) AuthorizeWith(key string) *TransCoordInitializer {
	t.AuthKey = keyPrefix + strings.TrimSpace(key)
	return t
}

func (t *TransCoordInitializer) Request(coord string) *TransCoordInitializer {
	if coord == "WGS84" || coord == "WCONGNAMUL" || coord == "CONGNAMUL" || coord == "WTM" || coord == "TM" || coord == "KTM" || coord == "UTM" || coord == " BESSEL" || coord == "WKTM" || coord == "WUTM" {
		t.InputCoord = coord
	}
	return t
}

func (t *TransCoordInitializer) Display(coord string) *TransCoordInitializer {
	if coord == "WGS84" || coord == "WCONGNAMUL" || coord == "CONGNAMUL" || coord == "WTM" || coord == "TM" || coord == "KTM" || coord == "UTM" || coord == " BESSEL" || coord == "WKTM" || coord == "WUTM" {
		t.OutputCoord = coord
	}
	return t
}

func (t *TransCoordInitializer) Collect() (res TransCoordResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/transcoord.%s?x=%f&y=%f&input_coord=%s&output_coord=%s", t.Format, t.X, t.Y, t.InputCoord, t.OutputCoord), nil)
	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(authorization, t.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if t.Format == JSON {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if t.Format == XML {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	return
}

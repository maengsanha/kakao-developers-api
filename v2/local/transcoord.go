// Package local provides the features of the Local API.
package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

// Coord ...
type Coord struct {
	X float64 `json:"x" xml:"x"`
	Y float64 `json:"y" xml:"y"`
}

// TransCoordInitializer ...
type TransCoordInitializer struct {
	X           float64
	Y           float64
	Format      string
	AuthKey     string
	InputCoord  string
	OutputCoord string
}

// TransCoordResult ...
type TransCoordResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []Coord `json:"documents" xml:"documents"`
}

// TransCoord ...
func TransCoord(x, y float64) *TransCoordInitializer {
	return &TransCoordInitializer{
		X:           x,
		Y:           y,
		Format:      "json",
		AuthKey:     "KakaoAK ",
		InputCoord:  "WGS84",
		OutputCoord: "WGS84",
	}
}

func (t *TransCoordInitializer) FormatJSON() *TransCoordInitializer {
	t.Format = "json"
	return t
}

func (t *TransCoordInitializer) FormatXML() *TransCoordInitializer {
	t.Format = "xml"
	return t
}

// AuthorizeWith ...
func (t *TransCoordInitializer) AuthorizeWith(key string) *TransCoordInitializer {
	t.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return t
}

// Input ...
func (t *TransCoordInitializer) Input(coord string) *TransCoordInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM", "KTM", "UTM", "BESSEL", "WKTM", "WUTM":
		t.InputCoord = coord
	}
	return t
}

// Output ...
func (t *TransCoordInitializer) Output(coord string) *TransCoordInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM", "KTM", "UTM", "BESSEL", "WKTM", "WUTM":
		t.OutputCoord = coord
	}
	return t
}

// Collect ...
func (t *TransCoordInitializer) Collect() (res TransCoordResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/transcoord.%s?x=%f&y=%f&input_coord=%s&output_coord=%s",
			t.Format, t.X, t.Y, t.InputCoord, t.OutputCoord), nil)

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set("Authorization", t.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if t.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if t.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	return
}

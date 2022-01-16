package local

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Coord represents X and Y coordinates of Local APIs.
type Coord struct {
	X float64 `json:"x" xml:"x"`
	Y float64 `json:"y" xml:"y"`
}

// TransCoordInitializer initializes parameters that used in transform coordinate.
type TransCoordInitializer struct {
	X           string
	Y           string
	Format      string
	AuthKey     string
	InputCoord  string
	OutputCoord string
}

// TransCoordResult represents a converted result.
type TransCoordResult struct {
	XMLName xml.Name `xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []Coord `json:"documents" xml:"documents"`
}

// TransCoord converts the specified X and Y coordinates ​​in a coordinate system
// to another X and Y coordinates in the designated coordinate system.
//
// Details can be referred to
// https://developers.kakao.com/docs/latest/ko/local/dev-guide#trans-coord
func TransCoord(x, y float64) *TransCoordInitializer {
	return &TransCoordInitializer{
		X:           strconv.FormatFloat(x, 'f', -1, 64),
		Y:           strconv.FormatFloat(y, 'f', -1, 64),
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

// AuthorizeWith sets the authorization key to @key.
func (t *TransCoordInitializer) AuthorizeWith(key string) *TransCoordInitializer {
	t.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return t
}

// Input specifies Coordinate system of the requested x and y.
//
// There are a few supported coordinate systems(@coord):
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
//
// KTM
//
// UTM
//
// BESSEL
//
// WKTM
//
// WUTM
func (t *TransCoordInitializer) Input(coord string) *TransCoordInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM", "KTM", "UTM", "BESSEL", "WKTM", "WUTM":
		t.InputCoord = coord
	}
	return t
}

// Output specifies type of coordinate system to be converted from the input coordinate system.
//
// There are a few supported coordinate systems(@coord):
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
//
// KTM
//
// UTM
//
// BESSEL
//
// WKTM
//
// WUTM
func (t *TransCoordInitializer) Output(coord string) *TransCoordInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM", "KTM", "UTM", "BESSEL", "WKTM", "WUTM":
		t.OutputCoord = coord
	}
	return t
}

// Collect sends a GET request and
// returns the result of the coordinate system conversion.
func (t *TransCoordInitializer) Collect() (res TransCoordResult, err error) {
	// at first, send request to the API server
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/transcoord.%s?x=%s&y=%s&input_coord=%s&output_coord=%s",
			t.Format, t.X, t.Y, t.InputCoord, t.OutputCoord), nil)

	if err != nil {
		return
	}
	// don't forget to close the request for concurrent request
	req.Close = true

	// set authorization header
	req.Header.Set("Authorization", t.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// don't forget to close the response body
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

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
	"strconv"
	"strings"
)

// Coord represents a document of coordinate transformation result.
type Coord struct {
	X float64 `json:"x" xml:"x"`
	Y float64 `json:"y" xml:"y"`
}

// TransCoordInitializer is a lazy coordinate converter.
type TransCoordInitializer struct {
	X           string
	Y           string
	Format      string
	AuthKey     string
	InputCoord  string
	OutputCoord string
}

// TransCoordResult represents a coordinate transformation result.
type TransCoordResult struct {
	XMLName xml.Name `json:"-" xml:"result"`
	Meta    struct {
		TotalCount int `json:"total_count" xml:"total_count"`
	} `json:"meta" xml:"meta"`
	Documents []Coord `json:"documents" xml:"documents"`
}

// String implements fmt.Stringer.
func (tr TransCoordResult) String() string {
	bs, _ := json.MarshalIndent(tr, "", "  ")
	return string(bs)
}

// SaveAs saves tr to @filename.
//
// The file extension could be either .json or .xml.
func (tr TransCoordResult) SaveAs(filename string) error {
	switch tokens := strings.Split(filename, "."); tokens[len(tokens)-1] {
	case "json":
		if bs, err := json.MarshalIndent(tr, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	case "xml":
		if bs, err := xml.MarshalIndent(tr, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	default:
		return ErrUnsupportedFormat
	}
}

// TransCoord converts @x and @y coordinates to another X and Y coordinates in the designated coordinate system.
//
// Details can be referred to
// https://developers.kakao.com/docs/latest/ko/local/dev-guide#trans-coord.
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

// FormatAs sets the request format to @format (json or xml).
func (ti *TransCoordInitializer) FormatAs(format string) *TransCoordInitializer {
	switch format {
	case "json", "xml":
		ti.Format = format
	default:
		panic(ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ti
}

// AuthorizeWith sets the authorization key to @key.
func (ti *TransCoordInitializer) AuthorizeWith(key string) *TransCoordInitializer {
	ti.AuthKey = "KakaoAK " + strings.TrimSpace(key)
	return ti
}

// Input sets the type of input coordinate system.
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
func (ti *TransCoordInitializer) Input(coord string) *TransCoordInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM", "KTM", "UTM", "BESSEL", "WKTM", "WUTM":
		ti.InputCoord = coord
	default:
		panic(errors.New("input coordinate system must be one of the following options:\nWGS84, WCONGNAMUL, CONGNAMUL, WTM, TM, KTM, UTM, BESSEL, WKTM, WUTM"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ti
}

// Output sets the type of output coordinate system.
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
func (ti *TransCoordInitializer) Output(coord string) *TransCoordInitializer {
	switch coord {
	case "WGS84", "WCONGNAMUL", "CONGNAMUL", "WTM", "TM", "KTM", "UTM", "BESSEL", "WKTM", "WUTM":
		ti.OutputCoord = coord
	default:
		panic(errors.New("output coordinate system must be one of the following options:\nWGS84, WCONGNAMUL, CONGNAMUL, WTM, TM, KTM, UTM, BESSEL, WKTM, WUTM"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ti
}

// Collect returns the coordinate system conversion result.
func (ti *TransCoordInitializer) Collect() (res TransCoordResult, err error) {
	// at first, send request to the API server
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("https://dapi.kakao.com/v2/local/geo/transcoord.%s?x=%s&y=%s&input_coord=%s&output_coord=%s",
			ti.Format, ti.X, ti.Y, ti.InputCoord, ti.OutputCoord), nil)

	if err != nil {
		return
	}
	// don't forget to close the request for concurrent request
	req.Close = true

	// set authorization header
	req.Header.Set("Authorization", ti.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	// don't forget to close the response body
	defer resp.Body.Close()

	if ti.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if ti.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	return
}

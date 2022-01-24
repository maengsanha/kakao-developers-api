package local

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"strings"
)

// PlaceSearchResult represents a place search result.
type PlaceSearchResult struct {
	XMLName xml.Name `json:"-" xml:"result"`
	Meta    struct {
		TotalCount    int        `json:"total_count" xml:"total_count"`
		PageableCount int        `json:"pageable_count" xml:"pageable_count"`
		IsEnd         bool       `json:"is_end" xml:"is_end"`
		SameName      RegionInfo `json:"same_name" xml:"same_name"`
	} `json:"meta" xml:"meta"`
	Documents []Place `json:"documents" xml:"documents"`
}

// String implements fmt.Stringer.
func (pr PlaceSearchResult) String() string {
	bs, _ := json.MarshalIndent(pr, "", "  ")
	return string(bs)
}

type PlaceSearchResults []PlaceSearchResult

// SaveAs saves prs to @filename.
//
// The file extension could be either .json or .xml.
func (prs PlaceSearchResults) SaveAs(filename string) error {
	switch tokens := strings.Split(filename, "."); tokens[len(tokens)-1] {
	case "json":
		if bs, err := json.MarshalIndent(prs, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	case "xml":
		if bs, err := xml.MarshalIndent(prs, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0644)
		}
	default:
		return ErrUnsupportedFormat
	}
}

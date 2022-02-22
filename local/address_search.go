package local

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
)

// ComplexAddress represents a document of an address search result.
type ComplexAddress struct {
	AddressName string `json:"address_name" xml:"address_name"`
	AddressType string `json:"address_type" xml:"address_type"`
	X           string `json:"x" xml:"x"`
	Y           string `json:"y" xml:"y"`
	Address     struct {
		AddressName       string `json:"address_name" xml:"address_name"`
		Region1depthName  string `json:"region_1depth_name" xml:"region_1depth_name"`
		Region2depthName  string `json:"region_2depth_name" xml:"region_2depth_name"`
		Region3depthName  string `json:"region_3depth_name" xml:"region_3depth_name"`
		Region3depthHName string `json:"region_3depth_h_name" xml:"region_3depth_h_name"`
		HCode             string `json:"h_code" xml:"h_code"`
		BCode             string `json:"b_code" xml:"b_code"`
		MountainYN        string `json:"mountain_yn" xml:"mountain_yn"`
		MainAddressNo     string `json:"main_address_no" xml:"main_address_no"`
		SubAddressNo      string `json:"sub_address_no" xml:"sub_address_no"`
		ZipCode           string `json:"zip_code" xml:"zip_code"`
		X                 string `json:"x" xml:"x"`
		Y                 string `json:"y" xml:"y"`
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
		X                string `json:"x" xml:"x"`
		Y                string `json:"y" xml:"y"`
	} `json:"road_address" xml:"road_address"`
}

// AddressSearchResult represents an address search result.
type AddressSearchResult struct {
	XMLName   xml.Name            `json:"-" xml:"result"`
	Meta      common.PageableMeta `json:"meta" xml:"meta"`
	Documents []ComplexAddress    `json:"documents" xml:"documents"`
}

// String implements fmt.Stringer.
func (ar AddressSearchResult) String() string { return common.String(ar) }

type AddressSearchResults []AddressSearchResult

// SaveAs saves ars to @filename.
//
// The file extension could be either .json or .xml.
func (ars AddressSearchResults) SaveAs(filename string) error {
	return common.SaveAsJSONorXML(ars, filename)
}

// AddressSearchIterator is a lazy address search iterator.
type AddressSearchIterator struct {
	Query       string
	Format      string
	AuthKey     string
	AnalyzeType string
	Page        int
	Size        int
	end         bool
}

// AddressSearch provides the coordinates of the requested address with @query.
//
// See https://developers.kakao.com/docs/latest/ko/local/dev-guide#address-coord for more details.
func AddressSearch(query string) *AddressSearchIterator {
	return &AddressSearchIterator{
		Query:       url.QueryEscape(strings.TrimSpace(query)),
		Format:      "json",
		AuthKey:     common.KeyPrefix,
		AnalyzeType: "similar",
		Page:        1,
		Size:        10,
		end:         false,
	}
}

// FormatAs sets the request format to @format (json or xml).
func (it *AddressSearchIterator) FormatAs(format string) *AddressSearchIterator {
	switch format {
	case "json", "xml":
		it.Format = format
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// AuthorizeWith sets the authorization key to @key.
func (it *AddressSearchIterator) AuthorizeWith(key string) *AddressSearchIterator {
	it.AuthKey = common.FormatKey(key)
	return it
}

// Analyze sets the analyze type to @typ (similar or exact).
func (it *AddressSearchIterator) Analyze(typ string) *AddressSearchIterator {
	switch typ {
	case "similar", "exact":
		it.AnalyzeType = typ
	default:
		panic(errors.New("analyze type must be either similar or exact"))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Result sets the result page number (a value between 1 and 45).
func (it *AddressSearchIterator) Result(page int) *AddressSearchIterator {
	if 1 <= page && page <= 45 {
		it.Page = page
	} else {
		panic(common.ErrPageOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Display sets the number of documents displayed on a single page (a value between 1 and 30).
func (it *AddressSearchIterator) Display(size int) *AddressSearchIterator {
	if 1 <= size && size <= 30 {
		it.Size = size
	} else {
		panic(common.ErrSizeOutOfBound)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return it
}

// Next returns the address search result and proceeds the iterator to the next page.
func (it *AddressSearchIterator) Next() (res AddressSearchResult, err error) {
	// if there is no more result, return error
	if it.end {
		return res, Done
	}

	// at first, send request to the API server
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%ssearch/address.%s?query=%s&analyze_type=%s&page=%d&size=%d",
			prefix, it.Format, it.Query, it.AnalyzeType, it.Page, it.Size), nil)

	if err != nil {
		return
	}
	// don't forget to close the request for concurrent request
	req.Close = true

	// set authorization header
	req.Header.Set(common.Authorization, it.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	// don't forget to close the response body
	defer resp.Body.Close()

	if it.Format == "json" {
		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	} else if it.Format == "xml" {
		if err = xml.NewDecoder(resp.Body).Decode(&res); err != nil {
			return
		}
	}

	it.end = res.Meta.IsEnd || 45 < it.Page

	it.Page++

	return
}

// CollectAll collects all the remaining address search results.
func (it *AddressSearchIterator) CollectAll() (results AddressSearchResults) {
	var (
		items  = make(AddressSearchResults, 46-it.Page)
		errors = make([]error, 46-it.Page)
		wg     sync.WaitGroup
	)

	// FIXME: it needs pre-profiling
	for page := it.Page; page <= 45; page++ {
		wg.Add(1)
		go func(page int) {
			defer wg.Done()
			worker := *it
			items[page-it.Page], errors[page-it.Page] = worker.Result(page).Next()
		}(page)
	}
	wg.Wait()

	// filter out
	for idx, err := range errors {
		if err == nil {
			results = append(results, items[idx])
		}
	}

	it.end = true

	return
}

// Package local provides the features of the Local API.
package local

import "fmt"

type Address struct {
	AddressName       string `json:"address_name"`
	Region1depthName  string `json:"region_1depth_name"`
	Region2depthName  string `json:"region_2depth_name"`
	Region3depthName  string `json:"region_3depth_name"`
	Region3depthHName string `json:"region_3depth_h_name"`
	HCode             string `json:"h_code"`
	BCode             string `json:"b_code"`
	MountainYN        string `json:"mountain_yn"`
	MainAddressNo     string `json:"main_address_no"`
	SubAddressNo      string `json:"sub_address_no"`
	ZipCode           string `json:"zip_code"`
	X                 string `json:"x"`
	Y                 string `json:"y"`
}

type RoadAddress struct {
	AddressName      string `json:"address_name"`
	Region1depthName string `json:"region_1depth_name"`
	Region2depthName string `json:"region_2depth_name"`
	Region3depthName string `json:"region_3depth_name"`
	RoadName         string `json:"road_name"`
	UndergroundYN    string `json:"underground_yn"`
	MainBuildingNo   string `json:"main_building_no"`
	SubBuildingNo    string `json:"sub_building_no"`
	BuildingName     string `json:"building_name"`
	ZoneNo           string `json:"zone_no"`
	X                string `json:"x"`
	Y                string `json:"y"`
}

type AddressSearchResponse struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	}
	Documents []struct {
		AddressName string      `json:"address_name"`
		AddressType string      `json:"address_type"`
		X           string      `json:"x"`
		Y           string      `json:"y"`
		Address     Address     `json:"address"`
		RoadAddress RoadAddress `json:"road_address"`
	}
}

type AddressSearcher struct {
	Query       string
	AnalyzeType string
	Size        int
	Format      string
	AuthKey     string
}

func NewAddressSearcher() *AddressSearcher {
	return &AddressSearcher{
		AnalyzeType: "similar",
		Size:        10,
		Format:      "json",
		AuthKey:     "KakaoAK ",
	}
}

func (as *AddressSearcher) SetQuery(query string) *AddressSearcher {
	as.Query = query
	return as
}

func (as *AddressSearcher) SetAnalyzeType(typ string) *AddressSearcher {
	switch typ {
	case "similar", "exact":
		as.AnalyzeType = typ
	}
	return as
}

func (as *AddressSearcher) SetSize(size int) *AddressSearcher {
	if 1 <= size && size <= 30 {
		as.Size = size
	}
	return as
}

func (as *AddressSearcher) SetFormat(format string) *AddressSearcher {
	switch format {
	case "JSON", "json":
		as.Format = "json"
	case "XML", "xml":
		as.Format = "xml"
	}
	return as
}

func (as *AddressSearcher) SetAuth(key string) *AddressSearcher {
	as.AuthKey = fmt.Sprintf("KakaoAK %s", key)
	return as
}

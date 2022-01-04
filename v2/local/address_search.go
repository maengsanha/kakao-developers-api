// Package local provides the features of the Local API.
package local

// Address represents a detailed information of Land-lot address.
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

// RoadAddress represents a detailed information of Road name address.
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

// AddressSearchResponse represents a Address search response.
type AddressSearchResponse struct {
	Meta struct {
		TotalCount    int  `json:"total_count"`
		PageableCount int  `json:"pageable_count"`
		IsEnd         bool `json:"is_end"`
	} `json:"meta"`
	Documents []struct {
		AddressName string `json:"address_name"`
		AddressType string `json:"address_type"`
		X           string `json:"x"`
		Y           string `json:"y"`
		Address     `json:"address"`
		RoadAddress `json:"road_address"`
	} `json:"documents"`
}

// AddressSearchIterator
type AddressSearchIterator struct {
	Query       string
	AnalyzeType string
	Page        int
	Size        int
}

func AddressSearch(query string) *AddressSearchIterator {
	return &AddressSearchIterator{
		Query:       query,
		AnalyzeType: "similar",
		Page:        1,
		Size:        10,
	}
}

func (a *AddressSearchIterator) Analyze(typ string) *AddressSearchIterator {
	a.AnalyzeType = typ
	return a
}

func (a *AddressSearchIterator) Result(page int) *AddressSearchIterator {
	a.Page = page
	return a
}

func (a *AddressSearchIterator) Display(size int) *AddressSearchIterator {
	a.Size = size
	return a
}

func (a *AddressSearchIterator) Next() (AddressSearchResponse, error) {
	a = nil
	return AddressSearchResponse{}, nil
}

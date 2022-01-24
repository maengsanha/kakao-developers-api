package local

// RegionInfo represents a region information of Local APIs.
type RegionInfo struct {
	Region         []string `json:"region" xml:"region"`
	Keyword        string   `json:"keyword" xml:"keyword"`
	SelectedRegion string   `json:"selected_region" xml:"selected_region"`
}

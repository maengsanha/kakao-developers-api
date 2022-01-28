package common

// Meta is the commonly used metadata in Kakao Developers.
type Meta struct {
	TotalCount int `json:"total_count" xml:"total_count"`
}

// PageableMeta is the commonly used metadata in Kakao Developers, with pageable state.
type PageableMeta struct {
	Meta
	PageableCount int  `json:"pageable_count" xml:"pageable_count"`
	IsEnd         bool `json:"is_end" xml:"is_end"`
}

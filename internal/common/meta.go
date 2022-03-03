// Copyright 2022 Sanha Maeng, Soyang Baek, Jinmyeong Kim
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

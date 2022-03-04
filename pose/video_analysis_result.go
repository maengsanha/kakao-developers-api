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

package pose

import (
	"fmt"
	"internal/common"
	"net/http"

	"github.com/goccy/go-json"
)

// CheckVideoResult represents the result of CheckVideo.
type CheckVideoResult struct {
	JobId       string       `json:"job_id"`
	Status      string       `json:"status"`
	Annotations []Annotation `json:"annotations"`
	Categories  []Category   `json:"categories"`
	Info        Info         `json:"info"`
	Video       Video        `json:"video"`
	Description string       `json:"description"`
}

// String implements fmt.Stringer.
func (cr CheckVideoResult) String() string { return common.String(cr) }

// SaveAs saves cr to @filename.
//
// The file extension could be .json.
func (cr CheckVideoResult) SaveAs(filename string) error { return common.SaveAsJSONorXML(cr, filename) }

// Annotation containing the coordinates and score of the key points detected in each frame,
// as an array with the size of the number of frames.
type Annotation struct {
	FrameNum int      `json:"frame_num"`
	Objects  []Person `json:"objects"`
}

// Person represents each person's 17 key points(person's eyes, nose, shoulders, elbows, wrists, pelvis, knees, and ankles).
type Person struct {
	Area       float64   `json:"area"`
	BBox       []float64 `json:"bbox"`
	CategoryId int       `json:"category_id"`
	KeyPoints  []float64 `json:"keypoints"`
	Score      float64   `json:"score"`
}

// Category contains the information about key points.
type Category struct {
	Id             int      `json:"id"`
	Keypoints      []string `json:"keypoints"`
	Name           string   `json:"name"`
	Skeleton       [][]int  `json:"skeleton"`
	SuperCathegory string   `json:"supercathegory"`
}

// Info contains information about the analyzed video such as version, creation date, URL, description, etc.
type Info struct {
	Contributer string  `json:"contributer"`
	DateCreated string  `json:"date_created"`
	Description string  `json:"description"`
	URL         string  `json:"url"`
	Version     float32 `json:"version"`
	Year        int     `json:"year"`
}

// Video containing information about the frames of the requested video, such as the number of frames per second,
// the total number of frames, the video frame size.
type Video struct {
	FPS    float32 `json:"fps"`
	Frames int     `json:"frames"`
	Height int     `json:"height"`
	Width  int     `json:"width"`
}

// CheckVideoInitalizer is a lazy video checker.
type CheckVideoInitializer struct {
	AuthKey string
	JobId   string
}

// CheckVideo returns the processing status and the video analysis results processed through the analyze_video API.
//
// For more details visit https://developers.kakao.com/docs/latest/en/pose/dev-guide#job-retrieval.
func CheckVideo(source string) *CheckVideoInitializer {
	return &CheckVideoInitializer{
		AuthKey: common.KeyPrefix,
		JobId:   source,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ci *CheckVideoInitializer) AuthorizeWith(key string) *CheckVideoInitializer {
	ci.AuthKey = common.FormatKey(key)
	return ci
}

// Collect returns the check video result.
func (ci *CheckVideoInitializer) Collect() (res CheckVideoResult, err error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/job/%s", prefix, ci.JobId), nil)
	if err != nil {
		return
	}

	req.Close = true
	req.Header.Set(common.Authorization, ci.AuthKey)

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	return
}

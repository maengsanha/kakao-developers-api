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
	"bytes"
	"encoding/json"
	"fmt"
	"internal/common"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// AnalyzeImageeResult represents a result of image analyze result.
type AnalyzeImageResult []struct {
	Area       float64   `json:"area"`
	BBox       []float64 `json:"bbox"`
	CategoryId int       `json:"category_id"`
	KeyPoints  []float64 `json:"keypoints"`
	Score      float64   `json:"score"`
}

// String implements fmt.Stringer.
func (ar AnalyzeImageResult) String() string { return common.String(ar) }

// SaveAs saves ar to @filename.
//
// The file extension could be .json.
func (ar AnalyzeImageResult) SaveAs(filename string) error {
	return common.SaveAsJSONorXML(ar, filename)
}

// AnalyzeImageInitializer is a lazy image analyzer.
type AnalyzeImageInitializer struct {
	AuthKey  string
	ImageURL string
	Filename string
	withFile bool
}

// AnalyzeImage detects people in the given image and extracts each person's 17 key points(person's eyes, nose, shoulders,
// elbows, wrists, pelvis, knees, and ankles) to determine their pose.
//
// For more details visit https://developers.kakao.com/docs/latest/en/pose/dev-guide#image-pose-estimation.
func AnalyzeImage() *AnalyzeImageInitializer {
	return &AnalyzeImageInitializer{
		AuthKey: common.KeyPrefix,
	}
}

// WithURL sets url to @url.
func (ai *AnalyzeImageInitializer) WithURL(url string) *AnalyzeImageInitializer {
	ai.ImageURL = url
	ai.withFile = false
	return ai
}

// WithFile sets image path to @filename.
func (ai *AnalyzeImageInitializer) WithFile(filename string) *AnalyzeImageInitializer {
	ai.Filename = filename
	ai.withFile = true
	return ai
}

// AuthorizeWith sets the authorization key to @key.
func (ai *AnalyzeImageInitializer) AuthorizeWith(key string) *AnalyzeImageInitializer {
	ai.AuthKey = common.FormatKey(key)
	return ai
}

// Collect returns the image analyze result.
func (ai *AnalyzeImageInitializer) Collect() (res AnalyzeImageResult, err error) {
	var req *http.Request
	if ai.withFile {
		file, err := os.Open(ai.Filename)
		if err != nil {
			return res, err
		}

		if stat, err := file.Stat(); err != nil {
			return res, err
		} else if 2*1024*1024 < stat.Size() {
			return res, common.ErrTooLargeFile
		}

		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("file", ai.Filename)
		if err != nil {
			return res, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return res, err
		}

		writer.Close()

		req, err = http.NewRequest(http.MethodPost, prefix, body)
		if err != nil {
			return res, err
		}

		req.Header.Add("Content-Type", writer.FormDataContentType())
	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s?image_url=%s", prefix, ai.ImageURL), nil)
		if err != nil {
			return res, err
		}
	}

	req.Close = true
	req.Header.Add(common.Authorization, ai.AuthKey)

	client := &http.Client{}
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

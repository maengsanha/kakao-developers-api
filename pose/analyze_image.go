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
	"path/filepath"
)

// AnalyzImageeResult represents a result of image analyze result.
type AnalyzeImageResult []struct {
	Area       float64   `json:"area"`
	BBox       []float64 `json:"bbox"`
	CategoryId int       `json:"category_id"`
	KeyPoints  []float64 `json:"keypoints"`
	Score      float64   `json:"score"`
}

// String implements fmt.Stringer.
func (ar AnalyzeImageResult) String() string { return common.String(ar) }

// SaveAs saves ir to @filename.
//
// The file extension could be .json.
func (ar AnalyzeImageResult) SaveAs(filename string) error {
	return common.SaveAsJSONorXML(ar, filename)
}

// AnalyzeImageInitializer is a lazy image analyzer.
type AnalyzeImageInitializer struct {
	AuthKey  string
	ImageURL string
	File     *os.File
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

// WithURL sets url to @imageURL.
func (ai *AnalyzeImageInitializer) WithURL(url string) *AnalyzeImageInitializer {
	ai.ImageURL = url
	return ai
}

// WithFile sets filepath to @File.
func (ai *AnalyzeImageInitializer) WithFile(filepath string) *AnalyzeImageInitializer {
	bs, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	if stat, _ := bs.Stat(); stat.Size() > 2*1024*1024 {
		panic("up to 2MB are allowed")
	} else {
		ai.File = bs
		return ai
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ii *AnalyzeImageInitializer) AuthorizeWith(key string) *AnalyzeImageInitializer {
	ii.AuthKey = common.FormatKey(key)
	return ii
}

// Collect returns the image analyze result.
func (ii *AnalyzeImageInitializer) Collect() (res AnalyzeImageResult, err error) {
	client := new(http.Client)
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	if err != nil {
		return
	}

	if ii.File != nil {
		part, err := writer.CreateFormFile("file", filepath.Base(ii.File.Name()))
		if err != nil {
			return res, err
		}
		io.Copy(part, ii.File)
	}
	defer writer.Close()

	var req *http.Request
	if ii.File != nil {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s?file=%s", prefix, filepath.Base(ii.File.Name())), body)
	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s?image_url=%s", prefix, ii.ImageURL), nil)
	}
	if err != nil {
		return
	}
	fmt.Println(req.ContentLength)
	req.Close = true

	req.Header.Set(common.Authorization, ii.AuthKey)
	req.Header.Set("Content-type", "multipart/form-data")

	defer ii.File.Close()

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

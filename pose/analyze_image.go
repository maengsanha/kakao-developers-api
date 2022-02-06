package pose

import (
	"encoding/json"
	"fmt"
	"internal/common"
	"io/ioutil"
	"net/http"
)

// ImageAnalyzeResult represents a result of image analyze result.
type ImageAnalyzeResult struct {
	Person []struct {
		Area       float64   `json:"area"`
		Bbox       []float64 `json:"bbox"`
		CathgoryID int       `json:"category_id"`
		KeyPoints  []float64 `json:"keypoints"`
		Score      float64   `json:"score"`
	} `json:" "`
}

// String implements fmt.Stringer.
func (ir ImageAnalyzeResult) String() string { return common.String(ir) }

// SaveAs saves ir to @filename.
//
// The file extension could be .json.
func (ir ImageAnalyzeResult) SaveAs(filename string) error {
	return common.SaveAsJSONorXML(ir, filename)
}

// ImageAnalyzeInitializer is a lazy image analyzer.
type ImageAnalyzeInitializer struct {
	AuthKey   string
	ImageURL  string
	ImageFile []byte
}

// For more details visit https://developers.kakao.com/docs/latest/en/pose/dev-guide#image-pose-estimation.
func ImageAnalyze(source string) *ImageAnalyzeInitializer {
	if source[0:4] == "http" {
		println("1")
		return &ImageAnalyzeInitializer{
			AuthKey:  common.KeyPrefix,
			ImageURL: source,
		}
	} else {
		println("2")
		bs, err := ioutil.ReadFile(source)
		if err != nil {
			panic(err)
		}
		return &ImageAnalyzeInitializer{
			AuthKey:   common.KeyPrefix,
			ImageFile: bs,
		}
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ii *ImageAnalyzeInitializer) AuthorizeWith(key string) *ImageAnalyzeInitializer {
	ii.AuthKey = common.FormatKey(key)
	return ii
}

func (ii *ImageAnalyzeInitializer) Collect() (res ImageAnalyzeResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s?image_url=%s", prefix, ii.ImageURL), nil)
	println("%s?image_url=%s", prefix, ii.ImageURL)
	println()
	if err != nil {
		return
	}

	req.Close = true
	req.Header.Set(common.Authorization, ii.AuthKey)
	req.Header.Set("Content-type", "application/x-www-form-urlencoded")

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

package vision

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

// MultiTagResult represents a document of a Multi-tag creation result.
type MultiTagResult struct {
	Label   []string `json:"label"`
	LabelKr []string `json:"label_kr"`
}

// MultiTagCreateResult represents a Multi-tag creation result.
type MultiTagCreateResult struct {
	RID    string         `json:"rid"`
	Result MultiTagResult `json:"result"`
}

// String implements fmt.Stringer.
func (mr MultiTagCreateResult) String() string { return common.String(mr) }

// SaveAs saves mr to @filename.
//
// The file extension must be .json.
func (mr MultiTagCreateResult) SaveAs(filename string) error {
	return common.SaveAsJSON(mr, filename)
}

// MultiTagCreateInitializer is a lazy Multi-tag creator.
type MultiTagCreateInitializer struct {
	AuthKey  string
	Image    *os.File
	ImageURL string
}

// MultiTagCreate creates a tag according to image content.
//
// Image can be either image URL or image file (JPG or PNG).
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-multi-tag for more details.
func MultiTagCreate() *MultiTagCreateInitializer {
	return &MultiTagCreateInitializer{
		AuthKey:  common.KeyPrefix,
		Image:    nil,
		ImageURL: "",
	}
}

// WithFile sets the file to request on @filepath.
func (mi *MultiTagCreateInitializer) WithFile(filepath string) *MultiTagCreateInitializer {
	mi.Image = OpenFile(filepath)
	return mi
}

// WithURL sets the URL to request to @url.
func (mi *MultiTagCreateInitializer) WithURL(url string) *MultiTagCreateInitializer {
	mi.ImageURL = url
	return mi
}

// AuthorizeWith sets the authorization key to @key.
func (mi *MultiTagCreateInitializer) AuthorizeWith(key string) *MultiTagCreateInitializer {
	mi.AuthKey = common.FormatKey(key)
	return mi
}

// Collect returns the Multi-tag creation result.
func (mi *MultiTagCreateInitializer) Collect() (res MultiTagCreateResult, err error) {
	client := new(http.Client)
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if mi.Image != nil {
		part, err := writer.CreateFormFile("image", filepath.Base(mi.Image.Name()))
		if err != nil {
			return res, err
		}
		io.Copy(part, mi.Image)
	}
	defer writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/multitag/generate?image_url=%s", prefix, mi.ImageURL), body)
	if err != nil {
		return res, err
	}
	req.Close = true

	req.Header.Set(common.Authorization, mi.AuthKey)
	if mi.Image != nil {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	defer mi.Image.Close()

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, err
	}
	return
}

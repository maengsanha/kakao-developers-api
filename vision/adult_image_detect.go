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

// AdultResult represents a document of a detected adult image result.
// If the soft or adult score of an image is high, it is likely to be nudity or porn images. (normal + soft + adult = 1.0)
type AdultResult struct {
	Normal float64 `json:"normal"`
	Soft   float64 `json:"soft"`
	Adult  float64 `json:"adult"`
}

// AdultImageDetectResult represents an Adult Image Detection result.
type AdultImageDetectResult struct {
	RID    string      `json:"rid"`
	Result AdultResult `json:"result"`
}

// String implements fmt.Stringer.
func (ar AdultImageDetectResult) String() string { return common.String(ar) }

// SaveAs saves ar to @filename.
//
// The file extension must be .json.
func (ar AdultImageDetectResult) SaveAs(filename string) error {
	return common.SaveAsJSON(ar, filename)
}

// AdultImageDetectInitializer is a lazy adult image detector.
type AdultImageDetectInitializer struct {
	AuthKey  string
	Image    *os.File
	ImageURL string
}

// AdultImageDetect determines the level of nudity or adult content in the given image.
//
// Image can be either the image file (JPG or PNG) or image URL.
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-adult-content for more details.
func AdultImageDetect() *AdultImageDetectInitializer {
	return &AdultImageDetectInitializer{
		AuthKey:  common.KeyPrefix,
		Image:    nil,
		ImageURL: "",
	}
}

// WithFile sets the file to request on @filepath.
func (ai *AdultImageDetectInitializer) WithFile(filepath string) *AdultImageDetectInitializer {
	ai.Image = OpenFile(filepath)
	return ai
}

// WithURL sets the URL to request to @url.
func (ai *AdultImageDetectInitializer) WithURL(url string) *AdultImageDetectInitializer {
	ai.ImageURL = url
	return ai
}

// AuthorizeWith sets the authorization key to @key.
func (ai *AdultImageDetectInitializer) AuthorizeWith(key string) *AdultImageDetectInitializer {
	ai.AuthKey = common.FormatKey(key)
	return ai
}

// Collect returns the adult image detection result.
func (ai *AdultImageDetectInitializer) Collect() (res AdultImageDetectResult, err error) {
	client := &http.Client{}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if ai.Image != nil {
		part, err := writer.CreateFormFile("image", filepath.Base(ai.Image.Name()))
		if err != nil {
			return res, err
		}
		io.Copy(part, ai.Image)
	}
	defer writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/adult/detect?image_url=%s", prefix, ai.ImageURL), body)
	if err != nil {
		return res, err
	}
	req.Close = true

	req.Header.Set(common.Authorization, ai.AuthKey)
	if ai.Image != nil {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	defer ai.Image.Close()

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

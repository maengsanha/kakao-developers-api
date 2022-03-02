package vision

import (
	"bytes"
	"encoding/json"
	"fmt"
	"internal/common"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

// AdultResult represents a document of a detected adult image result.
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
	Filename string
	ImageURL string
	withFile bool
}

// AdultImageDetect determines the level of nudity or adult content in the given image.
//
// Image can be either the image file (JPG or PNG) or image URL.
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-adult-content for more details.
func AdultImageDetect() *AdultImageDetectInitializer {
	return &AdultImageDetectInitializer{
		AuthKey:  common.KeyPrefix,
		Filename: "",
		ImageURL: "",
	}
}

// WithFile sets image path to @filename.
func (ai *AdultImageDetectInitializer) WithFile(filename string) *AdultImageDetectInitializer {
	switch format := strings.Split(filename, "."); format[len(format)-1] {
	case "jpg", "png":
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	ai.Filename = filename
	ai.withFile = true
	return ai
}

// WithURL sets url to @url.
func (ai *AdultImageDetectInitializer) WithURL(url string) *AdultImageDetectInitializer {
	ai.ImageURL = url
	ai.withFile = false
	return ai
}

// AuthorizeWith sets the authorization key to @key.
func (ai *AdultImageDetectInitializer) AuthorizeWith(key string) *AdultImageDetectInitializer {
	ai.AuthKey = common.FormatKey(key)
	return ai
}

// Collect returns the adult image detection result.
func (ai *AdultImageDetectInitializer) Collect() (res AdultImageDetectResult, err error) {
	var req *http.Request

	if ai.withFile {

		file, err := os.Open(ai.Filename)
		if err != nil {
			return res, err
		}

		if stat, _ := file.Stat(); 2*1024*1024 < stat.Size() {
			return res, common.ErrTooLargeFile
		}

		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("image", ai.Filename)
		if err != nil {
			return res, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return res, err
		}

		writer.Close()

		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/adult/detect", prefix), body)
		if err != nil {
			return res, err
		}

		req.Header.Add("Content-Type", writer.FormDataContentType())
	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/adult/detect?image_url=%s", prefix, ai.ImageURL), nil)
		if err != nil {
			return res, err
		}

	}
	if err != nil {
		return
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

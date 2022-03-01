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
	Filename string
	ImageURL string
	withFile bool
}

// MultiTagCreate creates a tag according to the given image.
//
// Image can be either image URL or image file (JPG or PNG).
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-multi-tag for more details.
func MultiTagCreate() *MultiTagCreateInitializer {
	return &MultiTagCreateInitializer{
		AuthKey:  common.KeyPrefix,
		Filename: "",
		ImageURL: "",
	}
}

// WithFile sets image path to @filename.
func (mi *MultiTagCreateInitializer) WithFile(filename string) *MultiTagCreateInitializer {
	switch format := strings.Split(filename, "."); format[len(format)-1] {
	case "jpg", "png":
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	mi.Filename = filename
	mi.withFile = true
	return mi
}

// WithURL sets url to @url.
func (mi *MultiTagCreateInitializer) WithURL(url string) *MultiTagCreateInitializer {
	mi.ImageURL = url
	mi.withFile = false
	return mi
}

// AuthorizeWith sets the authorization key to @key.
func (mi *MultiTagCreateInitializer) AuthorizeWith(key string) *MultiTagCreateInitializer {
	mi.AuthKey = common.FormatKey(key)
	return mi
}

// Collect returns the Multi-tag creation result.
func (mi *MultiTagCreateInitializer) Collect() (res MultiTagCreateResult, err error) {
	client := &http.Client{}
	var req *http.Request
	if mi.withFile {

		file, err := os.Open(mi.Filename)
		if err != nil {
			return res, err
		}
		if stat, _ := file.Stat(); 2*1024*1024 < stat.Size() {
			return res, common.ErrTooLargeFile
		}

		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)

		part, err := writer.CreateFormFile("image", mi.Filename)
		if err != nil {
			return res, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return res, err
		}

		writer.Close()

		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/multitag/generate", prefix), body)
		if err != nil {
			return res, err
		}
		req.Header.Add("Content-Type", writer.FormDataContentType())

	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/multitag/generate?image_url=%s", prefix, mi.ImageURL), nil)
		if err != nil {
			return res, err
		}
	}
	if err != nil {
		return res, err
	}

	req.Close = true

	req.Header.Add(common.Authorization, mi.AuthKey)

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

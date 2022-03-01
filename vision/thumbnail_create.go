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

// ThumbnailCreateInitializer is a lazy thumbnail creator.
type ThumbnailCreateInitializer struct {
	AuthKey  string
	Filename string
	ImageURL string
	Width    int
	Height   int
	withFile bool
}

// ThumbnailCreateResult represents a Thumbnail creation result.
type ThumbnailCreateResult struct {
	ThumbnailImageURL string `json:"thumbnail_image_url"`
}

// String implements fmt.Stringer.
func (tr ThumbnailCreateResult) String() string { return common.String(tr) }

// SaveAs saves tr to @filename.
//
// The file extension must be .json.
func (tr ThumbnailCreateResult) SaveAs(filename string) error { return common.SaveAsJSON(tr, filename) }

// ThumbnailCreate crops the representative area out of the given image and creates a thumbnail image.
//
// Image can be either image URL or image file (JPG or PNG).
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-thumbnail for more details.
func ThumbnailCreate() *ThumbnailCreateInitializer {
	return &ThumbnailCreateInitializer{
		AuthKey:  common.KeyPrefix,
		Filename: "",
		ImageURL: "",
		Width:    0,
		Height:   0,
	}
}

// WithURL sets url to @url.
func (ti *ThumbnailCreateInitializer) WithURL(url string) *ThumbnailCreateInitializer {

	ti.ImageURL = url
	ti.withFile = false
	return ti
}

// WithFile sets image path to @filename.
func (ti *ThumbnailCreateInitializer) WithFile(filename string) *ThumbnailCreateInitializer {
	switch format := strings.Split(filename, "."); format[len(format)-1] {
	case "jpg", "png":
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	ti.Filename = filename
	ti.withFile = true
	return ti
}

// AuthorizeWith sets the authorization key to @key.
func (ti *ThumbnailCreateInitializer) AuthorizeWith(key string) *ThumbnailCreateInitializer {
	ti.AuthKey = common.FormatKey(key)
	return ti
}

// WidthTo sets image width to @ratio.
func (ti *ThumbnailCreateInitializer) WidthTo(ratio int) *ThumbnailCreateInitializer {
	ti.Width = ratio
	return ti
}

// Height sets image height to @ratio.
func (ti *ThumbnailCreateInitializer) HeightTo(ratio int) *ThumbnailCreateInitializer {
	ti.Height = ratio
	return ti
}

// Collect returns the thumbnail creation result.
func (ti *ThumbnailCreateInitializer) Collect() (res ThumbnailCreateResult, err error) {
	var req *http.Request
	client := &http.Client{}

	if ti.withFile {

		file, err := os.Open(ti.Filename)
		if err != nil {
			return res, err
		}

		if stat, _ := file.Stat(); 2*1024*1024 < stat.Size() {
			return res, common.ErrTooLargeFile
		}

		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		writer.WriteField("width", fmt.Sprintf("%d", ti.Width))
		writer.WriteField("height", fmt.Sprintf("%d", ti.Height))
		part, err := writer.CreateFormFile("image", ti.Filename)

		if err != nil {
			return res, err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return res, err
		}

		writer.Close()
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/thumbnail/crop",
			prefix), body)
		if err != nil {
			return res, err
		}
		req.Header.Add("Content-Type", writer.FormDataContentType())

	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/thumbnail/crop?image_url=%s&width=%d&height=%d",
			prefix, ti.ImageURL, ti.Width, ti.Height), nil)
		if err != nil {
			return res, err
		}
	}

	req.Close = true
	req.Header.Add(common.Authorization, ti.AuthKey)

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

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

// ThumbnailCreateInitializer is a lazy thumbnail creator.
type ThumbnailCreateInitializer struct {
	AuthKey  string
	Image    *os.File
	ImageUrl string
	Width    int
	Height   int
}

// ThumbnailCreateResult represents a Thumbnail creation result.
type ThumbnailCreateResult struct {
	ThumbnailImageUrl string `json:"thumbnail_image_url"`
}

// String implements fmt.Stringer.
func (tr ThumbnailCreateResult) String() string { return common.String(tr) }

// SaveAs saves tr to @filename.
//
// The file extension must be .json.
func (tr ThumbnailCreateResult) SaveAs(filename string) error { return common.SaveAsJSON(tr, filename) }

// ThumbnailCreate crops the representative area out of the given image and creates a thumbnail image.
//
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#create-thumbnail for more details.
func ThumbnailCreate(source string) *ThumbnailCreateInitializer {
	url, file := CheckSourceType(source)
	return &ThumbnailCreateInitializer{
		AuthKey:  common.KeyPrefix,
		Image:    file,
		ImageUrl: url,
		Width:    0,
		Height:   0,
	}
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
	client := new(http.Client)
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if ti.Image != nil {
		part, err := writer.CreateFormFile("image", filepath.Base(ti.Image.Name()))

		if err != nil {
			return res, err
		}
		io.Copy(part, ti.Image)
	}
	defer writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/thumbnail/crop?image_url=%s&width=%d&height=%d",
		prefix, ti.ImageUrl, ti.Width, ti.Height), body)
	if err != nil {
		return res, err
	}
	req.Close = true
	req.Header.Set(common.Authorization, ti.AuthKey)
	if ti.Image != nil {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	defer ti.Image.Close()

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

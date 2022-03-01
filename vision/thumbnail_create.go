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
	ImageURL string
	Width    int
	Height   int
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
		Image:    nil,
		ImageURL: "",
		Width:    0,
		Height:   0,
	}
}

// WithURL sets the URL to request to @url.
func (ti *ThumbnailCreateInitializer) WithURL(url string) *ThumbnailCreateInitializer {
	ti.ImageURL = url
	return ti
}

// WithFile sets the file to request on @filepath.
func (ti *ThumbnailCreateInitializer) WithFile(filepath string) *ThumbnailCreateInitializer {
	ti.Image = OpenFile(filepath)
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
	client := &http.Client{}
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	if ti.Image != nil {
		writer.WriteField("width", fmt.Sprintf("%d", ti.Width))
		writer.WriteField("height", fmt.Sprintf("%d", ti.Height))
		part, err := writer.CreateFormFile("image", filepath.Base(ti.Image.Name()))

		if err != nil {
			return res, err
		}
		io.Copy(part, ti.Image)
	}
	defer writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/thumbnail/crop?image_url=%s&width=%d&height=%d",
		prefix, ti.ImageURL, ti.Width, ti.Height), body)
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

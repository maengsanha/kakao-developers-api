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

// Thumbnail represents coordinates of the point starting the thumbnail image and its width, height.
type Thumbnail struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}

// ThumbnailResult represents a document of a detected thumbnail result.
type ThumbnailResult struct {
	Width     int       `json:"width"`
	Height    int       `json:"height"`
	Thumbnail Thumbnail `json:"thumbnail"`
}

// ThumbnailDetectResult represents a Thumbnail Detection result.
type ThumbnailDetectResult struct {
	RID    string          `json:"rid"`
	Result ThumbnailResult `json:"result"`
}

// String implements fmt.Stringer.
func (tr ThumbnailDetectResult) String() string { return common.String(tr) }

// SaveAs saves tr to @filename.
//
// The file extension must be .json.
func (tr ThumbnailDetectResult) SaveAs(filename string) error {
	return common.SaveAsJSON(tr, filename)
}

// ThumbnailDetectIniailizer is a lazy thumbnail detector.
type ThumbnailDetectInitializer struct {
	AuthKey  string
	Filename string
	ImageURL string
	Width    int
	Height   int
	withFile bool
}

// ThumbnailDetect helps to create a thumbnail image by detecting the representative area out of the given image.
//
// Image can be either image URL or image file (JPG or PNG).
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#extract-thumbnail for more details.
func ThumbnailDetect() *ThumbnailDetectInitializer {
	return &ThumbnailDetectInitializer{
		AuthKey:  common.KeyPrefix,
		ImageURL: "",
		Filename: "",
		Width:    0,
		Height:   0,
	}
}

// WithFile sets image path to @filename.
func (ti *ThumbnailDetectInitializer) WithFile(filename string) *ThumbnailDetectInitializer {
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

// WithURL sets url to @url.
func (ti *ThumbnailDetectInitializer) WithURL(url string) *ThumbnailDetectInitializer {
	ti.ImageURL = url
	ti.withFile = false
	return ti
}

// AuthorizeWith sets the authorization key to @key
func (ti *ThumbnailDetectInitializer) AuthorizeWith(key string) *ThumbnailDetectInitializer {
	ti.AuthKey = common.FormatKey(key)
	return ti
}

// WidthTo sets Image width to @ratio.
func (ti *ThumbnailDetectInitializer) WidthTo(ratio int) *ThumbnailDetectInitializer {
	ti.Width = ratio
	return ti
}

// HeightTo sets Image Height to @ratio.
func (ti *ThumbnailDetectInitializer) HeightTo(ratio int) *ThumbnailDetectInitializer {
	ti.Height = ratio
	return ti
}

// Collect returns the thumbnail detection result.
func (ti *ThumbnailDetectInitializer) Collect() (res ThumbnailDetectResult, err error) {
	var req *http.Request

	if ti.withFile {

		file, err := os.Open(ti.Filename)
		if err != nil {
			return res, err
		}

		if stat, _ := file.Stat(); 2*1024*1024 < stat.Size() {
			return res, common.ErrTooLargeFile
		}

		defer file.Close()

		body := new(bytes.Buffer)
		writer := multipart.NewWriter(body)
		writer.WriteField("width", fmt.Sprintf("%d", ti.Width))
		writer.WriteField("height", fmt.Sprintf("%d", ti.Height))
		part, err := writer.CreateFormFile("image", ti.Filename)
		if err != nil {
			return res, err
		}
		io.Copy(part, file)

		writer.Close()
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/thumbnail/detect",
			prefix), body)
		if err != nil {
			return res, err
		}
		req.Header.Add("Content-Type", writer.FormDataContentType())

	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/thumbnail/detect?image_url=%s&width=%d&height=%d",
			prefix, ti.ImageURL, ti.Width, ti.Height), nil)
		if err != nil {
			return res, err
		}
	}
	if err != nil {
		return
	}

	req.Close = true
	req.Header.Add(common.Authorization, ti.AuthKey)
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

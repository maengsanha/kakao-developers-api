package vision

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"internal/common"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

// Product represents coordinates of the detected product area box.
type Product struct {
	XMin  float64 `json:"x1"`
	YMax  float64 `json:"y1"`
	XMax  float64 `json:"x2"`
	YMin  float64 `json:"y2"`
	Class string  `json:"class"`
}

// ProductResult represents a document of a detected product result.
type ProductResult struct {
	Width   int       `json:"width"`
	Height  int       `json:"height"`
	Objects []Product `json:"objects"`
}

// ProductDetectResult represents a Product Detection result.
type ProductDetectResult struct {
	RID    string        `json:"rid"`
	Result ProductResult `json:"result"`
}

// String implements fmt.Stringer.
func (pr ProductDetectResult) String() string { return common.String(pr) }

// SaveAs saves pr to @filename.
//
// The file extension must be .json.
func (pr ProductDetectResult) SaveAs(filename string) error { return common.SaveAsJSON(pr, filename) }

// ProductDetectInitializer is a lazy product detector.
type ProductDetectInitializer struct {
	AuthKey   string
	Image     *os.File
	ImageURL  string
	Threshold float64
}

// ProductDetect detects the position and type of products within the given image.
//
// Image can be either image URL or image file (JPG or PNG).
// Refer to https://developers.kakao.com/docs/latest/en/vision/dev-guide#recog-product for more details.
func ProductDetect() *ProductDetectInitializer {
	return &ProductDetectInitializer{
		AuthKey:   common.KeyPrefix,
		Image:     nil,
		ImageURL:  "",
		Threshold: 0.8,
	}
}

// WithFile sets the file to request on @filepath.
func (pi *ProductDetectInitializer) WithFile(filepath string) *ProductDetectInitializer {
	pi.Image = OpenFile(filepath)
	return pi
}

// WithURL sets the URL to request to @url.
func (pi *ProductDetectInitializer) WithURL(url string) *ProductDetectInitializer {
	pi.ImageURL = url
	return pi
}

// AuthorizeWith sets the authorization key to @key.
func (pi *ProductDetectInitializer) AuthorizeWith(key string) *ProductDetectInitializer {
	pi.AuthKey = common.FormatKey(key)
	return pi
}

// ThresholdAt sets the Threshold to @val. (a value between 0 and 1.0)
//
// Threshold is a reference value to detect as a product.
func (pi *ProductDetectInitializer) ThresholdAt(val float64) *ProductDetectInitializer {
	if 0.1 <= val && val <= 1.0 {
		pi.Threshold = val
	} else {
		panic(errors.New("threshold must be between 0.1 and 1.0"))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return pi
}

// Collect returns the product detection result.
func (pi *ProductDetectInitializer) Collect() (res ProductDetectResult, err error) {
	client := new(http.Client)
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	if pi.Image != nil {
		writer.WriteField("threshold", fmt.Sprintf("%f", pi.Threshold))
		part, err := writer.CreateFormFile("image", filepath.Base(pi.Image.Name()))
		if err != nil {
			return res, err
		}
		io.Copy(part, pi.Image)
	}
	defer writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/product/detect?threshold=%f&image_url=%s", prefix, pi.Threshold, pi.ImageURL), body)
	if err != nil {
		return res, err
	}
	req.Close = true

	req.Header.Set(common.Authorization, pi.AuthKey)
	if pi.Image != nil {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	defer pi.Image.Close()

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

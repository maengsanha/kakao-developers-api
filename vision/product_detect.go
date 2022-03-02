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
	"strings"
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
	Filename  string
	ImageURL  string
	Threshold float64
	withFile  bool
}

// ProductDetect detects the position and type of products within the given image.
//
// Image can be either image URL or image file (JPG or PNG).
// Refer to https://developers.kakao.com/docs/latest/en/vision/dev-guide#recog-product for more details.
func ProductDetect() *ProductDetectInitializer {
	return &ProductDetectInitializer{
		AuthKey:   common.KeyPrefix,
		Filename:  "",
		ImageURL:  "",
		Threshold: 0.8,
	}
}

// WithFile sets image path to @filename.
func (pi *ProductDetectInitializer) WithFile(filename string) *ProductDetectInitializer {
	switch format := strings.Split(filename, "."); format[len(format)-1] {
	case "jpg", "png":
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	pi.Filename = filename
	pi.withFile = true
	return pi
}

// WithURL sets url to @url.
func (pi *ProductDetectInitializer) WithURL(url string) *ProductDetectInitializer {
	pi.ImageURL = url
	pi.withFile = false
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
	var req *http.Request

	if pi.withFile {
		file, err := os.Open(pi.Filename)
		if err != nil {
			return res, err
		}
		if stat, _ := file.Stat(); 2*1024*1024 < stat.Size() {
			return res, common.ErrTooLargeFile
		}
		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		writer.WriteField("threshold", fmt.Sprintf("%f", pi.Threshold))
		part, err := writer.CreateFormFile("image", pi.Filename)
		if err != nil {
			return res, err
		}
		_, err = io.Copy(part, file)
		if err != nil {
			return res, err
		}
		writer.Close()

		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/product/detect", prefix), body)
		if err != nil {
			return res, err
		}
		req.Header.Add("Content-Type", writer.FormDataContentType())
	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/product/detect?threshold=%f&image_url=%s", prefix, pi.Threshold, pi.ImageURL), nil)
		if err != nil {
			return
		}
	}
	if err != nil {
		return
	}
	req.Close = true
	req.Header.Add(common.Authorization, pi.AuthKey)
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

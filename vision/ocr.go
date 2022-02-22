package vision

import (
	"bytes"
	"encoding/json"
	"errors"
	"internal/common"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// OCRInitializer is a lazy Optical Character Recognition.
type OCRInitializer struct {
	AuthKey string
	Image   *os.File
}

// Result represents a document of a Optical Character Recognition result.
type Result struct {
	Boxes            [][]float64 `json:"boxes"`
	RecognitionWords []string    `json:"recognition_words"`
}

// OCRResult represents an Optical Character Recognition result.
type OCRResult struct {
	Result Result `json:"result"`
}

// SaveAs saves or to @filename.
//
// The file extension must be .json.
func (or OCRResult) SaveAs(filename string) error { return common.SaveAsJSON(or, filename) }

// String implements fmt.Stringer.
func (or OCRResult) String() string { return common.String(or) }

// OCR detects and extracts text from the given @filepath.
//
// File format must be one of the BMP, DIB, JPEG, JPE, JP2, WEBP, PBM, PGM, PPM, SR, RAS, TIFF, TIF, PNG and JPG.
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#ocr for more details.
func OCR(filepath string) *OCRInitializer {
	switch format := strings.Split(filepath, "."); format[len(format)-1] {
	case "bmp", "dib", "jpeg", "jpg", "jpe", "jp2", "png", "webp",
		"pbm", "pgm", "ppm", "sr", "ras", "tiff", "tif":
	default:
		panic(errors.New("unsupport format"))
	}

	file, err := os.Open(filepath)

	if err != nil {
		panic(errors.New("invalid file path"))
	}
	return &OCRInitializer{
		AuthKey: common.KeyPrefix,
		Image:   file,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (oi *OCRInitializer) AuthorizeWith(key string) *OCRInitializer {
	oi.AuthKey = common.FormatKey(key)
	return oi
}

// Collect returns the OCR result.
func (oi *OCRInitializer) Collect() (res json.RawMessage, err error) {
	client := new(http.Client)
	body := new(bytes.Buffer)
	bodywriter := multipart.NewWriter(body)
	filewriter, err := bodywriter.CreateFormFile("image", filepath.Base(oi.Image.Name()))

	if err != nil {
		return res, err
	}
	io.Copy(filewriter, oi.Image)
	defer bodywriter.Close()

	req, err := http.NewRequest(http.MethodPost, "https://dapi.kakao.com/v2/vision/text/ocr", body)
	if err != nil {
		return res, err
	}

	req.Header.Set(common.Authorization, oi.AuthKey)
	req.Header.Set("Content-Type", "multipart/form-data")

	defer req.Body.Close()
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}

	defer resp.Body.Close()
	defer oi.Image.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, err
	}

	return
}

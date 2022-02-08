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
	Boxes            [][]int  `json:"boxes"`
	RecognitionWords []string `json:"recognition_words"`
}

// OCRResult represents a Optical Character Recognition result.
type OCRResult struct {
	Result Result `json:"result"`
}

// SaveAs saves or to @filename
//
// The file extension must be .json.
func (or OCRResult) SaveAs(filename string) error { return common.SaveAsJSON(or, filename) }

// String implements fmt.Stringer.
func (or OCRResult) String() string { return common.String(or) }

// OCR detects and extracts text from the given @filepath.
//
// @filepath must be image file path.
// file format must be one of the BMP, DIB, JPEG, JPE, JP2, WEBP, PBM, PGM, PPM, SR, RAS, TIFF, TIF, PNG and JPG.
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#ocr for more details.
func OCR(filepath string) *OCRInitializer {
	switch format := strings.Split(filepath, "."); format[len(format)-1] {
	case "bmp", "dib", "jpeg", "jpg", "jpe", "jp2", "png", "webp",
		"pbm", "pgm", "ppm", "sr", "ras", "tiff", "tif":
		break
	default:
		panic(ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}

	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
	}
	CheckFileSize(file)
	CheckImagePixel(filepath)
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
func (oi *OCRInitializer) Collect() (res OCRResult, err error) {
	client := new(http.Client)

	body := &bytes.Buffer{}
	bodywriter := multipart.NewWriter(body)

	filewriter, err := bodywriter.CreateFormFile("image", filepath.Base(oi.Image.Name()))
	if err != nil {
		return res, err
	}
	io.Copy(filewriter, oi.Image)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/text/ocr", prefix), body)
	if err != nil {
		return res, err
	}
	req.Close = true
	req.Header.Set(common.Authorization, oi.AuthKey)
	req.Header.Set("Content-Type", "multipart/form-data")

	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()
	defer bodywriter.Close()
	defer oi.Image.Close()

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, err
	}
	return
}

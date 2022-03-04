// Copyright 2022 Sanha Maeng, Soyang Baek, Jinmyeong Kim
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vision

import (
	"bytes"
	"fmt"
	"internal/common"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/goccy/go-json"
)

// OCRInitializer is a lazy Optical Character Recognition.
type OCRInitializer struct {
	AuthKey  string
	Filename string
}

// Result represents a document of a Optical Character Recognition result.
type Result struct {
	Boxes            [][]int  `json:"boxes"`
	RecognitionWords []string `json:"recognition_words"`
}

// OCRResult represents an Optical Character Recognition result.
type OCRResult struct {
	Result []Result `json:"result"`
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
func OCR(filename string) *OCRInitializer {
	switch format := strings.Split(filename, "."); format[len(format)-1] {
	case "jpg", "png", "bmp", "jpeg", "jpe", "jp2", "webp", "pbm", "pgm", "ppm", "sr",
		"ras", "tiff", "tif", "dib":
	default:
		panic(common.ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return &OCRInitializer{
		AuthKey:  common.KeyPrefix,
		Filename: filename,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (oi *OCRInitializer) AuthorizeWith(key string) *OCRInitializer {
	oi.AuthKey = common.FormatKey(key)
	return oi
}

// Collect returns the OCR result.
func (oi *OCRInitializer) Collect() (res OCRResult, err error) {
	file, err := os.Open(oi.Filename)
	if err != nil {
		return res, err
	}

	defer file.Close()

	if stat, err := file.Stat(); err != nil {
		return res, err
	} else if 2*1024*1024 < stat.Size() {
		return res, common.ErrTooLargeFile
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", oi.Filename)
	if err != nil {
		return
	}

	_, err = io.Copy(part, file)
	if err != nil {
		return
	}

	writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/text/ocr", prefix), body)
	if err != nil {
		return
	}

	req.Header.Add(common.Authorization, oi.AuthKey)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}

	return
}

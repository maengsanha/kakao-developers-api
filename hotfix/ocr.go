package hotfix

import (
	"bytes"
	"encoding/json"
	"internal/common"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

const API_URL = "https://dapi.kakao.com/v2/vision/text/ocr"

// OCR result type
// CAUTION: type of ``Result`` is []Result, not Result
type OCRResult struct {
	Result []Result `json:"result"`
}

// Result field of OCR result
type Result struct {
	Boxes            [][]int  `json:"boxes"`
	RecognitionWords []string `json:"recognition_words"`
}

// String implements fmt.Stringer
func (o OCRResult) String() string { return common.String(o) }

// convert OCR Python code to Go.
// FIXME: all errors are ignored
// FIXME: no image resizing
func OCR(filename, key string) (res OCRResult) {
	// open file using filename
	file, _ := os.Open(filename)
	defer file.Close()

	// create a new form-data with header to body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("image", filename)
	_, _ = io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest(http.MethodPost, API_URL, body)
	req.Header.Add(common.Authorization, common.FormatKey(key))
	req.Header.Add("Content-Type", writer.FormDataContentType())

	// CAUTION: do not use new(), use &Type{}, as this achieves zero-allocation
	client := &http.Client{}
	resp, _ := client.Do(req)

	_ = json.NewDecoder(resp.Body).Decode(&res)

	return
}

package pose

import (
	"bytes"
	"encoding/json"
	"fmt"
	"internal/common"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

// AnalyzeVideoResult returns the result code of analyze video.
type AnalyzeVideoResult struct {
	JobId string `json:"job_id"`
}

// AnalyzeVideoIterator is a lazy video analyzer.
type AnalyzeVideoInitializer struct {
	AuthKey     string
	VideoURL    string
	Filename    string
	Smoothing   bool
	CallbackURL string
	withFile    bool
}

// String implements fmt.Stringer.
func (ar AnalyzeVideoResult) String() string { return common.String(ar) }

// SaveAs saves ar to @filename.
//
// The file extension could be .json.
func (ar AnalyzeVideoResult) SaveAs(filename string) error {
	return common.SaveAsJSONorXML(ar, filename)
}

// AnalyzeVideo detects people in each frame of the requested video and extracts key points.
//
// For more details visit https://developers.kakao.com/docs/latest/en/pose/dev-guide#job-submit.
func AnalyzeVideo() *AnalyzeVideoInitializer {
	return &AnalyzeVideoInitializer{
		AuthKey:     common.KeyPrefix,
		Smoothing:   true,
		CallbackURL: "",
	}
}

// WithURL sets url to @url.
func (ai *AnalyzeVideoInitializer) WithURL(url string) *AnalyzeVideoInitializer {
	ai.VideoURL = url
	ai.withFile = false
	return ai
}

// WithFile sets filepath to @filename.
func (ai *AnalyzeVideoInitializer) WithFile(filename string) *AnalyzeVideoInitializer {
	ai.Filename = filename
	ai.withFile = true
	return ai
}

// AuthorizeWith sets the authorization key to @key.
func (ai *AnalyzeVideoInitializer) AuthorizeWith(key string) *AnalyzeVideoInitializer {
	ai.AuthKey = common.FormatKey(key)
	return ai
}

// SetSmoothing sets smoothing that apply the smoothing process to the position of the key points between the detected frames.
func (ai *AnalyzeVideoInitializer) SetSmoothing(set bool) *AnalyzeVideoInitializer {
	ai.Smoothing = set
	return ai
}

// ReceiveTo sets a callback URL to receive a callback when the video analysis is completed.
func (ai *AnalyzeVideoInitializer) ReceiveTo(url string) *AnalyzeVideoInitializer {
	ai.CallbackURL = url
	return ai
}

// Collect returns the result of AnalyzeVideo.
func (ai *AnalyzeVideoInitializer) Collect() (res AnalyzeVideoResult, err error) {
	var req *http.Request
	if ai.withFile {
		file, err := os.Open(ai.Filename)
		if err != nil {
			return res, err
		}

		if stat, _ := file.Stat(); 50*1024*1024 < stat.Size() {
			return res, common.ErrTooLargeFile
		}

		defer file.Close()

		body := &bytes.Buffer{}
		writer := multipart.NewWriter(body)
		part, err := writer.CreateFormFile("file", ai.Filename)
		if err != nil {
			return res, err
		}

		_, err = io.Copy(part, file)
		if err != nil {
			return res, err
		}

		writer.Close()

		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/job", prefix), body)
		if err != nil {
			return res, err
		}

		req.Header.Add("Content-Type", writer.FormDataContentType())
	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/job?video_url=%s", prefix, ai.VideoURL), nil)
		if err != nil {
			return
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	req.Close = true

	req.Header.Add(common.Authorization, ai.AuthKey)

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

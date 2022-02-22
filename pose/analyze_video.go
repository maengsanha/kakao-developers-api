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
	"path/filepath"
)

// AnalyzeVideoResult returns the result code of analyze video.
type AnalyzeVideoResult struct {
	JobId string `json:"job_id"`
}

// AnalyzeVideoIterator is a lazy video analyzer.
type AnalyzeVideoIterator struct {
	AuthKey     string
	VideoURL    string
	File        *os.File
	Smoothing   bool
	CallbackURL string
}

// AnalyzeVideo detects people in each frame of the requested video and extracts key points.
//
// For more details visit https://developers.kakao.com/docs/latest/en/pose/dev-guide#job-submit.
func AnalyzeVideo() *AnalyzeVideoIterator {
	return &AnalyzeVideoIterator{
		AuthKey:   common.KeyPrefix,
		Smoothing: true,
	}
}

// WithURL sets @url to VideoURL.
func (ai *AnalyzeVideoIterator) WithURL(url string) *AnalyzeVideoIterator {
	return &AnalyzeVideoIterator{
		AuthKey:   common.KeyPrefix,
		VideoURL:  url,
		Smoothing: true,
	}
}

// WithFile sets @filepath to File.
func (ai *AnalyzeVideoIterator) WithFile(filepath string) *AnalyzeVideoIterator {
	bs, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	if stat, _ := bs.Stat(); stat.Size() > 50*1024*1024 {
		panic("up to 50MB are allowed")
	} else {
		return &AnalyzeVideoIterator{
			AuthKey: common.KeyPrefix,
			File:    bs,
		}
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ai *AnalyzeVideoIterator) AuthorizeWith(key string) *AnalyzeVideoIterator {
	ai.AuthKey = common.FormatKey(key)
	return ai
}

// SetSmoothing sets smoothing that apply the smoothing process to the position of the key points between the detected frames.
func (ai *AnalyzeVideoIterator) SetSmoothing(set bool) *AnalyzeVideoIterator {
	ai.Smoothing = set
	return ai
}

// Callback sets a callback URL to receive a callback when the video analysis is completed.
func (ai *AnalyzeVideoIterator) Callback(url string) *AnalyzeVideoIterator {
	ai.CallbackURL = url
	return ai
}

// Collect returns the result of AnalyzeVideo.
func (ai *AnalyzeVideoIterator) Collect() (res AnalyzeVideoResult, err error) {
	client := new(http.Client)
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	if err != nil {
		return
	}

	if ai.File != nil {
		part, err := writer.CreateFormFile("file", filepath.Base(ai.File.Name()))
		if err != nil {
			return res, err
		}
		io.Copy(part, ai.File)
	}
	defer writer.Close()

	var req *http.Request
	if ai.File != nil {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/job?file=%s", prefix, ai.File.Name()), body)
	} else {
		req, err = http.NewRequest(http.MethodPost, fmt.Sprintf("%s/job?video_url=%s", prefix, ai.VideoURL), nil)
	}

	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(common.Authorization, ai.AuthKey)
	if ai.File != nil {
		req.Header.Set("Content-Type", "multipart/form-data")
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

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

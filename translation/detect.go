package translation

import (
	"encoding/json"
	"errors"
	"fmt"
	"internal/common"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// LanguageInfo represents list of detected languages sorted by highest confidence.
//
// Up to three results are returned.
type LanguageInfo struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

// DetectLanguageResult represents a detected language result.
type DetectLanguageResult struct {
	LanguageInfo []LanguageInfo `json:"language_info"`
}

// String implements fmt.Stringer.
func (dl DetectLanguageResult) String() string { return common.String(dl) }

// SaveAs saves dl to @filename.
func (dl *DetectLanguageResult) SaveAs(filename string) error {
	return common.SaveAsJSON(dl, filename)
}

// DetectLanguageInitializer is a lazy language detector.
type DetectLanguageInitializer struct {
	Query   string
	Authkey string
}

// DetectLanguage detects the language of the given @query.
//
// See https://developers.kakao.com/docs/latest/ko/translate/dev-guide#language-detect for more details.
func DetectLanguage(query string) *DetectLanguageInitializer {
	if 5000 < len(query) {
		panic(errors.New("query must be 5,000 bytes or less"))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}

	return &DetectLanguageInitializer{
		Query:   url.QueryEscape(strings.TrimSpace(query)),
		Authkey: common.KeyPrefix,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (dl *DetectLanguageInitializer) AuthorizeWith(key string) *DetectLanguageInitializer {
	dl.Authkey = common.FormatKey(key)
	return dl
}

// RequestBy returns the detected language result by sending a HTTP @method (GET or POST).
func (dl *DetectLanguageInitializer) RequestBy(method string) (res DetectLanguageResult, err error) {
	client := new(http.Client)
	switch method {
	case "GET", "POST":
		req, err := http.NewRequest(method,
			fmt.Sprintf("%s/v3/translation/language/detect?query=%s", prefix, dl.Query), nil)
		if err != nil {
			return res, err
		}

		req.Close = true

		req.Header.Set(common.Authorization, dl.Authkey)
		if method == "POST" {
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		}
		resp, err := client.Do(req)
		if err != nil {
			return res, err
		}

		defer resp.Body.Close()

		if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
			return res, err
		}
	default:
		return res, errors.New("method must be either GET or POST")
	}
	return
}

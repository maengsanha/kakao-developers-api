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

// LanguageInfo represents a detected language.
type LanguageInfo struct {
	Code       string  `json:"code"`
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

// DetectResult represents a language detection result.
type DetectResult struct {
	LanguageInfo []LanguageInfo `json:"language_info"`
}

// String implements fmt.Stringer.
func (dr DetectResult) String() string { return common.String(dr) }

// SaveAs saves dr to @filename.
//
// The file extension must be .json.
func (dr DetectResult) SaveAs(filename string) error { return common.SaveAsJSON(dr, filename) }

// DetectInitializer is a lazy language detector.
type DetectInitializer struct {
	Query   string
	Authkey string
}

// Detect detects the language of the given @text.
//
// See https://developers.kakao.com/docs/latest/ko/translate/dev-guide#language-detect for more details.
func Detect(text string) *DetectInitializer {
	if 5000 < len(text) {
		panic(errors.New("up to 5,000 characters are allowed"))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}

	return &DetectInitializer{
		Query:   url.QueryEscape(strings.TrimSpace(text)),
		Authkey: common.KeyPrefix,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (di *DetectInitializer) AuthorizeWith(key string) *DetectInitializer {
	di.Authkey = common.FormatKey(key)
	return di
}

// Collect returns the language detection result.
func (di *DetectInitializer) Collect() (res DetectResult, err error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/v3/translation/language/detect?query=%s", prefix, di.Query), nil)
	if err != nil {
		return res, err
	}

	req.Close = true

	req.Header.Set(common.Authorization, di.Authkey)

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

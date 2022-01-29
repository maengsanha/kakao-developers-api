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

// TranslateResult represents a result of translate.
type TranslateResult struct {
	TranslatedText [][]string `json:"translated_text"`
}

// String implements fmt.Stringer.
func (tr TranslateResult) String() string { return common.String(tr) }

// SaveAs saves tr to @filename.
//
// The file extension must be .json.
func (tr TranslateResult) SaveAs(filename string) error { return common.SaveAsJSON(tr, filename) }

// TranslateInitializer is a lazy translator.
type TranslateInitializer struct {
	Query      string
	SrcLang    string
	TargetLang string
	AuthKey    string
}

// Translate translates the input text into various languages.
//
// For more details visit https://developers.kakao.com/docs/latest/en/translate/dev-guide#trans-sentence.
func Translate(query string) *TranslateInitializer {
	if 5000 < len(query) {
		panic(errors.New("query must be 5,000 bytes or less"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return &TranslateInitializer{
		Query:      url.QueryEscape(strings.TrimSpace(query)),
		SrcLang:    "",
		TargetLang: "",
		AuthKey:    common.KeyPrefix,
	}
}

// AuthorizeWith sets the authorization key to @key.
func (ti *TranslateInitializer) AuthorizeWith(key string) *TranslateInitializer {
	ti.AuthKey = common.FormatKey(key)
	return ti
}

// Source sets the source language that input text to be translated.
//
// There are following source language exist:
//
// kr: Korean
//
// en: English
//
// jp: Japanese
//
// cn: Chinese
//
// vi: Vietnamese
//
// id: Indonesian
//
// ar: Arabic
//
// bn: Bangal language
//
// de: German
//
// es: Spanish
//
// fr: French
//
// hi: Hindustani
//
// it: Italian
//
// ms: Malaysian
//
// nl: Dutch
//
// pt: Portuguese
//
// ru: Russian
//
// th: Thai
//
// tr: Turkish
func (ti *TranslateInitializer) Source(src string) *TranslateInitializer {
	switch src {
	case "kr", "en", "jp", "cn", "vi", "id", "ar", "bn", "de", "es", "fr", "hi", "it", "ms", "nl", "pt", "ru", "th", "tr":
		ti.SrcLang = src
	default:
		panic(errors.New("source language must be one of the following options:\nkr, en, jp, cn, vi, id, ar, bn, de, es, fr, hi, it, ms, nl, pt, ru, th, tr"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ti
}

// Target sets the target langage that input text is translated into.
//
// There are following target language exist:
//
// kr: Korean
//
// en: English
//
// jp: Japanese
//
// cn: Chinese
//
// vi: Vietnamese
//
// id: Indonesian
//
// ar: Arabic
//
// bn: Bangal language
//
// de: German
//
// es: Spanish
//
// fr: French
//
// hi: Hindustani
//
// it: Italian
//
// ms: Malaysian
//
// nl: Dutch
//
// pt: Portuguese
//
// ru: Russian
//
// th: Thai
//
// tr: Turkish
func (ti *TranslateInitializer) Target(target string) *TranslateInitializer {
	switch target {
	case "kr", "en", "jp", "cn", "vi", "id", "ar", "bn", "de", "es", "fr", "hi", "it", "ms", "nl", "pt", "ru", "th", "tr":
		ti.TargetLang = target
	default:
		panic(errors.New("target language must be one of the following options:\nkr, en, jp, cn, vi, id, ar, bn, de, es, fr, hi, it, ms, nl, pt, ru, th, tr"))
	}
	if r := recover(); r != nil {
		log.Println(r)
	}
	return ti
}

// CollectByGET returns the translate result.
func (ti *TranslateInitializer) CollectByGET() (res TranslateResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/v2/translation/translate?src_lang=%s&target_lang=%s&query=%s",
			prefix, ti.SrcLang, ti.TargetLang, ti.Query), nil)
	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(common.Authorization, ti.AuthKey)

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

// CollectByPOST returns the translate result.
func (ti *TranslateInitializer) CollectByPOST() (res TranslateResult, err error) {
	client := new(http.Client)
	req, err := http.NewRequest(http.MethodPost,
		fmt.Sprintf("%s/v2/translation/translate?src_lang=%s&target_lang=%s&query=%s",
			prefix, ti.SrcLang, ti.TargetLang, ti.Query), nil)
	if err != nil {
		return
	}

	req.Close = true

	req.Header.Set(common.Authorization, ti.AuthKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return
	}

	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return
	}
	return
}

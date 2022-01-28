package common

import "strings"

const (
	Authorization = "Authorization"
	KeyPrefix     = "KakaoAK "
	REST_API_KEY  = "" // paste your REST API key here
)

// FormatKey formats @key to Kakao Developers' authorization key format.
func FormatKey(key string) string { return KeyPrefix + strings.TrimSpace(key) }

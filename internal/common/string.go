package common

import (
	"bytes"
	"encoding/json"
)

// String implements fmt.Stringer.
func String(data interface{}) string {
	// let's make two buffers
	// one for encoding, the other for indenting
	encBuf, indBuf := new(bytes.Buffer), new(bytes.Buffer)

	// okay, now encode without HTML escaping
	encoder := json.NewEncoder(encBuf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(data)

	// indent the encoded buffer
	json.Indent(indBuf, encBuf.Bytes(), "", "  ")
	return indBuf.String()
}

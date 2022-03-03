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

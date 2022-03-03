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
	"encoding/xml"
	"io/ioutil"
	"strings"
)

// SaveAsJSON saves @data to @filename.
//
// @filename should end with .json.
func SaveAsJSON(data interface{}, filename string) error {
	switch tokens := strings.Split(filename, "."); tokens[len(tokens)-1] {
	case "json":
		encBuf, indBuf := new(bytes.Buffer), new(bytes.Buffer)

		encoder := json.NewEncoder(encBuf)
		encoder.SetEscapeHTML(false)
		if err := encoder.Encode(data); err != nil {
			return err
		}

		if err := json.Indent(indBuf, encBuf.Bytes(), "", "  "); err != nil {
			return err
		}

		return ioutil.WriteFile(filename, indBuf.Bytes(), 0o644)
	default:
		return ErrUnsupportedFormat
	}
}

// SaveAsJSONorXML saves @data to @filename.
//
// @filename should end with .json or .xml.
func SaveAsJSONorXML(data interface{}, filename string) error {
	switch tokens := strings.Split(filename, "."); tokens[len(tokens)-1] {
	case "json":
		return SaveAsJSON(data, filename)
	case "xml":
		if bs, err := xml.MarshalIndent(data, "", "  "); err != nil {
			return err
		} else {
			return ioutil.WriteFile(filename, bs, 0o644)
		}
	default:
		return ErrUnsupportedFormat
	}
}

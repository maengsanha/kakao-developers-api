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

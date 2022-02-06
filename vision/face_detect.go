package vision

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"internal/common"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Face represents data of the detected face.
type Face struct {
	FacialAttributes FacialAttributes `json:"facial_attributes"`
	FacialPoints     FacialPoints     `json:"facial_points"`
	Score            float64          `json:"score"`
	ClassIdx         int              `json:"class_idx"`
	X                float64          `json:"x"`
	Y                float64          `json:"y"`
	W                float64          `json:"w"`
	H                float64          `json:"h"`
	Pitch            float64          `json:"pitch"`
	Yaw              float64          `json:"yaw"`
	Roll             float64          `json:"roll"`
}

// FacialAttributes represents estimated gender, age of the detected face.
type FacialAttributes struct {
	Gender Gender  `json:"gender"`
	Age    float64 `json:"age"`
}

// Gender represents confidence score that the detected face is considered as male or female.
type Gender struct {
	Male   float64 `json:"male"`
	Female float64 `json:"female"`
}

// FacialPoints represents array of coordinates of the detected face. (a value between 0 and 1.0)
type FacialPoints struct {
	Jaw          [][]float64 `json:"jaw"`
	RightEyebrow [][]float64 `json:"right_eyebrow"`
	LeftEyebrow  [][]float64 `json:"left_eyebrow"`
	Nose         [][]float64 `json:"nose"`
	RightEye     [][]float64 `json:"right_eye"`
	LeftEye      [][]float64 `json:"left_eye"`
	Lip          [][]float64 `json:"lip"`
}

// FaceResult represents a document of a detected face result.
type FaceResult struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Faces  []Face `json:"faces"`
}

// FaceDetectResult represents a Face Detect result.
type FaceDetectResult struct {
	Rid    string     `json:"rid"`
	Result FaceResult `json:"result"`
}

// String implements fmt.Stringer.
func (fr FaceDetectResult) String() string { return common.String(fr) }

// SaveAs saves fr to @filename
//
// The file extension must be .json.
func (fr FaceDetectResult) SaveAs(filename string) error { return common.SaveAsJSON(fr, filename) }

// FaceDetectInitializer is a lazy face detector.
type FaceDetectInitializer struct {
	AuthKey   string
	Image     *os.File
	ImageUrl  string
	Threshold float64
}

// FaceDetect detects a face in the given @source.
//
// @source can be requested with either the image or image_url, PNG and JPG format only.
// Refer to https://developers.kakao.com/docs/latest/ko/vision/dev-guide#recog-face for more details.
func FaceDetect(source string) *FaceDetectInitializer {
	switch format := strings.Split(source, "."); format[len(format)-1] {
	case "jpg", "png":
		break
	default:
		panic(ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	if source[0:4] == "http" {
		return &FaceDetectInitializer{
			AuthKey:   common.KeyPrefix,
			ImageUrl:  source,
			Image:     nil,
			Threshold: 0.7,
		}
	} else {
		bs, err := os.Open(source)
		if err != nil {
			panic(err)
		}
		if stat, _ := bs.Stat(); stat.Size() > 2*1024*1024 {
			panic(errors.New("file size must be 2 mb or less"))
		} else {
			return &FaceDetectInitializer{
				AuthKey:   common.KeyPrefix,
				ImageUrl:  "",
				Image:     bs,
				Threshold: 0.7,
			}
		}
	}
}

// AuthorizeWith sets the authorization key to @key.
func (fi *FaceDetectInitializer) AuthorizeWith(key string) *FaceDetectInitializer {
	fi.AuthKey = common.FormatKey(key)
	return fi
}

// ThresholdAt sets the Threshold to @val. (a value between 0 and 1.0)
//
// Threshold is reference value to detect as a face.
// *If this value is set too high, some faces may not be able to be detected as a face.
// *If this value is set too low, other area can be detected as a face.
func (fi *FaceDetectInitializer) ThresholdAt(val float64) *FaceDetectInitializer {
	if 0.1 <= val && val <= 1.0 {
		fi.Threshold = val
	} else {
		panic(errors.New("threshold must be between 0.1 and 1.0"))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
	return fi
}

// Collect returns the face detection result.
func (fi *FaceDetectInitializer) Collect() (res FaceDetectResult, err error) {
	client := new(http.Client)

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	if err != nil {
		return res, err
	}

	if fi.Image != nil {
		part, err := writer.CreateFormFile("image", filepath.Base(fi.Image.Name()))
		if err != nil {
			return res, err
		}
		io.Copy(part, fi.Image)
	}

	defer writer.Close()

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/face/detect?threshold=%f&image_url=%s", prefix, fi.Threshold, fi.ImageUrl), body)
	if err != nil {
		return res, err
	}
	req.Close = true

	req.Header.Set(common.Authorization, fi.AuthKey)
	if fi.Image != nil {
		req.Header.Set("Content-Type", writer.FormDataContentType())
	} else {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	defer fi.Image.Close()

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

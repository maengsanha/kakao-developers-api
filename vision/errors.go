package vision

import (
	"errors"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"strings"
)

const mb = 1024 * 1024

func OpenFile(filepath string) *os.File {

	// Open the file.
	file, err := os.Open(filepath)
	if err != nil {
		panic(errors.New("invalid file path"))
	}

	// Check the file size.
	if stat, _ := file.Stat(); stat.Size() > 2*mb {
		panic(errors.New("file size must be 2 mb or less"))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}

	// Check the file format.
	switch format := strings.Split(filepath, "."); format[len(format)-1] {
	case "jpg":
		// Open the file for Decode.
		imgfile, _ := os.Open(filepath)
		im, err := jpeg.Decode(imgfile)
		if err != nil {
			panic(errors.New("invalid image file"))
		}
		// Check the image pixel.
		if imgsize := im.Bounds().Max; imgsize.X > 2048 || imgsize.Y > 2048 {
			panic(errors.New("image pixel must be 2048px or less"))
		}
	case "png":
		// Open the file for Decode.
		imgfile, _ := os.Open(filepath)
		im, err := png.Decode(imgfile)
		if err != nil {
			panic(errors.New("invalid image file"))
		}
		// Check the image pixel.
		if imgsize := im.Bounds().Max; imgsize.X > 2048 || imgsize.Y > 2048 {
			panic(errors.New("image pixel must be 2048px or less"))
		}
	default:
		panic(errors.New("file format must be either jpg or png"))
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}

	return file
}

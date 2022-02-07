package vision

import (
	"errors"
	"log"
	"os"
	"strings"
)

var (
	ErrUnsupportedFormat = errors.New("file format must be either jpg or png")
	ErrOverTheFileSize   = errors.New("file size must be 2 mb or less")
	ErrOverTheImagePixel = errors.New("image pixel must be 2048px or less")
)

func CheckFileFormat(source string) {
	switch format := strings.Split(source, "."); format[len(format)-1] {
	case "jpg", "png":
		break
	default:
		panic(ErrUnsupportedFormat)
	}
	if r := recover(); r != nil {
		log.Panicln(r)
	}
}

func CheckFileSize(file os.File) {
	if stat, _ := file.Stat(); stat.Size() > 2*1024*1024 {
		panic(ErrOverTheFileSize)
	} else {
		return
	}
}

// func CheckImagePixel(file *os.File) {

// }

func CheckSourceType(source string) (string, *os.File) {
	CheckFileFormat(source)
	if source[0:4] == "http" {
		return source, nil
	} else {
		file, err := os.Open(source)
		if err != nil {
			log.Panicln(err)
		}
		CheckFileSize(*file)
		return "", file
	}
}

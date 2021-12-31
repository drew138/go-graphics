package encoding

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

var formats [3]string = [3]string{"png", "jpg", "jpeg"}

func isSupportedFormat(format string) error {
	for _, val := range formats {
		if val == format {
			return nil
		}
	}
	return fmt.Errorf("image format '%v' not supported", format)
}

// EncodeBytesArrayToFile takes an array of bytes
// to be decoded into an image and consequently encoded
func EncodeBytesArrayToFile(bytesArr []byte, w io.Writer) error {

	img, format, err := image.Decode(bytes.NewReader(bytesArr))
	if err != nil {
		return err
	}
	if err := isSupportedFormat(format); err != nil {
		return err
	}
	switch format {
	case "png":
		err := png.Encode(w, img)
		if err != nil {
			return err
		}
	case "jpg":
	case "jpeg":
		err := jpeg.Encode(w, img, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

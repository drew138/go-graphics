package encoding

import (
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

// EncodeBytesArrayToFile takes an image object
// to be encoded to the desired format
func EncodeImageToBytes(img image.Image, format string, w io.Writer) error {

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

package filters

import (
	"image"
	"image/color"

	"github.com/drew138/go-graphics/filters/kernels"
)

func roundRGBValue(val float32) float32 {
	if val < 0 {
		return 0
	} else if val > 255 {
		return 255
	}
	return val
}

func weighedSum(img image.Image, k kernels.Kernel, x int, y int) (uint8, uint8, uint8, uint8) {
	n, m := len(k), len(k[0])
	halfSide := int(n / 2)
	var r, g, b float32 = 0.0, 0.0, 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rAtPosition, gAtPosition, bAtPosition, _ := img.At(x+i-halfSide, y+j-halfSide).RGBA()
			rAtPosition >>= 8
			gAtPosition >>= 8
			bAtPosition >>= 8
			r += float32(rAtPosition) * k[i][j]
			g += float32(gAtPosition) * k[i][j]
			b += float32(bAtPosition) * k[i][j]
		}
	}
	r = roundRGBValue(r)
	g = roundRGBValue(g)
	b = roundRGBValue(b)
	_, _, _, a := img.At(x, y).RGBA()
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

// ApplyFilter Applies a convolutional matrix
// (kernel) to each pixel of the image provided
func ApplyFilter(i image.Image, k kernels.Kernel) *image.RGBA {
	bounds := i.Bounds()
	rgba := image.NewRGBA(bounds)
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 1; y+1 <= height; y++ {
		for x := 1; x+1 <= width; x++ {
			r, g, b, a := weighedSum(i, k, x, y)
			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			rgba.Set(x, y, c)
		}
	}
	for y := 0; y <= height; y++ {
		colorOne := rgba.At(1, y)
		rgba.Set(0, y, colorOne)
		colorTwo := rgba.At(width-2, y)
		rgba.Set(width-1, y, colorTwo)
	}
	for x := 0; x <= width; x++ {
		colorOne := rgba.At(x, 1)
		rgba.Set(x, 0, colorOne)
		colorTwo := rgba.At(x, height-2)
		rgba.Set(x, height-1, colorTwo)
	}
	return rgba
}

// CreateNegativeImage Creates negative of an image
// by changing each rgb value to its remainder to 255.
func CreateNegativeImage(i image.Image) *image.RGBA {
	bounds := i.Bounds()
	rgba := image.NewRGBA(bounds)
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			r, g, b, a := i.At(x, y).RGBA()
			c := color.RGBA{255 - uint8(r>>8), 255 - uint8(g>>8), 255 - uint8(b>>8), uint8(a)}
			rgba.Set(x, y, c)
		}
	}
	return rgba
}

package processing

import (
	"image"
	"image/color"

	kernels "github.com/drew138/go-graphics/processing/kernels"
)

func weighedSum(img image.Image, k *kernels.Kernel, x int, y int, maxY int, maxX int) (uint8, uint8, uint8, uint8) {
	if x == 0 || y == 0 || y == maxY || x == maxX {
		r, g, b, a := img.At(x, y).RGBA()
		return uint8(r), uint8(g), uint8(b), uint8(a)
	}
	n, m := len(k.Kernel), len(k.Kernel[0])
	l := int(n / 2)
	var r, g, b, _ float32 = 0.0, 0.0, 0.0, 0.0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rp, gp, bp, _ := img.At(x+j-l, y+i-l).RGBA()
			rp >>= 8
			gp >>= 8
			bp >>= 8
			r += float32(rp) * k.Kernel[i][j]
			g += float32(gp) * k.Kernel[i][j]
			b += float32(bp) * k.Kernel[i][j]
		}
	}
	if r < 0 {
		r = 0
	}
	if g < 0 {
		g = 0
	}
	if b < 0 {
		b = 0
	}
	if r > 255 {
		r = 255
	}
	if g > 255 {
		g = 255
	}
	if b > 255 {
		b = 255
	}
	_, _, _, a := img.At(x, y).RGBA()
	return uint8(r), uint8(g), uint8(b), uint8(a)
}

func TransformImage(i image.Image, k *kernels.Kernel) *image.RGBA {

	bounds := i.Bounds()
	rgba := image.NewRGBA(bounds)
	width, height := bounds.Max.X, bounds.Max.Y
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := weighedSum(i, k, x, y, height-1, width-1)
			c := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			rgba.Set(x, y, c)
		}
	}
	return rgba
}

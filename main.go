package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"time"

	filters "github.com/drew138/go-graphics/filters"
	kernels "github.com/drew138/go-graphics/filters/kernels"
)

func main() {
	// f, _ := os.Open("./ej.png")
	start := time.Now()
	f, _ := os.Open("./sarah.jpg")
	// f, _ := os.Open("./lenna.png")

	file := io.Reader(f)
	i, format, _ := image.Decode(file)
	img := filters.ApplyFilter(i, kernels.EdgeDetection)
	// img := processing.CreateNegativeImage(i)
	// _ = processing.TransformImage(i, &kernels.EdgeDetection)
	a, _ := os.Create("sarah." + format)
	if format == "png" {
		_ = png.Encode(a, img)
	} else {
		_ = jpeg.Encode(a, img, nil)
	}
	elapsed := time.Since(start)
	log.Printf("sharpen took %s", elapsed)
}

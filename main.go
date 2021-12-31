package main

import (
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"time"

	processing "github.com/drew138/go-graphics/processing"
	kernels "github.com/drew138/go-graphics/processing/kernels"
)

func main() {
	// f, _ := os.Open("./ej.png")
	start := time.Now()
	f, _ := os.Open("./sarah.jpg")
	// f, _ := os.Open("./lenna.png")

	file := io.Reader(f)
	i, format, _ := image.Decode(file)
	img := processing.ApplyFilter(i, kernels.EdgeDetection)
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

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200

	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// TODO 越接近边界n越大，颜色需要循环
			return getColor(0, iterations, float64(n))
		}
	}
	return color.Black
}

func getColor(min, max, value float64) color.Color {
	x := (value - min) / (max - min)
	var c color.RGBA
	switch {
	case x < 0.25: // #0000FF -> #00FFFF
		c = color.RGBA{0, uint8(255 * x / 0.25), 255, 255}
	case x < 0.5: // #00FFFF -> #00FF00
		c = color.RGBA{0, 255, uint8(255 * (0.5 - x) / 0.25), 255}
	case x < 0.75: // #00FF00 -> #FFFF00
		c = color.RGBA{uint8(255 * (x - 0.5) / 0.25), 255, 0, 255}
	default: // #FFFF00 -> #FF0000
		c = color.RGBA{255, uint8(255 * (1.0 - x) / 0.25), 0, 255}
	}
	return c
}

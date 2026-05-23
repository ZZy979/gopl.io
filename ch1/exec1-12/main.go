// Lissajous generates GIF animations of random Lissajous figures.
// Run with "web" command-line argument for web server.
// Read parameter values from the URL. For example, http://localhost:8000/?cycles=20
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type Param struct {
	Cycles  int     // number of complete x oscillator revolutions
	Res     float64 // angular resolution
	Size    int     // image canvas covers [-size..+size]
	Nframes int     // number of animation frames
	Delay   int     // delay between frames in 10ms units
}

var defaultParam = Param{5, 0.001, 100, 64, 8}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			param := defaultParam
			query := r.URL.Query()
			if v, ok := query["cycles"]; ok {
				if cycles, err := strconv.Atoi(v[0]); err == nil {
					param.Cycles = cycles
				}
			}
			lissajous(param, w)
		})
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(defaultParam, os.Stdout)
}

func lissajous(p Param, out io.Writer) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: p.Nframes}
	phase := 0.0 // phase difference
	for i := 0; i < p.Nframes; i++ {
		rect := image.Rect(0, 0, 2*p.Size+1, 2*p.Size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(p.Cycles)*2*math.Pi; t += p.Res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(p.Size+int(x*float64(p.Size)+0.5), p.Size+int(y*float64(p.Size)+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, p.Delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

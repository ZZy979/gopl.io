// Surface computes an SVG rendering of a 3-D surface function.
//
// Command-line argument:
//
//	[-web] -f {sinc|egg-box|moguls|saddle} [-w <width>] [-h <height>] [-c <cells>]
//	[-r <xyrange>] [-z <zrange>] [-zs <zscale>]
//
// Web request parameter:
//
//	http://localhost:8000/?func=sinc
//	http://localhost:8000/?func=egg-box&zrange=2&zscale=0.1
//	http://localhost:8000/?func=moguls&zscale=0.1
//	http://localhost:8000/?func=saddle&cells=50&xyrange=5&zrange=6.5&zscale=0.1
package main

import (
	"flag"
	"fmt"
	"image/color"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"slices"
	"strconv"
)

type Param struct {
	funcName     string  // function name
	width        int     // canvas width in pixels
	height       int     // canvas height in pixels
	cells        int     // number of grid cells
	xyrange      float64 // X/Y axis ranges (-xyrange..+xyrange)
	zrange       float64 // Z axis range (-zrange..+zrange)
	zScaleFactor float64 // scale factor for z
}

func (p *Param) f(x, y float64) float64 {
	switch p.funcName {
	case "sinc":
		r := math.Hypot(x, y) // distance from (0,0)
		return math.Sin(r) / r
	case "egg-box":
		return math.Sin(x) + math.Sin(y)
	case "moguls":
		return math.Sin(x) * math.Sin(y)
	case "saddle":
		return x*x - y*y
	default:
		return 0
	}
}

func (p *Param) xyscale() float64 {
	return float64(p.width) / 2 / p.xyrange // pixels per x or y unit
}

func (p *Param) zscale() float64 {
	return float64(p.height) * p.zScaleFactor // pixels per z unit
}

var defaultParam = Param{"sinc", 600, 320, 100, 30.0, 1.0, 0.4}

const angle = math.Pi / 6                           // angle of x, y axes (=30°)
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	param := defaultParam
	web := flag.Bool("web", false, "Web mode")
	flag.StringVar(&param.funcName, "f", defaultParam.funcName, "function name")
	flag.IntVar(&param.width, "w", defaultParam.width, "canvas width in pixels")
	flag.IntVar(&param.height, "h", defaultParam.height, "canvas height in pixels")
	flag.IntVar(&param.cells, "c", defaultParam.cells, "number of grid cells")
	flag.Float64Var(&param.xyrange, "r", defaultParam.xyrange, "X/Y axis ranges")
	flag.Float64Var(&param.zrange, "z", defaultParam.zrange, "Z axis range")
	flag.Float64Var(&param.zScaleFactor, "zs", defaultParam.zScaleFactor, "Z scale factor")
	flag.Parse()

	if *web {
		http.HandleFunc("/", plotHandler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}

	plot(&param, os.Stdout)
}

func plot(p *Param, out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", p.width, p.height)
	for i := 0; i < p.cells; i++ {
		for j := 0; j < p.cells; j++ {
			ax, ay, _ := corner(p, i+1, j)
			bx, by, bc := corner(p, i, j)
			cx, cy, _ := corner(p, i, j+1)
			dx, dy, _ := corner(p, i+1, j+1)
			if isAllValid([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Fprintf(out, "<polygon points='%.2f,%.2f %.2f,%.2f %.2f,%.2f %.2f,%.2f' style='fill:%s;' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, bc)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(p *Param, i, j int) (float64, float64, string) {
	// Find point (x,y) at corner of cell (i,j).
	x := p.xyrange * (float64(i)/float64(p.cells) - 0.5)
	y := p.xyrange * (float64(j)/float64(p.cells) - 0.5)

	// Compute surface height z.
	z := p.f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(p.width)/2 + (x-y)*cos30*p.xyscale()
	sy := float64(p.height)/2 + (x+y)*sin30*p.xyscale() - z*p.zscale()

	// Color of polygon (i,j)
	c := getColor(-p.zrange, p.zrange, z)
	return sx, sy, c
}

func isAllValid(s []float64) bool {
	return slices.IndexFunc(s, func(x float64) bool {
		return math.IsInf(x, 1) || math.IsInf(x, -1) || math.IsNaN(x)
	}) == -1
}

func getColor(min, max, value float64) string {
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
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

func plotHandler(w http.ResponseWriter, r *http.Request) {
	param := defaultParam
	query := r.URL.Query()
	getStringRequestParam(query, "func", &param.funcName)
	getIntRequestParam(query, "width", &param.width)
	getIntRequestParam(query, "height", &param.height)
	getIntRequestParam(query, "cells", &param.cells)
	getFloatRequestParam(query, "xyrange", &param.xyrange)
	getFloatRequestParam(query, "zrange", &param.zrange)
	getFloatRequestParam(query, "zscale", &param.zScaleFactor)

	w.Header().Set("Content-Type", "image/svg+xml")
	plot(&param, w)
}

func getStringRequestParam(q url.Values, key string, value *string) {
	if v, ok := q[key]; ok {
		*value = v[0]
	}
}

func getIntRequestParam(q url.Values, key string, value *int) {
	if v, ok := q[key]; ok {
		if vv, err := strconv.Atoi(v[0]); err == nil {
			*value = vv
		}
	}
}

func getFloatRequestParam(q url.Values, key string, value *float64) {
	if v, ok := q[key]; ok {
		if vv, err := strconv.ParseFloat(v[0], 64); err == nil {
			*value = vv
		}
	}
}

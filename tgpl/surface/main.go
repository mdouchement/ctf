package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * .4         // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes in radian
)

var sinAngle, cosAngle = math.Sin(angle), math.Cos(angle)

func main() {
	file, _ := os.Create("surface.svg")
	defer file.Close()
	w := bufio.NewWriter(file)

	header := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	w.Write([]byte(header))

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			row := fmt.Sprintf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			w.Write([]byte(row))
		}
	}
	w.Write([]byte("</svg>"))
	w.Flush()
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - .5)
	y := xyrange * (float64(j)/cells - .5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*cosAngle*xyscale
	sy := height/2 + (x+y)*sinAngle*xyscale - z*zscale
	return sx, sy
}

// sin(r)/r
// r => radius
func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

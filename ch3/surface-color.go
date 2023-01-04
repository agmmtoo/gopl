// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	//"strconv"
)

const (
	width, height = 1000, 1000            // canvas size in pixels
	cells         = 100                  // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			depth := color(i, j)
			h := 240 - depth / 1 * 360
			//h := depth
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:hsl(%f,100%%,50%%)' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, h)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sin(r) / r
	if math.IsNaN(result) {
		return 0
	}
	return result
}

func color(i, j int) float64 {
	// Find point (x,y) at corner of cell (i,j).
        x := xyrange * (float64(i)/cells - 0.5)
        y := xyrange * (float64(j)/cells - 0.5)

        // Compute surface height z.
        //z := math.Abs(f(x, y))
	z := f(x,y)
	return z
}


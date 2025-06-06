// Package main is the main program
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/pov"
)

//////////////////////////////
// MAIN
//////////////////////////////

func main() {
	scene := pov.NewScene()
	scene.AddInclude("colors.inc")
	scene.AddInclude("stones.inc")
	// scene.SetSize(800, 800)
	scene.AddLight(pov.NewLightSource(-4, 0, 1, "White"))
	scene.AddLight(pov.NewLightSource(20, 2, -5, "Gray"))
	scene.AddObject(pov.NewPlane(0.0))
	scene.AddObject(pov.NewBox(0, 0, 0, 2.5, 1.5, 2.5))
	scene.AddObject(pov.NewBackgroundRGB(0.1, 0.2, 0.5))
	scene.Camera.Position(0, 0.5, -8)

	count := 64.0
	for i := 0.0; i < count; i++ {
		t := i / count * blmath.Tau
		x := math.Cos(t) * 2.5
		z := math.Sin(t) * 2.5
		y := math.Cos(t*4)*0.5 + 0.5
		r := math.Sin(t*8)*0.1 + 0.2
		scene.AddObject(pov.NewSphere(x, y, z, r))
	}

	scene.Render("out/main.pov")

}

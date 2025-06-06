// Package main is the main program
package main

import (
	"math"

	"github.com/bit101/bitlib/blmath"
	"github.com/bit101/pov"
	"github.com/bit101/pov/tex"
)

//////////////////////////////
// MAIN
//////////////////////////////

func main() {
	scene := pov.NewScene()
	scene.AddInclude("stones.inc")
	// scene.AddInclude("woods.inc")
	scene.SetSize(1600, 1600)
	scene.SetAmbient(0.2, 0.2, 0.0)

	scene.AddLight(pov.NewLightSource(-4, 0, 1, "White"))
	scene.AddLight(pov.NewLightSource(20, 2, -5, "Gray"))

	sky := pov.NewSphere(0, 0, 0, 1000)
	sky.Texture = pov.NewPresetTexture("Bright_Blue_Sky", 100)
	scene.AddObject(sky)

	plane := pov.NewPlane(0.0)
	plane.Texture = pov.NewPresetTexture(tex.Stone15, 2)
	scene.AddObject(plane)

	box := pov.NewBox(0, 0, 0, 25, 0.1, 25)
	box.Texture = pov.NewPresetTexture(tex.Metal, 0.25)
	scene.AddObject(box)

	scene.Camera.Position(4, 2.25, -3)

	count := 512.0
	for i := 0.0; i < count; i++ {
		t := i / count * blmath.Tau
		x := math.Cos(t) * 2.5
		z := math.Sin(t) * 2.5
		y := math.Cos(t*4)*0.5 + 0.5
		r := math.Sin(t*8)*0.2 + 0.3
		s := pov.NewSphere(x, y, z, r)
		s.Texture = pov.NewPresetTexture(tex.Stone30, 2)
		scene.AddObject(s)
	}

	scene.Render("out/main.pov")

}

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
	scene.SetSize(1600, 1600)
	scene.SetAmbient(0.2, 0.2, 0.0)

	scene.AddLight(pov.NewLightSource(-4, 8, 1, "White"))
	scene.AddLight(pov.NewLightSource(20, 6, -5, "Gray"))

	sky := pov.NewSphere(0, 0, 0, 1000)
	sky.Texture = pov.NewPresetTexture("Blue_Sky")
	sky.Texture.UniScale(1000)
	scene.AddObject(sky)

	plane := pov.NewPlane(0.0)
	plane.Texture = pov.NewPresetTexture(tex.Stone33)
	plane.Texture.UniScale(5)
	scene.AddObject(plane)

	box := pov.NewBox(0, 0.5, 0, 1, 1, 1)
	box.Texture = pov.NewPresetTexture(tex.Stone38)
	box.Texture.UniScale(1)
	scene.AddObject(box)

	scene.Camera.Position(4, 2.25, -4)

	txt := pov.NewPresetTexture(tex.SilverFinish)
	count := 64.0
	for i := 0.0; i < count; i++ {
		t := i / count * blmath.Tau
		x := math.Cos(t) * 2.5
		z := math.Sin(t) * 2.5
		y := math.Cos(t*4)*0.5 + 0.5
		r := math.Sin(t*8)*0.2 + 0.3
		s := pov.NewSphere(x, y, z, r)
		s.Texture = txt
		scene.AddObject(s)
	}
	scene.Render("out/main.pov")
}

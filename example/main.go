// Package main is the main program
package main

import (
	"github.com/bit101/pov"
	"github.com/bit101/pov/tex/metal"
	"github.com/bit101/pov/tex/sky"
	"github.com/bit101/pov/tex/stone"
)

//////////////////////////////
// MAIN
//////////////////////////////

func main() {
	scene := pov.NewScene()
	scene.AddInclude("stones.inc")
	scene.AddInclude("woods.inc")
	scene.SetSize(1000, 800)
	// scene.SetAmbient(0.2, 0.2, 0.0)

	scene.AddLight(pov.NewLightSource(-4, 8, -1, "White"))
	scene.AddLight(pov.NewLightSource(20, 6, -5, "Gray"))

	s := pov.NewSphere(0, 0, 0, 10000)
	s.Texture = pov.PresetTexture(sky.BrightBlueSky)
	s.Texture.UniScale(5000)
	scene.AddObject(s)

	for i := -3.0; i <= 3.0; i += 3.0 {
		s = pov.NewSphere(i, 0, 3.5, 1.5)
		s.Texture = pov.PresetTexture(metal.MetallicFinish)
		// scene.AddObject(s)
	}

	plane := pov.NewPlane(0.0)
	plane.Texture = pov.PresetTexture(stone.Stone33)
	plane.Texture.UniScale(5)
	// scene.AddObject(plane)

	torus := pov.NewTorus(2.5, 0.5)
	torus.Texture = pov.PresetTexture(metal.MetallicFinish)
	torus.Rotate(90, 0, 0)
	torus.Translate(0, 0, 2)
	scene.AddObject(torus)

	// box := pov.NewBox(0, 0.5, 0, 1, 1, 1)
	// box.Texture = pov.PresetTexture(wood.TWood34)
	// box.Texture.Scale(0.2, 0.2, 0.2)
	// box.Scale(2, 2, 2)
	// scene.AddObject(box)

	// cone := pov.NewCylinder(0, 0, 0, 0, 3, 0, 1)
	// cone.Texture = pov.PresetTexture(stone.RedMarble)
	// scene.AddObject(cone)

	scene.Camera.Position(1, 1.5, -2.7)
	scene.Camera.LookAt(0.2, 1, 0)

	text := pov.NewText("/usr/share/fonts/TTF/DejaVuSans.ttf", "BIT-101", 1)
	text.Translate(-1.8, 0, 0)
	text.Scale(1, 1, 1000)
	text.Texture = pov.PresetTexture(metal.MetallicFinish)
	text.Texture.Scale(0.25, 0.25, 0.00025)
	scene.AddObject(text)

	// txt := pov.PresetTexture(metal.SilverFinish)
	// count := 64.0
	// for i := 0.0; i < count; i++ {
	// 	t := i / count * blmath.Tau
	// 	x := math.Cos(t) * 2.5
	// 	z := math.Sin(t) * 2.5
	// 	y := math.Cos(t*4)*0.5 + 0.5
	// 	r := math.Sin(t*8)*0.2 + 0.3
	// 	s := pov.NewSphere(x, y, z, r)
	// 	s.Texture = txt
	// 	scene.AddObject(s)
	// }
	scene.Render("out/main.pov")
}

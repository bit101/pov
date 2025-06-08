// Package main is the main program
package main

import (
	"github.com/bit101/bitlib/random"
	"github.com/bit101/pov"
	"github.com/bit101/pov/tex/metal"
	"github.com/bit101/pov/tex/stone"
	"github.com/bit101/pov/utils"
)

//////////////////////////////
// MAIN
//////////////////////////////

func main() {
	// test()
	fractalBox()
	// utils.ConvertToVideo("out", "out.mp4", 1280, 720, 30)
}

const maxDepth = 6

func fractalBox() {
	scene := pov.NewScene("out/main.pov")
	scene.AddInclude("stones.inc")
	scene.AddInclude("colors.inc")
	scene.AddInclude("woods.inc")
	scene.SetSize(1280, 720)
	scene.Add(pov.NewLightSource(-4, 6, -8, "White"))
	scene.Add(pov.NewLightSource(8, 1, -1, utils.RGB(1.0, 0.75, 0.75)))

	scene.Add(pov.NewSkySphere("Blue", "White"))

	plane := pov.NewPlane(0)
	plane.Texture = pov.PresetTexture(metal.NewPenny)
	scene.Add(plane)

	random.Seed(5)
	fractalize(-1, 0, -1, 1, 2, 1, scene, maxDepth)

	scene.Camera.Position(1.9, 0.8, -2.9)
	scene.Camera.LookAt(0, 0.9, 0)

	scene.WritePOV()
	// scene.RenderAnim()
	scene.Render()
}

func fractalize(x0, y0, z0, x1, y1, z1 float64, scene *pov.Scene, depth int) {
	if random.WeightedBool(0.45) && depth < maxDepth {
		return
	}
	if depth < 1 {
		// box := pov.NewBoxFromPoints(x0, y0, z0, x1, y1, z1)
		box := pov.NewSphere((x0+x1)/2, (y0+y1)/2, (z0+z1)/2, (z1-z0)/4*random.FloatRange(1, 10))
		if random.Boolean() {
			box.Texture = pov.PresetTexture(stone.RedMarble)
		} else {
			box.Texture = pov.PresetTexture(stone.Stone10)
		}
		box.StringTransform("rotate <0, 360*clock, 0>")
		scene.Add(box)
	} else {
		w := (x1 - x0) / 2
		h := (y1 - y0) / 2
		d := (z1 - z0) / 2
		fractalize(x0, y0, z0, x0+w, y0+h, z0+d, scene, depth-1)
		fractalize(x0+w, y0, z0, x0+w+w, y0+h, z0+d, scene, depth-1)
		fractalize(x0+w, y0+h, z0, x0+w+w, y0+h+h, z0+d, scene, depth-1)
		fractalize(x0+w, y0+h, z0+d, x0+w+w, y0+h+h, z0+d+d, scene, depth-1)
		fractalize(x0+w, y0, z0+d, x0+w+w, y0+h, z0+d+d, scene, depth-1)
		fractalize(x0, y0+h, z0, x0+w, y0+h+h, z0+d, scene, depth-1)
		fractalize(x0, y0+h, z0+d, x0+w, y0+h+h, z0+d+d, scene, depth-1)
		fractalize(x0, y0, z0+d, x0+w, y0+h, z0+d+d, scene, depth-1)
	}
}

func test() {
	scene := pov.NewScene("out/main.pov")
	scene.AddInclude("stones.inc")
	scene.AddInclude("woods.inc")
	scene.SetSize(1000, 1000)
	scene.SetAmbient(0.5, 0.5, 0.5)

	scene.Add(pov.NewLightSource(-4, 8, -1, "White"))
	scene.Add(pov.NewLightSource(20, 6, -5, "Gray"))

	scene.Camera.Position(0, 3, -3)

	box := pov.NewBox(0, 0, 0, 1, 1, 1)
	box.Texture = pov.PresetTexture(stone.RedMarble)
	box.Rotate(0, 45, 0)
	scene.Add(box)

	s := pov.NewSphere(0, 1, 0, 0.5)
	s.Texture = pov.PresetTexture(stone.Stone10)
	scene.Add(s)

	scene.WritePOV()
	scene.Render()
}

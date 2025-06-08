// Package main is the main program
package main

import (
	"github.com/bit101/bitlib/random"
	"github.com/bit101/pov"
	"github.com/bit101/pov/tex/metal"
	"github.com/bit101/pov/tex/sky"
	"github.com/bit101/pov/tex/stone"
	"github.com/bit101/pov/utils"
)

//////////////////////////////
// MAIN
//////////////////////////////

func main() {
	// test()
	fractalBox()
}

const maxDepth = 6

func fractalBox() {
	scene := pov.NewScene()
	scene.AddInclude("stones.inc")
	scene.AddInclude("colors.inc")
	scene.AddInclude("woods.inc")
	scene.SetSize(1600, 1600)
	scene.Add(pov.NewLightSource(-4, 6, -8, "White"))
	scene.Add(pov.NewLightSource(8, 1, -1, utils.RGB(1.0, 0.75, 0.75)))

	scene.Add(pov.NewSkySphere("Blue", "White"))

	plane := pov.NewPlane(0)
	plane.Texture = pov.PresetTexture(metal.SilverFinish)
	scene.Add(plane)

	// base := pov.NewCylinder(0, 0, 0, 0, 0.1, 0, 1.5)
	// base.Texture = pov.PresetTexture(stone.Stone33)
	// scene.Add(base)

	random.Seed(6)
	fractalize(-1, 0, -1, 1, 2, 1, scene, maxDepth)

	scene.Camera.Position(2, 3.5, -3)
	scene.Camera.LookAt(-0.5, 0.5, 0)

	scene.Render("out/main.pov")
}

func fractalize(x0, y0, z0, x1, y1, z1 float64, scene *pov.Scene, depth int) {
	if random.WeightedBool(0.25) && depth < maxDepth {
		return
	}
	if depth < 1 {
		box := pov.NewBoxFromPoints(x0, y0, z0, x1, y1, z1)
		// box := pov.NewSphere((x0+x1)/2, (y0+y1)/2, (z0+z1)/2, (z1-z0)/2)

		// color := fmt.Sprintf("rgb <%f, %f, %f>", random.Float(), random.Float(), random.Float())
		box.Texture = pov.PresetTexture(stone.Stone10)
		// box.Texture = pov.ColorTexture(utils.RandomGray())
		// box.Texture.UniScale(0.25)
		// box.Rotate(0, 0, 45)
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
	scene := pov.NewScene()
	scene.AddInclude("stones.inc")
	scene.AddInclude("woods.inc")
	scene.SetSize(1000, 1000)
	scene.SetAmbient(0.5, 0.5, 0.5)

	scene.Add(pov.NewLightSource(-4, 8, -1, "White"))
	scene.Add(pov.NewLightSource(20, 6, -5, "Gray"))

	scene.Camera.Position(0, 0, -5)

	s := pov.NewSphere(0, 0, 0, 1000)
	s.Texture = pov.PresetTexture(sky.BrightBlueSky)
	s.Texture.UniScale(10000)
	// scene.Add(s)

	scene.Add(pov.NewBackground("Red"))
	scene.Add(&pov.Fog{})

	blobs := []*pov.BlobSphere{}
	for range 1000 {
		size := random.FloatRange(0.1, 0.3)
		blob := pov.NewBlobSphere(
			random.FloatRange(-2, 2),
			random.FloatRange(-2, 2),
			random.FloatRange(-1, 1),
			size,
			size,
		)
		blobs = append(blobs, blob)
	}
	blob := pov.NewBlob(
		0.2,
		blobs,
	)
	blob.Texture = pov.PresetTexture(stone.Stone15)
	scene.Add(blob)

	scene.Render("out/main.pov")
}

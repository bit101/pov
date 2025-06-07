// Package main is the main program
package main

import (
	"github.com/bit101/bitlib/random"
	"github.com/bit101/pov"
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
	scene.SetSize(1000, 1000)
	scene.SetAmbient(0.5, 0.5, 0.5)

	scene.AddLight(pov.NewLightSource(-4, 8, -1, "White"))
	scene.AddLight(pov.NewLightSource(20, 6, -5, "Gray"))

	scene.Camera.Position(0, 0, -5)

	s := pov.NewSphere(0, 0, 0, 1000)
	s.Texture = pov.PresetTexture(sky.BrightBlueSky)
	s.Texture.UniScale(10000)
	scene.AddObject(s)

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
	scene.AddObject(blob)

	scene.Render("out/main.pov")
}

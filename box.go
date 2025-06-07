// Package pov ...
package pov

import "fmt"

// Box ...
type Box struct {
	cornerA, cornerB Vector3
	transform        Transform
	Texture          *Texture
}

// UniScale ...
func (b *Box) UniScale(scale float64) {
	b.transform.UniScale(scale)
}

// Scale ...
func (b *Box) Scale(x, y, z float64) {
	b.transform.Scale(x, y, z)
}

// Translate ...
func (b *Box) Translate(x, y, z float64) {
	b.transform.Translate(x, y, z)
}

// Rotate ...
func (b *Box) Rotate(x, y, z float64) {
	b.transform.Rotate(x, y, z)
}

// NewBox ...
func NewBox(x, y, z, w, h, d float64) *Box {
	return &Box{
		Vector3{x - w/2, y - h/2, z - d/2},
		Vector3{x + w/2, y + h/2, z + d/2},
		Transform{},
		NewColorTexture("White"),
	}
}

func (b *Box) String() string {
	return fmt.Sprintf(`
box {
  %s
  %s
  %s
  %s
}
`, b.cornerA.String(), b.cornerB.String(), b.Texture.String(), b.transform.String())
}

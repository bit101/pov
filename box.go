// Package pov ...
package pov

import "fmt"

// Box ...
type Box struct {
	cornerA, cornerB Vector3
	Texture          *Texture
}

// NewBox ...
func NewBox(x, y, z, w, h, d float64) *Box {
	return &Box{
		Vector3{x - w/2, y - h/2, z - d/2},
		Vector3{x + w/2, y + h/2, z + d/2},
		NewColorTexture("White", UniScale(1)),
	}
}

func (b *Box) String() string {
	return fmt.Sprintf(`
box {
  %s
  %s
  %s
}
`, b.cornerA.String(), b.cornerB.String(), b.Texture.String())
}

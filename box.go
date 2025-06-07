// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Box ...
type Box struct {
	Object
	cornerA, cornerB Vector3
}

// NewBox ...
func NewBox(x, y, z, w, h, d float64) *Box {
	return &Box{
		NewObject(),
		Vector3{x - w/2, y - h/2, z - d/2},
		Vector3{x + w/2, y + h/2, z + d/2},
	}
}

// NewBoxFromPoints ...
func NewBoxFromPoints(x0, y0, z0, x1, y1, z1 float64) *Box {
	return &Box{
		NewObject(),
		Vector3{x0, y0, z0},
		Vector3{x1, y1, z1},
	}
}

func (b *Box) String() string {
	str := fmt.Sprintf(`
box {
  %s
  %s
  %s
  %s
}
`, b.cornerA.String(), b.cornerB.String(), b.Texture.String(), b.transform.String())
	return utils.RemoveEmptyLines(str)
}

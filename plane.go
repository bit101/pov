// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Plane ...
type Plane struct {
	Object
	y float64
}

// NewPlane ...
func NewPlane(y float64) *Plane {
	return &Plane{NewObject(), y}
}

func (p *Plane) String() string {
	str := fmt.Sprintf(`
plane {
  y
  %f
  %s
  %s
}`, p.y, p.Texture.String(), p.transform.String())
	return utils.RemoveEmptyLines(str)
}

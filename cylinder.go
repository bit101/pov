// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Cylinder ...
type Cylinder struct {
	Object
	location0 string
	location1 string
	radius    string
}

// NewCylinder ...
func NewCylinder(x0, y0, z0, x1, y1, z1, radius float64) *Cylinder {
	return &Cylinder{NewObject(), NewVector3(x0, y0, z0).String(), NewVector3(x1, y1, z1).String(), utils.Ftos(radius)}
}

// String ...
func (c *Cylinder) String() string {
	str := fmt.Sprintf(`
cylinder {
  %s
  %s
  %s
  %s
  %s
}`, c.location0, c.location1, c.radius, c.Texture.String(), c.transform.String())
	return utils.RemoveEmptyLines(str)
}

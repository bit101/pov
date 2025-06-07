// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Cylinder ...
type Cylinder struct {
	Object
	location0 Vector3
	location1 Vector3
	radius    float64
}

// NewCylinder ...
func NewCylinder(x0, y0, z0, x1, y1, z1, radius float64) Cylinder {
	return Cylinder{NewObject(), Vector3{x0, y0, z0}, Vector3{x1, y1, z1}, radius}
}

// String ...
func (c Cylinder) String() string {
	str := fmt.Sprintf(`
cylinder {
  %s
  %s
  %f
  %s
  %s
}`, c.location0.String(), c.location1.String(), c.radius, c.Texture.String(), c.transform.String())
	return utils.RemoveEmptyLines(str)
}

// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Cone ...
type Cone struct {
	Object
	location0 Vector3
	radius0   float64
	location1 Vector3
	radius1   float64
}

// NewCone ...
func NewCone(x0, y0, z0, radius0, x1, y1, z1, radius1 float64) Cone {
	return Cone{NewObject(), Vector3{x0, y0, z0}, radius0, Vector3{x1, y1, z1}, radius1}
}

// String ...
func (c Cone) String() string {
	str := fmt.Sprintf(`
cone {
  %s
  %f
  %s
  %f
  %s
  %s
}`, c.location0.String(), c.radius0, c.location1.String(), c.radius1, c.Texture.String(), c.transform.String())
	return utils.RemoveEmptyLines(str)
}

// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Cone ...
type Cone struct {
	Object
	location0 string
	radius0   string
	location1 string
	radius1   string
}

// NewCone ...
func NewCone(x0, y0, z0, radius0, x1, y1, z1, radius1 float64) *Cone {
	return &Cone{
		NewObject(),
		NewVector3(x0, y0, z0).String(),
		utils.Ftos(radius0),
		NewVector3(x1, y1, z1).String(),
		utils.Ftos(radius1),
	}
}

// String ...
func (c *Cone) String() string {
	str := fmt.Sprintf(`
cone {
  %s
  %s
  %s
  %s
  %s
  %s
}`, c.location0, c.radius0, c.location1, c.radius1, c.Texture.String(), c.transform.String())
	return utils.RemoveEmptyLines(str)
}

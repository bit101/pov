// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Torus ...
type Torus struct {
	Object
	majorRadius string
	minorRadius string
}

// NewTorus ...
func NewTorus(majorRadius, minorRadius float64) *Torus {
	return &Torus{NewObject(), utils.Ftos(majorRadius), utils.Ftos(minorRadius)}
}

// String ...
func (t *Torus) String() string {
	str := fmt.Sprintf(`
torus {
  %s
  %s
  %s
  %s
}`, t.majorRadius, t.minorRadius, t.Texture.String(), t.transform.String())
	return utils.RemoveEmptyLines(str)
}

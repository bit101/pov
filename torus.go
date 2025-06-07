// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Torus ...
type Torus struct {
	Object
	majorRadius float64
	minorRadius float64
}

// NewTorus ...
func NewTorus(majorRadius, minorRadius float64) *Torus {
	return &Torus{NewObject(), majorRadius, minorRadius}
}

// String ...
func (t *Torus) String() string {
	str := fmt.Sprintf(`
torus {
  %f
  %f
  %s
  %s
}`, t.majorRadius, t.minorRadius, t.Texture.String(), t.transform.String())
	return utils.RemoveEmptyLines(str)
}

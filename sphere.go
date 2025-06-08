// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Sphere ...
type Sphere struct {
	Object
	location string
	radius   string
}

// NewSphere ...
func NewSphere(x, y, z, radius float64) *Sphere {
	return &Sphere{NewObject(), NewVector3(x, y, z).String(), fmt.Sprintf("%f", radius)}
}

// String ...
func (s *Sphere) String() string {
	str := fmt.Sprintf(`
sphere {
  %s
  %s
  %s
  %s
}`, s.location, s.radius, s.Texture.String(), s.transform.String())
	return utils.RemoveEmptyLines(str)
}

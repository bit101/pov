// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Sphere ...
type Sphere struct {
	Object
	location Vector3
	radius   float64
}

// NewSphere ...
func NewSphere(x, y, z, radius float64) *Sphere {
	return &Sphere{NewObject(), Vector3{x, y, z}, radius}
}

// String ...
func (s *Sphere) String() string {
	str := fmt.Sprintf(`
sphere {
  %s
  %f
  %s
  %s
}`, s.location.String(), s.radius, s.Texture.String(), s.transform.String())
	return utils.RemoveEmptyLines(str)
}

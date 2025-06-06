// Package pov ...
package pov

import "fmt"

// Sphere ...
type Sphere struct {
	location Vector3
	radius   float64
}

// NewSphere ...
func NewSphere(x, y, z, radius float64) Sphere {
	return Sphere{Vector3{x, y, z}, radius}
}

// String ...
func (s Sphere) String() string {
	return fmt.Sprintf(`
sphere {
  %s
	%f
  texture {
    T_Stone30
  }
}
`, s.location.String(), s.radius)
}

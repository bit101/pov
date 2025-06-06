// Package pov ...
package pov

import "fmt"

// Sphere ...
type Sphere struct {
	location  Vector3
	radius    float64
	Texture   *Texture
	Transform Transform
}

// NewSphere ...
func NewSphere(x, y, z, radius float64) Sphere {
	return Sphere{Vector3{x, y, z}, radius, NewColorTexture("White", UniScale(1)), UniScale(1)}
}

// String ...
func (s Sphere) String() string {
	return fmt.Sprintf(`
sphere {
  %s
  %f
  %s
  %s
}
`, s.location.String(), s.radius, s.Texture.String(), s.Transform.String())
}

/*


texture
  - pigment
		- color


color
	string: Red
	values: red 1.0 green 0.8 blue 0.5
	rgb: rgb <1, 0.8, 0.5>



*/

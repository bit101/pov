// Package pov ...
package pov

import "fmt"

// Sphere ...
type Sphere struct {
	location  Vector3
	radius    float64
	Texture   *Texture
	transform Transform
}

// NewSphere ...
func NewSphere(x, y, z, radius float64) Sphere {
	return Sphere{Vector3{x, y, z}, radius, NewColorTexture("White"), Transform{}}
}

// UniScale ...
func (s *Sphere) UniScale(scale float64) {
	s.transform.UniScale(scale)
}

// Scale ...
func (s *Sphere) Scale(x, y, z float64) {
	s.transform.Scale(x, y, z)
}

// Translate ...
func (s *Sphere) Translate(x, y, z float64) {
	s.transform.Translate(x, y, z)
}

// Rotate ...
func (s *Sphere) Rotate(x, y, z float64) {
	s.transform.Rotate(x, y, z)
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
`, s.location.String(), s.radius, s.Texture.String(), s.transform.String())
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

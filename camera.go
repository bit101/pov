// Package pov ...
package pov

import "fmt"

// Camera ...
type Camera struct {
	location Vector3
	lookAt   Vector3
}

// NewCamera ...
func NewCamera(x, y, z float64) *Camera {
	return &Camera{
		Vector3{x, y, z},
		Vector3{0, 0, 0},
	}
}

// LookAt ...
func (c *Camera) LookAt(x, y, z float64) {
	c.lookAt = Vector3{x, y, z}
}

// Position ...
func (c *Camera) Position(x, y, z float64) {
	c.location = Vector3{x, y, z}
}

// String ...
func (c *Camera) String() string {
	return fmt.Sprintf(`
camera {
  right x * image_width / image_height
  location %s
  look_at  %s
}`,
		c.location.String(),
		c.lookAt.String(),
	)
}

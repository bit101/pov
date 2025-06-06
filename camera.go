// Package pov ...
package pov

import "fmt"

// Camera ...
type Camera struct {
	location Location
	lookAt   Location
}

// NewCamera ...
func NewCamera(x, y, z float64) *Camera {
	return &Camera{
		Location{x, y, z},
		Location{0, 0, 0},
	}
}

// LookAt ...
func (c *Camera) LookAt(x, y, z float64) {
	c.lookAt = Location{x, y, z}
}

// Position ...
func (c *Camera) Position(x, y, z float64) {
	c.location = Location{x, y, z}
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

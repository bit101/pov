// Package pov ...
package pov

import "fmt"

// Camera ...
type Camera struct {
	location string
	lookAt   string
}

// NewCamera ...
func NewCamera(x, y, z float64) *Camera {
	return &Camera{
		NewVector3(x, y, z).String(),
		NewVector3(0, 0, 0).String(),
	}
}

// LookAt ...
func (c *Camera) LookAt(x, y, z float64) {
	c.lookAt = NewVector3(x, y, z).String()
}

// Position ...
func (c *Camera) Position(x, y, z float64) {
	c.location = NewVector3(x, y, z).String()
}

// String ...
func (c *Camera) String() string {
	return fmt.Sprintf(`
camera {
  right x * image_width / image_height
  location %s
  look_at  %s
}`,
		c.location,
		c.lookAt,
	)
}

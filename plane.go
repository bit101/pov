// Package pov ...
package pov

import "fmt"

// Plane ...
type Plane struct {
	Object
	y float64
}

// NewPlane ...
func NewPlane(y float64) *Plane {
	return &Plane{NewObject(), y}
}

func (p *Plane) String() string {
	return fmt.Sprintf(`
plane {
  y
  %f
  %s
}`, p.y, p.Texture.String())
}

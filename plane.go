// Package pov ...
package pov

import "fmt"

// Plane ...
type Plane struct {
	y float64
}

// NewPlane ...
func NewPlane(y float64) *Plane {
	return &Plane{y}
}

func (p *Plane) String() string {
	return fmt.Sprintf(`
plane {
  y, %f
  texture {
    T_Stone27
  }
}`, p.y)
}

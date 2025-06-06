// Package pov ...
package pov

import "fmt"

// Plane ...
type Plane struct {
	y       float64
	Texture *Texture
}

// NewPlane ...
func NewPlane(y float64) *Plane {
	return &Plane{y, NewColorTexture("White", UniScale(1))}
}

func (p *Plane) String() string {
	return fmt.Sprintf(`
plane {
  y
  %f
  %s
}`, p.y, p.Texture.String())
}

// Package pov ...
package pov

import "fmt"

// SkySphere ...
type SkySphere struct {
	colorA, colorB string
}

// NewSkySphere ...
func NewSkySphere(colorA, colorB string) *SkySphere {
	return &SkySphere{colorA, colorB}
}

// String ...
func (g *SkySphere) String() string {
	return fmt.Sprintf(`
sky_sphere {
  pigment {
    gradient y
    color_map {
      [ 0.0 %s ]
      [ 1.0 %s ]
    }
    scale 2
    translate 1
  }
  // emission rgb <0.8,0.8,1>
}`, g.colorA, g.colorB)
}

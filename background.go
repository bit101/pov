// Package pov ...
package pov

import "fmt"

// Background ...
type Background struct {
	color string
}

// NewBackground ...
func NewBackground(color string) *Background {
	return &Background{color}
}

// NewBackgroundRGB ...
func NewBackgroundRGB(r, g, b float64) *Background {
	color := Vector3{r, g, b}
	return &Background{color.String()}
}

// String ...
func (b *Background) String() string {
	return fmt.Sprintf("background {color %s}", b.color)
}

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

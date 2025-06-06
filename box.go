// Package pov ...
package pov

import "fmt"

// Box ...
type Box struct {
	cornerA, cornerB *Location
}

// NewBox ...
func NewBox(x, y, z, w, h, d float64) *Box {
	return &Box{
		NewLocation(x-w/2, y-h/2, z-d/2),
		NewLocation(x+w/2, y+h/2, z+d/2),
	}
}

func (b *Box) String() string {
	return fmt.Sprintf(`
box {
  %s
  %s
  texture {
    T_Stone4
  }
}
`, b.cornerA, b.cornerB)
}

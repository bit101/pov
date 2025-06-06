// Package pov ...
package pov

import "fmt"

// Location ...
type Location struct {
	x, y, z float64
}

// NewLocation ...
func NewLocation(x, y, z float64) *Location {
	return &Location{x, y, z}
}

// String ...
func (l *Location) String() string {
	return fmt.Sprintf("<%f, %f, %f>", l.x, l.y, l.z)
}

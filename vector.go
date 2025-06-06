// Package pov ...
package pov

import "fmt"

// Vector3 ...
type Vector3 struct {
	x, y, z float64
}

// NewVector3 ...
func NewVector3(x, y, z float64) *Vector3 {
	return &Vector3{x, y, z}
}

// String ...
func (l *Vector3) String() string {
	return fmt.Sprintf("<%f, %f, %f>", l.x, l.y, l.z)
}

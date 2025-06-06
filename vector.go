// Package pov ...
package pov

import "fmt"

// Vector3 ...
type Vector3 struct {
	x, y, z float64
}

// String ...
func (l *Vector3) String() string {
	return fmt.Sprintf("<%f, %f, %f>", l.x, l.y, l.z)
}

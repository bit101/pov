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
	return &Background{fmt.Sprintf("<%f, %f, %f>", r, g, b)}
}

// String ...
func (b *Background) String() string {
	return fmt.Sprintf("background {color %s}", b.color)
}

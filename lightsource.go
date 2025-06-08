// Package pov ...
package pov

import "fmt"

// LightSource ...
type LightSource struct {
	location string
	color    string
}

// NewLightSource ...
func NewLightSource(x, y, z float64, color string) *LightSource {
	return &LightSource{NewVector3(x, y, z).String(), color}
}

func (l *LightSource) String() string {
	return fmt.Sprintf(`
light_source {
  %s
  color %s
}`, l.location, l.color)
}

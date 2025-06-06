// Package pov ...
package pov

import "fmt"

// LightSource ...
type LightSource struct {
	location *Location
	color    string
}

// NewLightSource ...
func NewLightSource(x, y, z float64, color string) *LightSource {
	return &LightSource{NewLocation(x, y, z), color}
}

func (l *LightSource) String() string {
	return fmt.Sprintf(`
light_source { %s color %s}
`, l.location.String(), l.color)
}

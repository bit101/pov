// Package pov ...
package pov

import "fmt"

// LightSource ...
type LightSource struct {
	location Vector3
	color    string
}

// NewLightSource ...
func NewLightSource(x, y, z float64, color string) *LightSource {
	return &LightSource{Vector3{x, y, z}, color}
}

func (l *LightSource) String() string {
	return fmt.Sprintf(`
light_source { %s color %s}
`, l.location.String(), l.color)
}

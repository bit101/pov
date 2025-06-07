// Package pov ...
package pov

import "fmt"

// Transform ...
type Transform struct {
	transforms []string
}

// UniScale ...
func (t *Transform) UniScale(scale float64) {
	t.transforms = append(t.transforms, fmt.Sprintf("scale %f", scale))
}

// Scale ...
func (t *Transform) Scale(x, y, z float64) {
	t.transforms = append(t.transforms, fmt.Sprintf("scale <%f, %f, %f>", x, y, z))
}

// Translate ...
func (t *Transform) Translate(x, y, z float64) {
	t.transforms = append(t.transforms, fmt.Sprintf("translate <%f, %f, %f>", x, y, z))
}

// Rotate ...
func (t *Transform) Rotate(x, y, z float64) {
	t.transforms = append(t.transforms, fmt.Sprintf("rotate <%f, %f, %f>", x, y, z))
}

// String ...
func (t *Transform) String() string {
	str := ""
	for _, transform := range t.transforms {
		str += transform
	}
	return str
}

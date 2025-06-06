// Package pov ...
package pov

import "fmt"

// Transform ...
type Transform struct {
	scale     *float64
	scale3    *Vector3
	translate *Vector3
	rotate    *Vector3
}

// UniScale ...
func UniScale(scale float64) Transform {
	return Transform{scale: &scale}
}

// Scale ...
func Scale(x, y, z float64) Transform {
	return Transform{scale3: &Vector3{x, y, z}}
}

// Translate ...
func Translate(x, y, z float64) Transform {
	return Transform{translate: &Vector3{x, y, z}}
}

// Rotate ...
func Rotate(x, y, z float64) Transform {
	return Transform{rotate: &Vector3{x, y, z}}
}

// Scale ...
func (t *Transform) Scale(scale float64) {
	t.scale3 = nil
	t.scale = &scale
}

// Scale3 ...
func (t *Transform) Scale3(x, y, z float64) {
	t.scale = nil
	t.scale3 = NewVector3(x, y, z)
}

// Rotate ...
func (t *Transform) Rotate(x, y, z float64) {
	t.rotate = NewVector3(x, y, z)
}

// Translate ...
func (t *Transform) Translate(x, y, z float64) {
	t.translate = NewVector3(x, y, z)
}

// String ...
func (t *Transform) String() string {
	str := ""
	if t.scale != nil {
		str += fmt.Sprintf("scale %f\n", *t.scale)
	} else if t.scale3 != nil {
		str += fmt.Sprintf("scale %s\n", t.scale3.String())
	}

	if t.translate != nil {
		str += fmt.Sprintf("translate %s\n", t.translate.String())
	}

	if t.rotate != nil {
		str += fmt.Sprintf("rotate %s\n", t.rotate.String())
	}

	return str

}

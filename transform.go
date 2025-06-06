// Package pov ...
package pov

import "fmt"

// Transform ...
type Transform struct {
	Scale     *float64
	Scale3    *Vector3
	Translate *Vector3
	Rotate    *Vector3
}

// UniScale ...
func UniScale(scale float64) Transform {
	return Transform{Scale: &scale}
}

// Scale ...
func Scale(x, y, z float64) Transform {
	return Transform{Scale3: &Vector3{x, y, z}}
}

// Translate ...
func Translate(x, y, z float64) Transform {
	return Transform{Translate: &Vector3{x, y, z}}
}

// Rotate ...
func Rotate(x, y, z float64) Transform {
	return Transform{Rotate: &Vector3{x, y, z}}
}

// String ...
func (t *Transform) String() string {
	str := ""
	if t.Scale != nil {
		str += fmt.Sprintf("scale %f\n", *t.Scale)
	} else if t.Scale3 != nil {
		str += fmt.Sprintf("scale %s\n", t.Scale3.String())
	}

	if t.Translate != nil {
		str += fmt.Sprintf("translate %s\n", t.Translate.String())
	}

	if t.Rotate != nil {
		str += fmt.Sprintf("rotate %s\n", t.Rotate.String())
	}

	return str

}

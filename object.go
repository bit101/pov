// Package pov ...
package pov

// Object ...
type Object struct {
	transform Transform
	Texture   *Texture
}

// NewObject ...
func NewObject() Object {
	return Object{
		Transform{},
		ColorTexture("White"),
	}
}

// UniScale ...
func (o *Object) UniScale(scale float64) {
	o.transform.UniScale(scale)
}

// Scale ...
func (o *Object) Scale(x, y, z float64) {
	o.transform.Scale(x, y, z)
}

// Translate ...
func (o *Object) Translate(x, y, z float64) {
	o.transform.Translate(x, y, z)
}

// Rotate ...
func (o *Object) Rotate(x, y, z float64) {
	o.transform.Rotate(x, y, z)
}

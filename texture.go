// Package pov ...
package pov

import "fmt"

// Texture ...
type Texture struct {
	preset    string
	color     string
	transform Transform
}

// PresetTexture ...
func PresetTexture(preset string) *Texture {
	return &Texture{preset: preset, transform: Transform{}}
}

// ColorTexture ...
func ColorTexture(color string) *Texture {
	return &Texture{color: color, transform: Transform{}}
}

// RGBTexture ...
func RGBTexture(r, g, b float64) *Texture {
	return &Texture{color: fmt.Sprintf("rgb %s" + NewVector3(r, g, b).String()), transform: Transform{}}
}

// UniScale ...
func (t *Texture) UniScale(scale float64) {
	t.transform.UniScale(scale)
}

// Scale ...
func (t *Texture) Scale(x, y, z float64) {
	t.transform.Scale(x, y, z)
}

// Translate ...
func (t *Texture) Translate(x, y, z float64) {
	t.transform.Translate(x, y, z)
}

// Rotate ...
func (t *Texture) Rotate(x, y, z float64) {
	t.transform.Rotate(x, y, z)
}

// String ...
func (t *Texture) String() string {
	if t.preset != "" {
		return fmt.Sprintf(`
  texture {
    %s
    %s
  }`, t.preset, t.transform.String())
	}

	if t.color != "" {
		return fmt.Sprintf(`
  texture {
    pigment { color %s }
    %s
  }`, t.color, t.transform.String())
	}

	return "texture { pigment { color White } }"
}

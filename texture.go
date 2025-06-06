// Package pov ...
package pov

import "fmt"

// Texture ...
type Texture struct {
	preset    string
	color     string
	Transform Transform
	rgb       *Vector3
}

// NewPresetTexture ...
func NewPresetTexture(preset string, transform Transform) *Texture {
	return &Texture{preset: preset, Transform: transform}
}

// NewColorTexture ...
func NewColorTexture(color string, transform Transform) *Texture {
	return &Texture{color: color, Transform: transform}
}

// NewRGBTexture ...
func NewRGBTexture(r, g, b float64, transform Transform) *Texture {
	return &Texture{rgb: &Vector3{r, g, b}, Transform: transform}
}

// String ...
func (t *Texture) String() string {
	if t.preset != "" {
		return fmt.Sprintf(`
  texture {
    %s
    %s
  }`, t.preset, t.Transform.String())
	}

	if t.color != "" {
		return fmt.Sprintf(`
  texture {
    pigment { color %s }
    %s
  }`, t.color, t.Transform.String())
	}

	if t.rgb != nil {
		return fmt.Sprintf(`
  texture {
    pigment { rgb %s }
    %s
  }`, t.rgb.String(), t.Transform.String())
	}
	return "texture { pigment { color White } }"
}

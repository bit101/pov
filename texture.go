// Package pov ...
package pov

import "fmt"

// Texture ...
type Texture struct {
	preset string
	color  string
	scale  float64
	rgb    *Vector3
}

// NewPresetTexture ...
func NewPresetTexture(preset string, scale float64) *Texture {
	return &Texture{preset: preset, scale: scale}
}

// NewColorTexture ...
func NewColorTexture(color string, scale float64) *Texture {
	return &Texture{color: color, scale: scale}
}

// NewRGBTexture ...
func NewRGBTexture(r, g, b float64, scale float64) *Texture {
	return &Texture{rgb: &Vector3{r, g, b}, scale: scale}
}

// String ...
func (t *Texture) String() string {
	if t.preset != "" {
		return fmt.Sprintf(`
  texture {
    %s
    scale <%f, %f, %f>
  }`, t.preset, t.scale, t.scale, t.scale)
	}

	if t.color != "" {
		return fmt.Sprintf(`
  texture {
    pigment { color %s }
    scale <%f, %f, %f>
  }`, t.color, t.scale, t.scale, t.scale)
	}

	if t.rgb != nil {
		return fmt.Sprintf(`
  texture {
    pigment { rgb %s }
    scale <%f, %f, %f>
  }`, t.rgb.String(), t.scale, t.scale, t.scale)
	}
	return "texture { pigment { color White } }"
}

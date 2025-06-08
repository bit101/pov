// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Text ...
type Text struct {
	Object
	font      string
	text      string
	thickness string
}

// NewText ...
func NewText(font, text string, thickness float64) *Text {
	return &Text{
		NewObject(),
		font,
		text,
		utils.Ftos(thickness),
	}
}

func (t *Text) String() string {
	str := fmt.Sprintf(`
text {
  ttf %q
  %q
  %s
  %s
  %s
  %s
}
`, t.font, t.text, t.thickness, "0.0", t.Texture.String(), t.transform.String())
	return utils.RemoveEmptyLines(str)
}

// Package pov ...
package pov

import (
	"fmt"
	"strings"

	"github.com/bit101/pov/utils"
)

// Merge ...
type Merge struct {
	Object
	objects []Renderable
}

// NewMerge ...
func NewMerge(objects []Renderable) *Merge {
	return &Merge{NewObject(), objects}
}

// String ...
func (m *Merge) String() string {
	builder := strings.Builder{}
	for _, obj := range m.objects {
		builder.WriteString(obj.String() + "\n")
	}
	str := fmt.Sprintf(`
merge {
  %s
  %s
  %s
}`, builder.String(), m.Texture.String(), m.transform.String())
	return utils.RemoveEmptyLines(str)
}

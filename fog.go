// Package pov ...
package pov

import "github.com/bit101/pov/utils"

// Fog ...
type Fog struct {
}

// String ...
func (f *Fog) String() string {
	str := `
fog {
	distance 10
	color rgbt<1, 1, 1, 0.5>
}`
	return utils.RemoveEmptyLines(str)
}

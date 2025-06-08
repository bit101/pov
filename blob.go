// Package pov ...
package pov

import (
	"fmt"
	"strings"

	"github.com/bit101/pov/utils"
)

// Blob ...
type Blob struct {
	Object
	threshold string
	spheres   []*BlobSphere
}

// NewBlob ...
func NewBlob(threshold float64, spheres []*BlobSphere) *Blob {
	return &Blob{NewObject(), utils.Ftos(threshold), spheres}
}

// String ...
func (b *Blob) String() string {
	builder := strings.Builder{}
	for _, obj := range b.spheres {
		builder.WriteString(obj.String() + "\n")
	}
	str := fmt.Sprintf(`
blob {
  %s
  %s
  %s
  %s
}`, b.threshold, builder.String(), b.Texture.String(), b.transform.String())
	return utils.RemoveEmptyLines(str)
}

//////////////////////////////
// BlobSpheres
//////////////////////////////

// BlobSphere ...
type BlobSphere struct {
	Sphere
	strength string
}

// NewBlobSphere ...
func NewBlobSphere(x, y, z, radius, strength float64) *BlobSphere {
	return &BlobSphere{
		*NewSphere(x, y, z, radius),
		utils.Ftos(strength),
	}
}

func (b *BlobSphere) String() string {
	str := fmt.Sprintf(`
sphere {
  %s
  %s
  %s
}`, b.location, b.radius, b.strength)
	return utils.RemoveEmptyLines(str)
}

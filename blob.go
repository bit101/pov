// Package pov ...
package pov

import (
	"fmt"

	"github.com/bit101/pov/utils"
)

// Blob ...
type Blob struct {
	Object
	threshold float64
	spheres   []*BlobSphere
}

// NewBlob ...
func NewBlob(threshold float64, spheres []*BlobSphere) *Blob {
	return &Blob{NewObject(), threshold, spheres}
}

// String ...
func (b *Blob) String() string {
	spheres := ""
	for _, s := range b.spheres {
		spheres += s.String() + "\n"
	}
	str := fmt.Sprintf(`
blob {
  threshold %f
  %s
  %s
  %s
}`, b.threshold, spheres, b.Texture.String(), b.transform.String())
	return utils.RemoveEmptyLines(str)
}

//////////////////////////////
// BlobSpheres
//////////////////////////////

// BlobSphere ...
type BlobSphere struct {
	Sphere
	strength float64
}

// NewBlobSphere ...
func NewBlobSphere(x, y, z, radius, strength float64) *BlobSphere {
	return &BlobSphere{
		*NewSphere(x, y, z, radius),
		strength,
	}
}

func (b *BlobSphere) String() string {
	str := fmt.Sprintf(`
sphere {
  %s
  %f
  %f
}`, b.location.String(), b.radius, b.strength)
	return utils.RemoveEmptyLines(str)
}

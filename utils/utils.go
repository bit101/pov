// Package utils ...
package utils

import (
	"fmt"
	"regexp"

	"github.com/bit101/bitlib/blcolor"
	"github.com/bit101/bitlib/random"
)

// RemoveEmptyLines removes blank lines
func RemoveEmptyLines(str string) string {
	re := regexp.MustCompile(`\n\s*\n`)
	return re.ReplaceAllString(str, "\n")
}

// RGB ...
func RGB(r, g, b float64) string {
	return fmt.Sprintf("rgb <%f, %f, %f>", r, g, b)
}

// RandomRGB ...
func RandomRGB() string {
	r := random.Float()
	g := random.Float()
	b := random.Float()
	return RGB(r, g, b)
}

// Gray ...
func Gray(g float64) string {
	return RGB(g, g, g)
}

// RandomGray ...
func RandomGray() string {
	return Gray(random.Float())
}

// HSV ...
func HSV(h, s, v float64) string {
	col := blcolor.HSV(h, s, v)
	return RGB(col.R, col.G, col.B)
}

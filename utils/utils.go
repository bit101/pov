// Package utils ...
package utils

import "regexp"

// RemoveEmptyLines removes blank lines
func RemoveEmptyLines(str string) string {
	re := regexp.MustCompile(`\n\s*\n`)
	return re.ReplaceAllString(str, "\n")
}

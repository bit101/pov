// Package utils ...
package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

// ConvertToVideo converts a folder of pngs into an mp4 video file. Requires ffmpeg.
func ConvertToVideo(folder, outFileName string, w, h float64, fps int) {
	fmt.Println("Converting to video...")
	os.RemoveAll(outFileName)
	path := folder + "/main%03d.png"
	fpsArg := fmt.Sprintf("%d", fps)
	sizeArg := fmt.Sprintf("%dx%d", int(w), int(h))

	cmd := exec.Command("ffmpeg", "-framerate", fpsArg, "-i", path, "-s:v", sizeArg,
		"-c:v", "libx264", "-profile:v", "high", "-crf", "20",
		"-pix_fmt", "yuv420p", outFileName)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Video complete!")
}

// Ftos ...
func Ftos(x float64) string {
	return fmt.Sprintf("%f", x)
}

// Package pov ...
package pov

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

// Scene ...
type Scene struct {
	languageDirectives []string
	Camera             *Camera
	renderables        []Renderable
	atmosphericEffects []AtmosphericEffect
	globalSettings     []string
	width, height      int
	ambient            Vector3
}

// Renderable ...
type Renderable interface {
	String() string
}

// AtmosphericEffect ...
type AtmosphericEffect string

var currentTime time.Time

// NewScene ...
func NewScene() *Scene {
	fmt.Print("building scene...  ")
	currentTime = time.Now()
	scene := &Scene{}
	scene.Camera = NewCamera(2, 2, -5)
	scene.SetAmbient(1, 1, 1)
	scene.SetSize(640, 480)
	scene.AddInclude("stdinc.inc")
	scene.AddInclude("textures.inc")
	return scene
}

// AddLanguageDirective ...
func (s *Scene) AddLanguageDirective(ld string) {
	s.languageDirectives = append(s.languageDirectives, ld)
}

// AddInclude ...
func (s *Scene) AddInclude(include string) {
	s.AddLanguageDirective(fmt.Sprintf("#include %q", include))
}

// SetSize ...
func (s *Scene) SetSize(w, h int) {
	s.width = w
	s.height = h
}

// SetAmbient ...
func (s *Scene) SetAmbient(r, g, b float64) {
	s.ambient = *&Vector3{r, g, b}
}

// Add ...
func (s *Scene) Add(obj Renderable) {
	s.renderables = append(s.renderables, obj)
}

// AddAtmosphericEffect ...
func (s *Scene) AddAtmosphericEffect(effect AtmosphericEffect) {
	s.atmosphericEffects = append(s.atmosphericEffects, effect)
}

// AddGlobalSetting ...
func (s *Scene) AddGlobalSetting(setting string) {
	s.globalSettings = append(s.globalSettings, setting)
}

// Render ...
func (s *Scene) Render(filename string) {
	now := time.Now()
	elapsed := now.Sub(currentTime)
	currentTime = now
	fmt.Printf("%v\n", elapsed)
	fmt.Print("creating pov file...  ")
	file, err := os.Create(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	for _, directive := range s.languageDirectives {
		file.WriteString(string(directive) + "\n")
	}
	str := fmt.Sprintf("global_settings { ambient_light rgb %s }\n", s.ambient.String())
	file.WriteString(str)
	file.WriteString(s.Camera.String() + "\n")
	for _, obj := range s.renderables {
		file.WriteString(obj.String() + "\n")
	}
	for _, setting := range s.globalSettings {
		file.WriteString(string(setting) + "\n")
	}

	w := fmt.Sprintf("+W%d", s.width)
	h := fmt.Sprintf("+H%d", s.height)
	cmd := exec.Command("povray", filename, "display=false", w, h)
	now = time.Now()
	elapsed = now.Sub(currentTime)
	currentTime = now
	fmt.Printf("%v\n", elapsed)
	fmt.Print("rendering image...  ")
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	now = time.Now()
	elapsed = now.Sub(currentTime)
	currentTime = now
	fmt.Printf("%v\n", elapsed)
}

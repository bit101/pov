// Package pov ...
package pov

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Scene ...
type Scene struct {
	languageDirectives []string
	Camera             *Camera
	lights             []*LightSource
	objects            []Object
	atmosphericEffects []AtmosphericEffect
	globalSettings     []string
	width, height      int
}

// Object ...
type Object interface {
	String() string
}

// AtmosphericEffect ...
type AtmosphericEffect string

// NewScene ...
func NewScene() *Scene {
	scene := &Scene{}
	camera := NewCamera(2, 2, -5)
	camera.LookAt(0, 0, 0)
	scene.Camera = camera
	scene.width = 640
	scene.height = 480
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

// AddLight ...
func (s *Scene) AddLight(light *LightSource) {
	s.lights = append(s.lights, light)
}

// AddObject ...
func (s *Scene) AddObject(object Object) {
	s.objects = append(s.objects, object)
}

// AddAtmosphericEffect ...
func (s *Scene) AddAtmosphericEffect(effect AtmosphericEffect) {
	s.atmosphericEffects = append(s.atmosphericEffects, effect)
}

// AddGlobalSetting ...
func (s *Scene) AddGlobalSetting(setting string) {
	s.globalSettings = append(s.globalSettings, setting)
}

// WriteScene ...
func (s *Scene) Render(filename string) {
	str := ""
	for _, directive := range s.languageDirectives {
		str += string(directive) + "\n"
	}
	str += s.Camera.String() + "\n"
	for _, light := range s.lights {
		str += light.String() + "\n"
	}
	for _, object := range s.objects {
		str += object.String() + "\n"
	}
	for _, effect := range s.atmosphericEffects {
		str += string(effect) + "\n"
	}
	for _, setting := range s.globalSettings {
		str += string(setting) + "\n"
	}

	os.WriteFile(filename, []byte(str), 0644)
	w := fmt.Sprintf("+W%d", s.width)
	h := fmt.Sprintf("+H%d", s.height)
	cmd := exec.Command("povray", filename, "display=false", w, h)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

}

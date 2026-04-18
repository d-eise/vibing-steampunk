package steampunk

import (
	"strings"
	"testing"
)

func TestBoilerFrames(t *testing.T) {
	frames := BoilerFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty boiler frames")
	}
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestBoilerScene(t *testing.T) {
	s := BoilerScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestBoilerSceneFramesNonEmpty(t *testing.T) {
	s := BoilerScene()
	if len(s.Frames) == 0 {
		t.Fatal("expected scene to have frames")
	}
}

func TestBoilerSceneWithThemeDefault(t *testing.T) {
	s := BoilerSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
	if len(s.Frames) == 0 {
		t.Fatal("expected frames")
	}
}

func TestBoilerSceneWithThemeMono(t *testing.T) {
	s := BoilerSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestBoilerSceneDefaultFPS(t *testing.T) {
	s := BoilerScene()
	if s.FPS <= 0 {
		t.Errorf("expected positive FPS, got %d", s.FPS)
	}
}

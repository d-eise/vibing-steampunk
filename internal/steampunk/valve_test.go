package steampunk

import (
	"strings"
	"testing"
)

func TestValveFrames(t *testing.T) {
	frames := ValveFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty ValveFrames")
	}
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestValveScene(t *testing.T) {
	s := ValveScene()
	if s == nil {
		t.Fatal("ValveScene returned nil")
	}
}

func TestValveSceneFramesNonEmpty(t *testing.T) {
	s := ValveScene()
	if len(s.Frames) == 0 {
		t.Fatal("ValveScene has no frames")
	}
}

func TestValveSceneWithThemeDefault(t *testing.T) {
	s := ValveSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("ValveSceneWithTheme(DefaultTheme) returned nil")
	}
	if len(s.Frames) != len(ValveFrames()) {
		t.Errorf("expected %d frames, got %d", len(ValveFrames()), len(s.Frames))
	}
}

func TestValveSceneWithThemeMono(t *testing.T) {
	s := ValveSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("ValveSceneWithTheme(MonoTheme) returned nil")
	}
}

func TestValveSceneDefaultFPS(t *testing.T) {
	s := ValveScene()
	if s.FPS != 6 {
		t.Errorf("expected FPS 6, got %d", s.FPS)
	}
}

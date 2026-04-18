package steampunk

import (
	"strings"
	"testing"
)

func TestClockFrames(t *testing.T) {
	frames := ClockFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty ClockFrames")
	}
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestClockScene(t *testing.T) {
	s := ClockScene()
	if s == nil {
		t.Fatal("ClockScene returned nil")
	}
}

func TestClockSceneFramesNonEmpty(t *testing.T) {
	s := ClockScene()
	if len(s.Frames) == 0 {
		t.Fatal("ClockScene has no frames")
	}
}

func TestClockSceneWithThemeDefault(t *testing.T) {
	s := ClockSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("ClockSceneWithTheme(DefaultTheme) returned nil")
	}
	if len(s.Frames) == 0 {
		t.Fatal("expected frames")
	}
}

func TestClockSceneWithThemeMono(t *testing.T) {
	s := ClockSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("ClockSceneWithTheme(MonoTheme) returned nil")
	}
}

func TestClockSceneDefaultFPS(t *testing.T) {
	s := ClockScene()
	if s.FPS <= 0 {
		t.Errorf("expected positive FPS, got %d", s.FPS)
	}
}

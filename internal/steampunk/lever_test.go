package steampunk

import (
	"strings"
	"testing"
)

func TestLeverFrames(t *testing.T) {
	frames := LeverFrames()
	if len(frames) == 0 {
		t.Fatal("LeverFrames returned no frames")
	}
}

func TestLeverFramesCount(t *testing.T) {
	frames := LeverFrames()
	if len(frames) != 4 {
		t.Errorf("expected 4 lever frames, got %d", len(frames))
	}
}

func TestLeverFramesNonEmpty(t *testing.T) {
	for i, f := range LeverFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestLeverScene(t *testing.T) {
	s := LeverScene()
	if s == nil {
		t.Fatal("LeverScene returned nil")
	}
}

func TestLeverSceneFramesNonEmpty(t *testing.T) {
	s := LeverScene()
	for i := range LeverFrames() {
		f := s.Frame(i)
		if strings.TrimSpace(f) == "" {
			t.Errorf("scene frame %d is empty", i)
		}
	}
}

func TestLeverSceneWithThemeDefault(t *testing.T) {
	s := LeverSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("LeverSceneWithTheme(DefaultTheme) returned nil")
	}
}

func TestLeverSceneWithThemeMono(t *testing.T) {
	s := LeverSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("LeverSceneWithTheme(MonoTheme) returned nil")
	}
}

func TestLeverSceneDefaultFPS(t *testing.T) {
	s := LeverScene()
	if s.FPS() != 6 {
		t.Errorf("expected FPS 6, got %d", s.FPS())
	}
}

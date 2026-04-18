package steampunk

import (
	"strings"
	"testing"
)

func TestPendulumFrames(t *testing.T) {
	frames := PendulumFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty PendulumFrames")
	}
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestPendulumScene(t *testing.T) {
	s := PendulumScene()
	if s == nil {
		t.Fatal("PendulumScene returned nil")
	}
}

func TestPendulumSceneFramesNonEmpty(t *testing.T) {
	s := PendulumScene()
	if len(s.Frames) == 0 {
		t.Fatal("PendulumScene has no frames")
	}
}

func TestPendulumSceneWithThemeDefault(t *testing.T) {
	s := PendulumSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("PendulumSceneWithTheme(DefaultTheme) returned nil")
	}
}

func TestPendulumSceneWithThemeMono(t *testing.T) {
	s := PendulumSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("PendulumSceneWithTheme(MonoTheme) returned nil")
	}
}

func TestPendulumSceneDefaultFPS(t *testing.T) {
	s := PendulumScene()
	if s.FPS <= 0 {
		t.Errorf("expected positive FPS, got %d", s.FPS)
	}
}

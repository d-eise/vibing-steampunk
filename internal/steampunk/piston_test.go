package steampunk

import (
	"testing"
)

func TestPistonScene(t *testing.T) {
	s := PistonScene(12)
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
	if s.FPS() != 12 {
		t.Errorf("expected fps 12, got %d", s.FPS())
	}
	if s.Len() == 0 {
		t.Error("expected at least one frame")
	}
}

func TestPistonSceneFramesNonEmpty(t *testing.T) {
	s := PistonScene(8)
	for i := 0; i < s.Len(); i++ {
		f := s.Frame(i)
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestPistonSceneWithThemeDefault(t *testing.T) {
	s := PistonSceneWithTheme(10, DefaultTheme())
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
	if s.Len() == 0 {
		t.Error("expected frames")
	}
}

func TestPistonSceneWithThemeMono(t *testing.T) {
	s := PistonSceneWithTheme(10, MonoTheme())
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
	if s.Len() == 0 {
		t.Error("expected frames")
	}
}

// TestPistonSceneDefaultFPS checks that passing 0 fps doesn't produce a negative value.
// Note: ideally PistonScene(0) should fall back to a sensible default (e.g. 24).
func TestPistonSceneDefaultFPS(t *testing.T) {
	s := PistonScene(0)
	if s.FPS() < 0 {
		t.Error("fps should not be negative")
	}
}

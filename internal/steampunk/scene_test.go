package steampunk

import (
	"strings"
	"testing"
)

func TestNewScene(t *testing.T) {
	frames := []string{"a", "b", "c"}
	s := NewScene("test", frames, 10)
	if s.Name != "test" {
		t.Errorf("expected name 'test', got %q", s.Name)
	}
	if s.FPS != 10 {
		t.Errorf("expected fps 10, got %d", s.FPS)
	}
	if s.Len() != 3 {
		t.Errorf("expected 3 frames, got %d", s.Len())
	}
}

func TestSceneDefaultFPS(t *testing.T) {
	s := NewScene("x", []string{"f"}, 0)
	if s.FPS != 8 {
		t.Errorf("expected default fps 8, got %d", s.FPS)
	}
}

func TestSceneFrame(t *testing.T) {
	frames := []string{"a", "b", "c"}
	s := NewScene("wrap", frames, 8)
	if s.Frame(0) != "a" {
		t.Errorf("expected 'a', got %q", s.Frame(0))
	}
	if s.Frame(3) != "a" {
		t.Errorf("expected wrap to 'a', got %q", s.Frame(3))
	}
}

func TestSceneString(t *testing.T) {
	s := NewScene("gear", GearFrames(), 8)
	str := s.String()
	if !strings.Contains(str, "gear") {
		t.Errorf("expected scene string to contain 'gear', got %q", str)
	}
}

func TestDefaultScenes(t *testing.T) {
	scenes := DefaultScenes()
	if len(scenes) == 0 {
		t.Error("expected at least one default scene")
	}
	for _, sc := range scenes {
		if sc.Len() == 0 {
			t.Errorf("scene %q has no frames", sc.Name)
		}
	}
}

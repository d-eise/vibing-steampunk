package steampunk

import (
	"strings"
	"testing"
)

func TestSmokeFrames(t *testing.T) {
	frames := SmokeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty smoke frames")
	}
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestSmokeScene(t *testing.T) {
	s := SmokeScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
	if s.Len() != len(SmokeFrames()) {
		t.Errorf("expected %d frames, got %d", len(SmokeFrames()), s.Len())
	}
}

func TestSmokeSceneWithThemeDefault(t *testing.T) {
	s := SmokeSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestSmokeSceneWithThemeMono(t *testing.T) {
	s := SmokeSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestSmokePlumeFrames(t *testing.T) {
	frames := SmokePlumeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty plume frames")
	}
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("plume frame %d is blank", i)
		}
	}
}

func TestNewSmokePlumeScene(t *testing.T) {
	s := NewSmokePlumeScene()
	if s == nil {
		t.Fatal("expected non-nil plume scene")
	}
	if s.Len() != len(SmokePlumeFrames()) {
		t.Errorf("expected %d frames, got %d", len(SmokePlumeFrames()), s.Len())
	}
}

// TestSmokePlumeSceneFrameCount verifies the plume scene has more than one frame,
// ensuring animation is actually possible (not a static single-frame scene).
func TestSmokePlumeSceneFrameCount(t *testing.T) {
	s := NewSmokePlumeScene()
	if s == nil {
		t.Fatal("expected non-nil plume scene")
	}
	if s.Len() < 2 {
		t.Errorf("expected at least 2 frames for animation, got %d", s.Len())
	}
}

package steampunk

import (
	"strings"
	"testing"
)

func TestFlywheelFrames(t *testing.T) {
	frames := FlywheelFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty flywheel frames")
	}
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestFlywheelFramesCount(t *testing.T) {
	frames := FlywheelFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestFlywheelScene(t *testing.T) {
	s := FlywheelScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestFlywheelSceneFramesNonEmpty(t *testing.T) {
	s := FlywheelScene()
	if s.String() == "" {
		t.Error("expected non-empty scene string")
	}
}

func TestFlywheelSceneWithThemeDefault(t *testing.T) {
	s := FlywheelSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("expected non-nil scene with default theme")
	}
}

func TestFlywheelSceneWithThemeMono(t *testing.T) {
	s := FlywheelSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("expected non-nil scene with mono theme")
	}
}

func TestFlywheelSceneDefaultFPS(t *testing.T) {
	s := FlywheelScene()
	if s.FPS() <= 0 {
		t.Errorf("expected positive FPS, got %d", s.FPS())
	}
}

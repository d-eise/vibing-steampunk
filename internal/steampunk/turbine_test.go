package steampunk

import (
	"strings"
	"testing"
)

func TestTurbineFrames(t *testing.T) {
	frames := TurbineFrames()
	if len(frames) == 0 {
		t.Fatal("TurbineFrames returned no frames")
	}
}

func TestTurbineFramesCount(t *testing.T) {
	frames := TurbineFrames()
	if len(frames) != 4 {
		t.Errorf("expected 4 turbine frames, got %d", len(frames))
	}
}

func TestTurbineFramesNonEmpty(t *testing.T) {
	for i, f := range TurbineFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestTurbineScene(t *testing.T) {
	s := TurbineScene()
	if s == nil {
		t.Fatal("TurbineScene returned nil")
	}
}

func TestTurbineSceneFramesNonEmpty(t *testing.T) {
	s := TurbineScene()
	if !sceneHasFrames(s) {
		t.Error("TurbineScene has no frames")
	}
}

func TestTurbineSceneWithThemeDefault(t *testing.T) {
	s := TurbineSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("TurbineSceneWithTheme(DefaultTheme) returned nil")
	}
	if !sceneHasFrames(s) {
		t.Error("TurbineSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestTurbineSceneWithThemeMono(t *testing.T) {
	s := TurbineSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("TurbineSceneWithTheme(MonoTheme) returned nil")
	}
}

func TestTurbineSceneDefaultFPS(t *testing.T) {
	s := TurbineScene()
	if !sceneHasFPS(s, 8) {
		t.Errorf("expected TurbineScene FPS=8")
	}
}

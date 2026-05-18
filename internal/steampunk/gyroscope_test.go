package steampunk

import (
	"strings"
	"testing"
)

func TestGyroscopeFrames(t *testing.T) {
	frames := GyroscopeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty frames")
	}
}

func TestGyroscopeFramesCount(t *testing.T) {
	frames := GyroscopeFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestGyroscopeFramesNonEmpty(t *testing.T) {
	for i, f := range GyroscopeFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestGyroscopeScene(t *testing.T) {
	s := GyroscopeScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestGyroscopeSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(GyroscopeScene()) {
		t.Error("scene has no frames")
	}
}

func TestGyroscopeSceneWithThemeDefault(t *testing.T) {
	s := GyroscopeSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("scene has no frames with default theme")
	}
}

func TestGyroscopeSceneWithThemeMono(t *testing.T) {
	s := GyroscopeSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("scene has no frames with mono theme")
	}
}

func TestGyroscopeSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(GyroscopeScene(), 6) {
		t.Error("unexpected FPS for gyroscope scene")
	}
}

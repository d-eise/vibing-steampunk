package steampunk_test

import (
	"strings"
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestPressureFrames(t *testing.T) {
	frames := steampunk.PressureFrames()
	if len(frames) == 0 {
		t.Fatal("PressureFrames returned no frames")
	}
}

func TestPressureFramesCount(t *testing.T) {
	frames := steampunk.PressureFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestPressureFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.PressureFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty or whitespace", i)
		}
	}
}

func TestPressureScene(t *testing.T) {
	s := steampunk.PressureScene()
	if s == nil {
		t.Fatal("PressureScene returned nil")
	}
}

func TestPressureSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.PressureScene()) {
		t.Error("PressureScene has no frames")
	}
}

func TestPressureSceneWithThemeDefault(t *testing.T) {
	s := steampunk.PressureSceneWithTheme(steampunk.DefaultTheme)
	if !sceneHasFrames(s) {
		t.Error("PressureSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestPressureSceneWithThemeMono(t *testing.T) {
	s := steampunk.PressureSceneWithTheme(steampunk.MonoTheme)
	if !sceneHasFrames(s) {
		t.Error("PressureSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestPressureSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.PressureScene(), 6) {
		t.Error("PressureScene does not have expected FPS of 6")
	}
}

package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestCompassFrames(t *testing.T) {
	frames := steampunk.CompassFrames()
	if len(frames) == 0 {
		t.Fatal("CompassFrames returned no frames")
	}
}

func TestCompassFramesCount(t *testing.T) {
	frames := steampunk.CompassFrames()
	if len(frames) != 4 {
		t.Errorf("expected 4 frames, got %d", len(frames))
	}
}

func TestCompassFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.CompassFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestCompassScene(t *testing.T) {
	s := steampunk.CompassScene()
	if s == nil {
		t.Fatal("CompassScene returned nil")
	}
}

func TestCompassSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.CompassScene()) {
		t.Error("CompassScene has no frames")
	}
}

func TestCompassSceneWithThemeDefault(t *testing.T) {
	s := steampunk.CompassSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("CompassSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestCompassSceneWithThemeMono(t *testing.T) {
	s := steampunk.CompassSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("CompassSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestCompassSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.CompassScene(), 4) {
		t.Error("CompassScene does not have expected FPS of 4")
	}
}

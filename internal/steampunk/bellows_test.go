package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestBellowsFrames(t *testing.T) {
	frames := steampunk.BellowsFrames()
	if len(frames) == 0 {
		t.Fatal("BellowsFrames returned no frames")
	}
}

func TestBellowsFramesCount(t *testing.T) {
	frames := steampunk.BellowsFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestBellowsFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.BellowsFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestBellowsScene(t *testing.T) {
	s := steampunk.BellowsScene()
	if s == nil {
		t.Fatal("BellowsScene returned nil")
	}
}

func TestBellowsSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.BellowsScene()) {
		t.Error("BellowsScene has no frames")
	}
}

func TestBellowsSceneWithThemeDefault(t *testing.T) {
	s := steampunk.BellowsSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("BellowsSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestBellowsSceneWithThemeMono(t *testing.T) {
	s := steampunk.BellowsSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("BellowsSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestBellowsSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.BellowsScene(), 4) {
		t.Error("BellowsScene does not have expected FPS of 4")
	}
}

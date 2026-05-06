package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestOrreryFrames(t *testing.T) {
	frames := steampunk.OrreryFrames()
	if len(frames) == 0 {
		t.Fatal("OrreryFrames returned no frames")
	}
}

func TestOrreryFramesCount(t *testing.T) {
	frames := steampunk.OrreryFrames()
	if len(frames) < 4 {
		t.Errorf("expected at least 4 frames, got %d", len(frames))
	}
}

func TestOrreryFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.OrreryFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestOrreryScene(t *testing.T) {
	s := steampunk.OrreryScene()
	if s == nil {
		t.Fatal("OrreryScene returned nil")
	}
}

func TestOrrerySceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.OrreryScene()) {
		t.Error("OrreryScene has no frames")
	}
}

func TestOrrerySceneWithThemeDefault(t *testing.T) {
	s := steampunk.OrrerySceneWithTheme(steampunk.DefaultTheme)
	if s == nil {
		t.Fatal("OrrerySceneWithTheme(DefaultTheme) returned nil")
	}
}

func TestOrrerySceneWithThemeMono(t *testing.T) {
	s := steampunk.OrrerySceneWithTheme(steampunk.MonoTheme)
	if s == nil {
		t.Fatal("OrrerySceneWithTheme(MonoTheme) returned nil")
	}
}

func TestOrrerySceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.OrreryScene(), 6) {
		t.Error("OrreryScene does not have expected FPS of 6")
	}
}

package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestCapacitorFrames(t *testing.T) {
	frames := steampunk.CapacitorFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty CapacitorFrames")
	}
}

func TestCapacitorFramesCount(t *testing.T) {
	frames := steampunk.CapacitorFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestCapacitorFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.CapacitorFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestCapacitorScene(t *testing.T) {
	s := steampunk.CapacitorScene()
	if s == nil {
		t.Fatal("expected non-nil CapacitorScene")
	}
}

func TestCapacitorSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.CapacitorScene()) {
		t.Fatal("CapacitorScene has no frames")
	}
}

func TestCapacitorSceneWithThemeDefault(t *testing.T) {
	s := steampunk.CapacitorSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("CapacitorSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestCapacitorSceneWithThemeMono(t *testing.T) {
	s := steampunk.CapacitorSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("CapacitorSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestCapacitorSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.CapacitorScene(), 6) {
		t.Fatal("CapacitorScene does not have expected FPS of 6")
	}
}

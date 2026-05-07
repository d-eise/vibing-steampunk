package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestAutomatonFrames(t *testing.T) {
	frames := steampunk.AutomatonFrames()
	if len(frames) == 0 {
		t.Fatal("AutomatonFrames returned no frames")
	}
}

func TestAutomatonFramesCount(t *testing.T) {
	frames := steampunk.AutomatonFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestAutomatonFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.AutomatonFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestAutomatonScene(t *testing.T) {
	s := steampunk.AutomatonScene()
	if s == nil {
		t.Fatal("AutomatonScene returned nil")
	}
}

func TestAutomatonSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.AutomatonScene()) {
		t.Fatal("AutomatonScene has no frames")
	}
}

func TestAutomatonSceneWithThemeDefault(t *testing.T) {
	s := steampunk.AutomatonSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("AutomatonSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestAutomatonSceneWithThemeMono(t *testing.T) {
	s := steampunk.AutomatonSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("AutomatonSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestAutomatonSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.AutomatonScene(), 4) {
		t.Fatal("AutomatonScene does not have expected FPS")
	}
}

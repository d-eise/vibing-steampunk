package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestCondenserFrames(t *testing.T) {
	frames := steampunk.CondenserFrames()
	if len(frames) == 0 {
		t.Fatal("CondenserFrames returned no frames")
	}
}

func TestCondenserFramesCount(t *testing.T) {
	frames := steampunk.CondenserFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestCondenserFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.CondenserFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestCondenserScene(t *testing.T) {
	s := steampunk.CondenserScene()
	if s == nil {
		t.Fatal("CondenserScene returned nil")
	}
}

func TestCondenserSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.CondenserScene()) {
		t.Error("CondenserScene has no frames")
	}
}

func TestCondenserSceneWithThemeDefault(t *testing.T) {
	s := steampunk.CondenserSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("CondenserSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestCondenserSceneWithThemeMono(t *testing.T) {
	s := steampunk.CondenserSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("CondenserSceneWithTheme(MonoTheme) has no frames")
	}
}

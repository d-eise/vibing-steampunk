package steampunk_test

import (
	"strings"
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestThermometerFrames(t *testing.T) {
	frames := steampunk.ThermometerFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty thermometer frames")
	}
}

func TestThermometerFramesCount(t *testing.T) {
	frames := steampunk.ThermometerFrames()
	if len(frames) != 5 {
		t.Errorf("expected 5 frames, got %d", len(frames))
	}
}

func TestThermometerFramesNonEmpty(t *testing.T) {
	frames := steampunk.ThermometerFrames()
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestThermometerScene(t *testing.T) {
	s := steampunk.ThermometerScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestThermometerSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.ThermometerScene()) {
		t.Error("expected thermometer scene to have frames")
	}
}

func TestThermometerSceneWithThemeDefault(t *testing.T) {
	s := steampunk.ThermometerSceneWithTheme(steampunk.DefaultTheme)
	if !sceneHasFrames(s) {
		t.Error("expected thermometer scene with default theme to have frames")
	}
}

func TestThermometerSceneWithThemeMono(t *testing.T) {
	s := steampunk.ThermometerSceneWithTheme(steampunk.MonoTheme)
	if !sceneHasFrames(s) {
		t.Error("expected thermometer scene with mono theme to have frames")
	}
}

func TestThermometerSceneDefaultFPS(t *testing.T) {
	s := steampunk.ThermometerScene()
	if !sceneHasFPS(s, 4) {
		t.Errorf("expected thermometer scene fps=4")
	}
}

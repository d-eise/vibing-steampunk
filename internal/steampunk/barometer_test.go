package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestBarometerFrames(t *testing.T) {
	frames := steampunk.BarometerFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty barometer frames")
	}
}

func TestBarometerFramesCount(t *testing.T) {
	frames := steampunk.BarometerFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestBarometerFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.BarometerFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestBarometerScene(t *testing.T) {
	s := steampunk.BarometerScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestBarometerSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.BarometerScene()) {
		t.Fatal("expected scene to have frames")
	}
}

func TestBarometerSceneWithThemeDefault(t *testing.T) {
	s := steampunk.BarometerSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with default theme to have frames")
	}
}

func TestBarometerSceneWithThemeMono(t *testing.T) {
	s := steampunk.BarometerSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with mono theme to have frames")
	}
}

func TestBarometerSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.BarometerScene(), 4) {
		t.Fatal("expected barometer scene to have fps=4")
	}
}

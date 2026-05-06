package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestAnemometerFrames(t *testing.T) {
	frames := steampunk.AnemometerFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty frames")
	}
}

func TestAnemometerFramesCount(t *testing.T) {
	frames := steampunk.AnemometerFrames()
	if len(frames) != 8 {
		t.Fatalf("expected 8 frames, got %d", len(frames))
	}
}

func TestAnemometerFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.AnemometerFrames() {
		if f == "" {
			t.Fatalf("frame %d is empty", i)
		}
	}
}

func TestAnemometerScene(t *testing.T) {
	s := steampunk.AnemometerScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestAnemometerSceneFramesNonEmpty(t *testing.T) {
	s := steampunk.AnemometerScene()
	if !sceneHasFrames(s) {
		t.Fatal("expected scene to have frames")
	}
}

func TestAnemometerSceneWithThemeDefault(t *testing.T) {
	s := steampunk.AnemometerSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with default theme to have frames")
	}
}

func TestAnemometerSceneWithThemeMono(t *testing.T) {
	s := steampunk.AnemometerSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with mono theme to have frames")
	}
}

func TestAnemometerSceneDefaultFPS(t *testing.T) {
	s := steampunk.AnemometerScene()
	if !sceneHasFPS(s, 8) {
		t.Fatalf("expected FPS 8")
	}
}

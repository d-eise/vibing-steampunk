package steampunk

import (
	"testing"
)

func TestGalvanometerFrames(t *testing.T) {
	frames := GalvanometerFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty GalvanometerFrames")
	}
}

func TestGalvanometerFramesCount(t *testing.T) {
	frames := GalvanometerFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestGalvanometerFramesNonEmpty(t *testing.T) {
	for i, f := range GalvanometerFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestGalvanometerScene(t *testing.T) {
	s := GalvanometerScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestGalvanometerSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(GalvanometerScene()) {
		t.Fatal("expected scene to have frames")
	}
}

func TestGalvanometerSceneWithThemeDefault(t *testing.T) {
	s := GalvanometerSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with default theme to have frames")
	}
}

func TestGalvanometerSceneWithThemeMono(t *testing.T) {
	s := GalvanometerSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with mono theme to have frames")
	}
}

func TestGalvanometerSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(GalvanometerScene(), 4) {
		t.Fatal("expected scene FPS to be 4")
	}
}

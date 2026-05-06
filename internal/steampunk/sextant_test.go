package steampunk

import (
	"strings"
	"testing"
)

func TestSextantFrames(t *testing.T) {
	frames := SextantFrames()
	if len(frames) == 0 {
		t.Fatal("SextantFrames returned no frames")
	}
}

func TestSextantFramesCount(t *testing.T) {
	frames := SextantFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestSextantFramesNonEmpty(t *testing.T) {
	for i, f := range SextantFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestSextantScene(t *testing.T) {
	s := SextantScene()
	if s == nil {
		t.Fatal("SextantScene returned nil")
	}
}

func TestSextantSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(SextantScene()) {
		t.Error("SextantScene has no frames")
	}
}

func TestSextantSceneWithThemeDefault(t *testing.T) {
	s := SextantSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("SextantSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestSextantSceneWithThemeMono(t *testing.T) {
	s := SextantSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("SextantSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestSextantSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(SextantScene(), 6) {
		t.Error("SextantScene does not have expected FPS of 6")
	}
}

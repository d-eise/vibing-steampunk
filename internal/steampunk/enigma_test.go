package steampunk

import (
	"strings"
	"testing"
)

func TestEnigmaFrames(t *testing.T) {
	frames := EnigmaFrames()
	if len(frames) == 0 {
		t.Fatal("EnigmaFrames returned no frames")
	}
}

func TestEnigmaFramesCount(t *testing.T) {
	frames := EnigmaFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestEnigmaFramesNonEmpty(t *testing.T) {
	for i, f := range EnigmaFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestEnigmaScene(t *testing.T) {
	s := EnigmaScene()
	if s == nil {
		t.Fatal("EnigmaScene returned nil")
	}
}

func TestEnigmaSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(EnigmaScene()) {
		t.Error("EnigmaScene has no frames")
	}
}

func TestEnigmaSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(EnigmaScene(), 6) {
		t.Errorf("expected FPS 6, got %d", EnigmaScene().FPS)
	}
}

func TestEnigmaSceneWithThemeDefault(t *testing.T) {
	s := EnigmaSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("EnigmaSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestEnigmaSceneWithThemeMono(t *testing.T) {
	s := EnigmaSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("EnigmaSceneWithTheme(MonoTheme) has no frames")
	}
}

package steampunk

import (
	"strings"
	"testing"
)

func TestEscapementFrames(t *testing.T) {
	frames := EscapementFrames()
	if len(frames) == 0 {
		t.Fatal("EscapementFrames returned no frames")
	}
}

func TestEscapementFramesCount(t *testing.T) {
	frames := EscapementFrames()
	if len(frames) != 4 {
		t.Errorf("expected 4 frames, got %d", len(frames))
	}
}

func TestEscapementFramesNonEmpty(t *testing.T) {
	for i, f := range EscapementFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestEscapementScene(t *testing.T) {
	s := EscapementScene()
	if s == nil {
		t.Fatal("EscapementScene returned nil")
	}
}

func TestEscapementSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(EscapementScene()) {
		t.Error("EscapementScene has no frames")
	}
}

func TestEscapementSceneWithThemeDefault(t *testing.T) {
	s := EscapementSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("EscapementSceneWithTheme(DefaultTheme) returned nil")
	}
	if !sceneHasFrames(s) {
		t.Error("EscapementSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestEscapementSceneWithThemeMono(t *testing.T) {
	s := EscapementSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("EscapementSceneWithTheme(MonoTheme) returned nil")
	}
	if !sceneHasFrames(s) {
		t.Error("EscapementSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestEscapementSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(EscapementScene(), 8) {
		t.Errorf("EscapementScene expected FPS 8")
	}
}

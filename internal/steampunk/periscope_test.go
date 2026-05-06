package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestPeriscopeFrames(t *testing.T) {
	frames := steampunk.PeriscopeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty PeriscopeFrames")
	}
}

func TestPeriscopeFramesCount(t *testing.T) {
	frames := steampunk.PeriscopeFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestPeriscopeFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.PeriscopeFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestPeriscopeScene(t *testing.T) {
	s := steampunk.PeriscopeScene()
	if !sceneHasFrames(s) {
		t.Fatal("PeriscopeScene has no frames")
	}
}

func TestPeriscopeSceneFramesNonEmpty(t *testing.T) {
	s := steampunk.PeriscopeScene()
	for i := 0; i < len(steampunk.PeriscopeFrames()); i++ {
		if s.Frame(i) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestPeriscopeSceneWithThemeDefault(t *testing.T) {
	s := steampunk.PeriscopeSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("PeriscopeSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestPeriscopeSceneWithThemeMono(t *testing.T) {
	s := steampunk.PeriscopeSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("PeriscopeSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestPeriscopeSceneDefaultFPS(t *testing.T) {
	s := steampunk.PeriscopeScene()
	if !sceneHasFPS(s, 4) {
		t.Fatalf("expected FPS 4, got different value")
	}
}

package steampunk_test

import (
	"strings"
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestPortholeFrames(t *testing.T) {
	frames := steampunk.PortholeFrames()
	if len(frames) == 0 {
		t.Fatal("PortholeFrames returned no frames")
	}
}

func TestPortholeFramesCount(t *testing.T) {
	frames := steampunk.PortholeFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestPortholeFramesNonEmpty(t *testing.T) {
	frames := steampunk.PortholeFrames()
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestPortholeScene(t *testing.T) {
	s := steampunk.PortholeScene()
	if !sceneHasFrames(s) {
		t.Error("PortholeScene has no frames")
	}
}

func TestPortholeSceneFramesNonEmpty(t *testing.T) {
	s := steampunk.PortholeScene()
	for i := 0; i < len(steampunk.PortholeFrames()); i++ {
		if strings.TrimSpace(s.Frame(i)) == "" {
			t.Errorf("scene frame %d is empty", i)
		}
	}
}

func TestPortholeSceneWithThemeDefault(t *testing.T) {
	s := steampunk.PortholeSceneWithTheme(steampunk.DefaultTheme)
	if !sceneHasFrames(s) {
		t.Error("PortholeSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestPortholeSceneWithThemeMono(t *testing.T) {
	s := steampunk.PortholeSceneWithTheme(steampunk.MonoTheme)
	if !sceneHasFrames(s) {
		t.Error("PortholeSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestPortholeSceneDefaultFPS(t *testing.T) {
	s := steampunk.PortholeScene()
	if !sceneHasFPS(s, 4) {
		t.Errorf("expected FPS=4")
	}
}

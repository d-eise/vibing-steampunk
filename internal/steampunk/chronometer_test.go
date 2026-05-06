package steampunk_test

import (
	"strings"
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestChronometerFrames(t *testing.T) {
	frames := steampunk.ChronometerFrames()
	if len(frames) == 0 {
		t.Fatal("ChronometerFrames returned no frames")
	}
}

func TestChronometerFramesCount(t *testing.T) {
	frames := steampunk.ChronometerFrames()
	if len(frames)%5 != 0 {
		t.Errorf("expected frame count divisible by 5, got %d", len(frames))
	}
}

func TestChronometerFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.ChronometerFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestChronometerScene(t *testing.T) {
	s := steampunk.ChronometerScene()
	if s == nil {
		t.Fatal("ChronometerScene returned nil")
	}
}

func TestChronometerSceneFramesNonEmpty(t *testing.T) {
	s := steampunk.ChronometerScene()
	if !sceneHasFrames(s) {
		t.Error("ChronometerScene has no frames")
	}
}

func TestChronometerSceneWithThemeDefault(t *testing.T) {
	s := steampunk.ChronometerSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("ChronometerSceneWithTheme(Default) has no frames")
	}
}

func TestChronometerSceneWithThemeMono(t *testing.T) {
	s := steampunk.ChronometerSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("ChronometerSceneWithTheme(Mono) has no frames")
	}
}

func TestChronometerSceneDefaultFPS(t *testing.T) {
	s := steampunk.ChronometerScene()
	if !sceneHasFPS(s, 6) {
		t.Errorf("expected FPS 6")
	}
}

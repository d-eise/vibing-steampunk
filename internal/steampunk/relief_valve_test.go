package steampunk_test

import (
	"strings"
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestReliefValveFrames(t *testing.T) {
	frames := steampunk.ReliefValveFrames()
	if len(frames) == 0 {
		t.Fatal("ReliefValveFrames returned no frames")
	}
}

func TestReliefValveFramesCount(t *testing.T) {
	frames := steampunk.ReliefValveFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestReliefValveFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.ReliefValveFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty or whitespace", i)
		}
	}
}

func TestReliefValveScene(t *testing.T) {
	s := steampunk.ReliefValveScene()
	if s == nil {
		t.Fatal("ReliefValveScene returned nil")
	}
}

func TestReliefValveSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.ReliefValveScene()) {
		t.Error("ReliefValveScene has no frames")
	}
}

func TestReliefValveSceneWithThemeDefault(t *testing.T) {
	s := steampunk.ReliefValveSceneWithTheme(steampunk.DefaultTheme)
	if !sceneHasFrames(s) {
		t.Error("ReliefValveSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestReliefValveSceneWithThemeMono(t *testing.T) {
	s := steampunk.ReliefValveSceneWithTheme(steampunk.MonoTheme)
	if !sceneHasFrames(s) {
		t.Error("ReliefValveSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestReliefValveSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.ReliefValveScene(), 8) {
		t.Error("ReliefValveScene does not have expected FPS of 8")
	}
}

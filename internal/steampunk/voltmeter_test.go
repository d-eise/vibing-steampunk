package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestVoltmeterFrames(t *testing.T) {
	frames := steampunk.VoltmeterFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty VoltmeterFrames")
	}
}

func TestVoltmeterFramesCount(t *testing.T) {
	frames := steampunk.VoltmeterFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestVoltmeterFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.VoltmeterFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestVoltmeterScene(t *testing.T) {
	s := steampunk.VoltmeterScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestVoltmeterSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.VoltmeterScene()) {
		t.Error("VoltmeterScene has no frames")
	}
}

func TestVoltmeterSceneWithThemeDefault(t *testing.T) {
	s := steampunk.VoltmeterSceneWithTheme(steampunk.DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("VoltmeterSceneWithTheme(Default) has no frames")
	}
}

func TestVoltmeterSceneWithThemeMono(t *testing.T) {
	s := steampunk.VoltmeterSceneWithTheme(steampunk.MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("VoltmeterSceneWithTheme(Mono) has no frames")
	}
}

func TestVoltmeterSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.VoltmeterScene(), 4) {
		t.Error("VoltmeterScene does not have expected FPS of 4")
	}
}

func TestBuzzFrames(t *testing.T) {
	frames := steampunk.BuzzFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty BuzzFrames")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("BuzzFrame %d is empty", i)
		}
	}
}

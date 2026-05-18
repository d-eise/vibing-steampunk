package steampunk

import (
	"testing"
)

func TestTelescopeFrames(t *testing.T) {
	frames := TelescopeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty frames")
	}
}

func TestTelescopeFramesCount(t *testing.T) {
	frames := TelescopeFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestTelescopeFramesNonEmpty(t *testing.T) {
	for i, f := range TelescopeFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestTelescopeScene(t *testing.T) {
	s := TelescopeScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestTelescopeSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(TelescopeScene()) {
		t.Fatal("expected scene to have frames")
	}
}

func TestTelescopeSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(TelescopeScene(), 6) {
		t.Fatal("expected scene to have fps=6")
	}
}

func TestTelescopeSceneWithThemeDefault(t *testing.T) {
	s := TelescopeSceneWithTheme(DefaultTheme)
	if !sceneHasFrames(s) {
		t.Fatal("expected themed scene to have frames")
	}
}

func TestTelescopeSceneWithThemeMono(t *testing.T) {
	s := TelescopeSceneWithTheme(MonoTheme)
	if !sceneHasFrames(s) {
		t.Fatal("expected mono themed scene to have frames")
	}
}

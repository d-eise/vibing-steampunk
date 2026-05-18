package steampunk

import (
	"strings"
	"testing"
)

func TestAstrolabeFrames(t *testing.T) {
	frames := AstrolabeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty AstrolabeFrames")
	}
}

func TestAstrolabeFramesCount(t *testing.T) {
	frames := AstrolabeFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestAstrolabeFramesNonEmpty(t *testing.T) {
	for i, f := range AstrolabeFrames() {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestAstrolabeScene(t *testing.T) {
	s := AstrolabeScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestAstrolabeSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(AstrolabeScene()) {
		t.Fatal("expected scene to have frames")
	}
}

func TestAstrolabeSceneWithThemeDefault(t *testing.T) {
	s := AstrolabeSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with default theme to have frames")
	}
}

func TestAstrolabeSceneWithThemeMono(t *testing.T) {
	s := AstrolabeSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with mono theme to have frames")
	}
}

func TestAstrolabeSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(AstrolabeScene(), 4) {
		t.Fatal("expected astrolabe scene to have fps=4")
	}
}

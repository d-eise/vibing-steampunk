package steampunk

import (
	"testing"
)

func TestSpeakingTubeFrames(t *testing.T) {
	frames := SpeakingTubeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty frames")
	}
}

func TestSpeakingTubeFramesCount(t *testing.T) {
	frames := SpeakingTubeFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestSpeakingTubeFramesNonEmpty(t *testing.T) {
	frames := SpeakingTubeFrames()
	for i, f := range frames {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestSpeakingTubeScene(t *testing.T) {
	s := SpeakingTubeScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestSpeakingTubeSceneFramesNonEmpty(t *testing.T) {
	s := SpeakingTubeScene()
	if !sceneHasFrames(s) {
		t.Fatal("expected scene to have frames")
	}
}

func TestSpeakingTubeSceneWithThemeDefault(t *testing.T) {
	s := SpeakingTubeSceneWithTheme(DefaultTheme)
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
	if !sceneHasFrames(s) {
		t.Fatal("expected scene to have frames")
	}
}

func TestSpeakingTubeSceneWithThemeMono(t *testing.T) {
	s := SpeakingTubeSceneWithTheme(MonoTheme)
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestSpeakingTubeSceneDefaultFPS(t *testing.T) {
	s := SpeakingTubeScene()
	if !sceneHasFPS(s, 4) {
		t.Fatalf("expected FPS 4")
	}
}

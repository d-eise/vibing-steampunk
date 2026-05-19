package steampunk

import "testing"

func TestCoilgunFrames(t *testing.T) {
	frames := CoilgunFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty CoilgunFrames")
	}
}

func TestCoilgunFramesCount(t *testing.T) {
	frames := CoilgunFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestCoilgunFramesNonEmpty(t *testing.T) {
	for i, f := range CoilgunFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestCoilgunScene(t *testing.T) {
	s := CoilgunScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestCoilgunSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(CoilgunScene()) {
		t.Fatal("expected scene to have frames")
	}
}

func TestCoilgunSceneWithThemeDefault(t *testing.T) {
	s := CoilgunSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with default theme to have frames")
	}
}

func TestCoilgunSceneWithThemeMono(t *testing.T) {
	s := CoilgunSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with mono theme to have frames")
	}
}

func TestCoilgunSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(CoilgunScene(), 6) {
		t.Fatal("expected default FPS of 6")
	}
}

func TestChargeFrames(t *testing.T) {
	frames := ChargeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty ChargeFrames")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("charge frame %d is empty", i)
		}
	}
}

func TestNewCoilgunSceneWithSound(t *testing.T) {
	s := NewCoilgunSceneWithSound(DefaultTheme())
	if s == nil {
		t.Fatal("expected non-nil scene with sound")
	}
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with sound to have frames")
	}
}

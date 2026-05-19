package steampunk

import (
	"testing"
)

func TestPneumaticFrames(t *testing.T) {
	frames := PneumaticFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty frames")
	}
}

func TestPneumaticFramesCount(t *testing.T) {
	frames := PneumaticFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestPneumaticFramesNonEmpty(t *testing.T) {
	for i, f := range PneumaticFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestPneumaticScene(t *testing.T) {
	scene := PneumaticScene()
	if scene == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestPneumaticSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(PneumaticScene()) {
		t.Fatal("expected scene to have frames")
	}
}

func TestPneumaticSceneWithThemeDefault(t *testing.T) {
	scene := PneumaticSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(scene) {
		t.Fatal("expected scene with default theme to have frames")
	}
}

func TestPneumaticSceneWithThemeMono(t *testing.T) {
	scene := PneumaticSceneWithTheme(MonoTheme())
	if !sceneHasFrames(scene) {
		t.Fatal("expected scene with mono theme to have frames")
	}
}

func TestPneumaticSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(PneumaticScene(), 6) {
		t.Fatal("expected pneumatic scene to have 6 FPS")
	}
}

func TestPneumaticWhooshFrames(t *testing.T) {
	frames := PneumaticWhooshFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty whoosh frames")
	}
}

func TestNewPneumaticSceneWithSound(t *testing.T) {
	scene, sfx := NewPneumaticSceneWithSound(DefaultTheme())
	if scene == nil {
		t.Fatal("expected non-nil scene")
	}
	if sfx == nil {
		t.Fatal("expected non-nil sound effect")
	}
}

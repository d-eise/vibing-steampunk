package steampunk

import (
	"testing"
)

func TestFocusFrames(t *testing.T) {
	frames := FocusFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty FocusFrames")
	}
}

func TestFocusFramesNonEmpty(t *testing.T) {
	for i, f := range FocusFrames() {
		if f == "" {
			t.Errorf("FocusFrames[%d] is empty", i)
		}
	}
}

func TestNewTelescopeSoundEffect(t *testing.T) {
	se := NewTelescopeSoundEffect()
	if se == nil {
		t.Fatal("expected non-nil SoundEffect")
	}
}

func TestTelescopeSoundEffectFrame(t *testing.T) {
	se := NewTelescopeSoundEffect()
	frame := se.Frame(0)
	if frame == "" {
		t.Fatal("expected non-empty sound frame at index 0")
	}
}

func TestNewTelescopeSceneWithSound(t *testing.T) {
	s := NewTelescopeSceneWithSound(DefaultTheme)
	if s == nil {
		t.Fatal("expected non-nil scene with sound")
	}
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with sound to have frames")
	}
}

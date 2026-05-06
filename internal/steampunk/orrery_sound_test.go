package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestGearClickFrames(t *testing.T) {
	frames := steampunk.GearClickFrames()
	if len(frames) == 0 {
		t.Fatal("GearClickFrames returned no frames")
	}
}

func TestGearClickFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.GearClickFrames() {
		if f == "" {
			t.Errorf("GearClickFrames frame %d is empty", i)
		}
	}
}

func TestNewOrrerySoundEffect(t *testing.T) {
	se := steampunk.NewOrrerySoundEffect()
	if se == nil {
		t.Fatal("NewOrrerySoundEffect returned nil")
	}
}

func TestOrrerySoundEffectFrame(t *testing.T) {
	se := steampunk.NewOrrerySoundEffect()
	f := se.Frame(0)
	if f == "" {
		t.Error("SoundEffect.Frame(0) returned empty string")
	}
}

func TestNewOrrerySceneWithSound(t *testing.T) {
	s := steampunk.NewOrrerySceneWithSound(steampunk.DefaultTheme)
	if s == nil {
		t.Fatal("NewOrrerySceneWithSound returned nil")
	}
	if !sceneHasFrames(s) {
		t.Error("NewOrrerySceneWithSound scene has no frames")
	}
}

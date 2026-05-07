package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestRatchetFrames(t *testing.T) {
	frames := steampunk.RatchetFrames()
	if len(frames) == 0 {
		t.Fatal("RatchetFrames returned no frames")
	}
}

func TestRatchetFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.RatchetFrames() {
		if f == "" {
			t.Errorf("RatchetFrames frame %d is empty", i)
		}
	}
}

func TestNewAutomatonSoundEffect(t *testing.T) {
	se := steampunk.NewAutomatonSoundEffect()
	if se == nil {
		t.Fatal("NewAutomatonSoundEffect returned nil")
	}
}

func TestAutomatonSoundEffectFrame(t *testing.T) {
	se := steampunk.NewAutomatonSoundEffect()
	f := se.Frame(0)
	if f == "" {
		t.Fatal("SoundEffect.Frame(0) returned empty string")
	}
}

func TestNewAutomatonSceneWithSound(t *testing.T) {
	sws := steampunk.NewAutomatonSceneWithSound(steampunk.DefaultTheme())
	if sws == nil {
		t.Fatal("NewAutomatonSceneWithSound returned nil")
	}
}

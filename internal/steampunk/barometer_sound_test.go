package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestTickTockFrames(t *testing.T) {
	frames := steampunk.TickTockFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty tick-tock frames")
	}
}

func TestTickTockFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.TickTockFrames() {
		if f == "" {
			t.Errorf("tick-tock frame %d is empty", i)
		}
	}
}

func TestNewBarometerSoundEffect(t *testing.T) {
	se := steampunk.NewBarometerSoundEffect()
	if se == nil {
		t.Fatal("expected non-nil sound effect")
	}
}

func TestBarometerSoundEffectFrame(t *testing.T) {
	se := steampunk.NewBarometerSoundEffect()
	f := se.Frame(0)
	if f == "" {
		t.Fatal("expected non-empty frame at index 0")
	}
}

func TestNewBarometerSceneWithSound(t *testing.T) {
	s := steampunk.NewBarometerSceneWithSound(steampunk.DefaultTheme())
	if s == nil {
		t.Fatal("expected non-nil scene with sound")
	}
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with sound to have frames")
	}
}

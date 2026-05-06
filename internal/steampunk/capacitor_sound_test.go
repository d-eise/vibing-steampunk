package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestZapFrames(t *testing.T) {
	frames := steampunk.ZapFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty ZapFrames")
	}
}

func TestZapFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.ZapFrames() {
		if f == "" {
			t.Errorf("ZapFrames frame %d is empty", i)
		}
	}
}

func TestNewCapacitorSoundEffect(t *testing.T) {
	se := steampunk.NewCapacitorSoundEffect()
	if se == nil {
		t.Fatal("expected non-nil SoundEffect")
	}
}

func TestCapacitorSoundEffectFrame(t *testing.T) {
	se := steampunk.NewCapacitorSoundEffect()
	f := se.Frame(0)
	if f == "" {
		t.Fatal("expected non-empty frame from CapacitorSoundEffect")
	}
}

func TestNewCapacitorSceneWithSound(t *testing.T) {
	sws := steampunk.NewCapacitorSceneWithSound()
	if sws == nil {
		t.Fatal("expected non-nil SceneWithSound")
	}
}

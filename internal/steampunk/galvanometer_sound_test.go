package steampunk

import (
	"testing"
)

func TestDeflectFrames(t *testing.T) {
	frames := DeflectFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty DeflectFrames")
	}
}

func TestDeflectFramesNonEmpty(t *testing.T) {
	for i, f := range DeflectFrames() {
		if f == "" {
			t.Errorf("deflect frame %d is empty", i)
		}
	}
}

func TestNewGalvanometerSoundEffect(t *testing.T) {
	se := NewGalvanometerSoundEffect()
	if se == nil {
		t.Fatal("expected non-nil sound effect")
	}
}

func TestGalvanometerSoundEffectFrame(t *testing.T) {
	se := NewGalvanometerSoundEffect()
	f := se.Frame(0)
	if f == "" {
		t.Fatal("expected non-empty sound frame at index 0")
	}
}

func TestNewGalvanometerSceneWithSound(t *testing.T) {
	sws := NewGalvanometerSceneWithSound()
	if sws == nil {
		t.Fatal("expected non-nil SceneWithSound")
	}
}

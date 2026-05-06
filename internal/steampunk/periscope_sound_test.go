package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestSwishFrames(t *testing.T) {
	frames := steampunk.SwishFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty SwishFrames")
	}
}

func TestSwishFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.SwishFrames() {
		if f == "" {
			t.Errorf("SwishFrames frame %d is empty", i)
		}
	}
}

func TestNewPeriscopeSoundEffect(t *testing.T) {
	se := steampunk.NewPeriscopeSoundEffect()
	if se.Len() == 0 {
		t.Fatal("expected non-zero sound effect length")
	}
}

func TestPeriscopeSoundEffectFrame(t *testing.T) {
	se := steampunk.NewPeriscopeSoundEffect()
	for i := 0; i < se.Len(); i++ {
		f := se.Frame(i)
		if f == "" {
			t.Errorf("sound effect frame %d is empty", i)
		}
	}
}

func TestNewPeriscopeSceneWithSound(t *testing.T) {
	sws := steampunk.NewPeriscopeSceneWithSound()
	if sws.Scene.Frame(0) == "" {
		t.Fatal("SceneWithSound scene frame 0 is empty")
	}
	if sws.Sound.Len() == 0 {
		t.Fatal("SceneWithSound sound has no frames")
	}
}

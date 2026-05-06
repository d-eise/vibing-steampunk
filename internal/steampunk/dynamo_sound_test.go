package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestHumFrames(t *testing.T) {
	frames := steampunk.HumFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty HumFrames")
	}
}

func TestHumFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.HumFrames() {
		if f == "" {
			t.Errorf("HumFrames: frame %d is empty", i)
		}
	}
}

func TestNewDynamoSoundEffect(t *testing.T) {
	se := steampunk.NewDynamoSoundEffect()
	if se == nil {
		t.Fatal("expected non-nil SoundEffect")
	}
}

func TestDynamoSoundEffectFrame(t *testing.T) {
	se := steampunk.NewDynamoSoundEffect()
	f := se.Frame(0)
	if f == "" {
		t.Fatal("expected non-empty frame from DynamoSoundEffect")
	}
}

func TestNewDynamoSceneWithSound(t *testing.T) {
	s := steampunk.NewDynamoSceneWithSound()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
	if !sceneHasFrames(s) {
		t.Fatal("NewDynamoSceneWithSound has no frames")
	}
}

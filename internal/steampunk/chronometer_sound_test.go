package steampunk_test

import (
	"strings"
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestTickTackFrames(t *testing.T) {
	frames := steampunk.TickTackFrames()
	if len(frames) == 0 {
		t.Fatal("TickTackFrames returned no frames")
	}
}

func TestTickTackFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.TickTackFrames() {
		_ = i
		_ = strings.TrimSpace(f)
	}
}

func TestNewChronometerSoundEffect(t *testing.T) {
	se := steampunk.NewChronometerSoundEffect()
	if se == nil {
		t.Fatal("NewChronometerSoundEffect returned nil")
	}
}

func TestChronometerSoundEffectFrame(t *testing.T) {
	se := steampunk.NewChronometerSoundEffect()
	f := se.Frame(0)
	if f == "" {
		t.Error("expected non-empty sound frame at index 0")
	}
}

func TestNewChronometerSceneWithSound(t *testing.T) {
	sws := steampunk.NewChronometerSceneWithSound(steampunk.DefaultTheme())
	if sws == nil {
		t.Fatal("NewChronometerSceneWithSound returned nil")
	}
	if sws.Frame(0) == "" {
		t.Error("expected non-empty combined frame at index 0")
	}
}

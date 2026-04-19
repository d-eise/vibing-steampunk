package steampunk

import (
	"testing"
)

func TestChuggingFrames(t *testing.T) {
	frames := ChuggingFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty ChuggingFrames")
	}
	for _, f := range frames {
		_ = f // frames are strings, just ensure non-nil
	}
}

func TestWhistleFrames(t *testing.T) {
	frames := WhistleFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty WhistleFrames")
	}
}

func TestHissFrames(t *testing.T) {
	frames := HissFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty HissFrames")
	}
}

func TestNewSoundEffect(t *testing.T) {
	se := NewSoundEffect("test", []string{"a", "b", "c"})
	if se.Label != "test" {
		t.Errorf("expected label 'test', got %q", se.Label)
	}
	if len(se.Frames) != 3 {
		t.Errorf("expected 3 frames, got %d", len(se.Frames))
	}
}

func TestSoundEffectFrame(t *testing.T) {
	se := NewSoundEffect("chug", ChuggingFrames())
	n := len(se.Frames)
	// iterate 3 full cycles to better exercise wrap-around behaviour
	for i := 0; i < n*3; i++ {
		f := se.Frame(i)
		if f == "" && se.Frames[i%n] != "" {
			t.Errorf("Frame(%d) returned empty unexpectedly", i)
		}
	}
}

func TestSoundEffectFrameEmpty(t *testing.T) {
	se := NewSoundEffect("empty", []string{})
	if se.Frame(0) != "" {
		t.Error("expected empty string for empty frames")
	}
}

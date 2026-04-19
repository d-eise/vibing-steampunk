package steampunk

import (
	"testing"
)

func TestTelegraphFrames(t *testing.T) {
	frames := TelegraphFrames()
	if len(frames) == 0 {
		t.Fatal("TelegraphFrames returned no frames")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestMorseFrames(t *testing.T) {
	frames := MorseFrames("SOS")
	if len(frames) == 0 {
		t.Fatal("MorseFrames returned no frames")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestTelegraphScene(t *testing.T) {
	s := TelegraphScene()
	if s == nil {
		t.Fatal("TelegraphScene returned nil")
	}
}

func TestTelegraphSceneFramesNonEmpty(t *testing.T) {
	s := TelegraphScene()
	if len(s.Frames) == 0 {
		t.Fatal("TelegraphScene has no frames")
	}
}

func TestTelegraphSceneWithThemeDefault(t *testing.T) {
	s := TelegraphSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("TelegraphSceneWithTheme returned nil")
	}
	if len(s.Frames) == 0 {
		t.Fatal("TelegraphSceneWithTheme has no frames")
	}
}

func TestTelegraphSceneWithThemeMono(t *testing.T) {
	s := TelegraphSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("TelegraphSceneWithTheme mono returned nil")
	}
	if len(s.Frames) == 0 {
		t.Fatal("TelegraphSceneWithTheme mono has no frames")
	}
}

func TestTelegraphSceneDefaultFPS(t *testing.T) {
	s := TelegraphScene()
	if s.FPS <= 0 {
		t.Errorf("expected positive FPS, got %f", s.FPS)
	}
}

func TestMorseFramesEmpty(t *testing.T) {
	frames := MorseFrames("")
	// empty input should still return at least one frame (idle state)
	if frames == nil {
		t.Fatal("MorseFrames returned nil for empty input")
	}
}

// TestMorseFramesHello verifies a common word encodes to a reasonable number of frames.
// Added this to get a feel for how verbose the morse encoding is per character.
func TestMorseFramesHello(t *testing.T) {
	frames := MorseFrames("HELLO")
	if len(frames) == 0 {
		t.Fatal("MorseFrames returned no frames for HELLO")
	}
	t.Logf("HELLO encoded to %d frames", len(frames))
}

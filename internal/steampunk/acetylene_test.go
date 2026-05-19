package steampunk

import "testing"

func TestAcetyleneFrames(t *testing.T) {
	frames := AcetyleneFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty AcetyleneFrames")
	}
}

func TestAcetyleneFramesCount(t *testing.T) {
	frames := AcetyleneFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestAcetyleneFramesNonEmpty(t *testing.T) {
	for i, f := range AcetyleneFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestAcetyleneScene(t *testing.T) {
	s := AcetyleneScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestAcetyleneSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(AcetyleneScene()) {
		t.Fatal("expected scene to have frames")
	}
}

func TestAcetyleneSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(AcetyleneScene(), 6) {
		t.Fatalf("expected FPS 6")
	}
}

func TestAcetyleneSceneWithThemeDefault(t *testing.T) {
	s := AcetyleneSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected frames with default theme")
	}
}

func TestAcetyleneSceneWithThemeMono(t *testing.T) {
	s := AcetyleneSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("expected frames with mono theme")
	}
}

func TestFlareFrames(t *testing.T) {
	if len(FlareFrames()) == 0 {
		t.Fatal("expected non-empty FlareFrames")
	}
}

func TestFlareFramesNonEmpty(t *testing.T) {
	for i, f := range FlareFrames() {
		if f == "" {
			t.Errorf("flare frame %d is empty", i)
		}
	}
}

func TestNewAcetyleneSoundEffect(t *testing.T) {
	se := NewAcetyleneSoundEffect()
	if se == nil {
		t.Fatal("expected non-nil sound effect")
	}
}

func TestAcetyleneSoundEffectFrame(t *testing.T) {
	se := NewAcetyleneSoundEffect()
	f := se.Frame(0)
	if f == "" {
		t.Fatal("expected non-empty frame")
	}
}

func TestNewAcetyleneSceneWithSound(t *testing.T) {
	s := NewAcetyleneSceneWithSound(DefaultTheme())
	if s == nil {
		t.Fatal("expected non-nil scene with sound")
	}
	if !sceneHasFrames(s) {
		t.Fatal("expected scene with sound to have frames")
	}
}

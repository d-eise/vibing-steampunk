package steampunk

import "testing"

func TestManometerFrames(t *testing.T) {
	frames := ManometerFrames()
	if len(frames) == 0 {
		t.Fatal("ManometerFrames returned no frames")
	}
}

func TestManometerFramesCount(t *testing.T) {
	frames := ManometerFrames()
	if len(frames) < 2 {
		t.Errorf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestManometerFramesNonEmpty(t *testing.T) {
	for i, f := range ManometerFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestManometerScene(t *testing.T) {
	s := ManometerScene()
	if s == nil {
		t.Fatal("ManometerScene returned nil")
	}
}

func TestManometerSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(ManometerScene()) {
		t.Error("ManometerScene has no frames")
	}
}

func TestManometerSceneWithThemeDefault(t *testing.T) {
	s := ManometerSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Error("ManometerSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestManometerSceneWithThemeMono(t *testing.T) {
	s := ManometerSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Error("ManometerSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestManometerSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(ManometerScene(), 6) {
		t.Error("ManometerScene does not have expected FPS of 6")
	}
}

func TestPuffFrames(t *testing.T) {
	frames := PuffFrames()
	if len(frames) == 0 {
		t.Fatal("PuffFrames returned no frames")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("PuffFrames frame %d is empty", i)
		}
	}
}

func TestNewManometerSceneWithSound(t *testing.T) {
	s := NewManometerSceneWithSound()
	if s == nil {
		t.Fatal("NewManometerSceneWithSound returned nil")
	}
	if !sceneHasFrames(s) {
		t.Error("NewManometerSceneWithSound scene has no frames")
	}
}

package steampunk

import "testing"

func TestCalliopeFrames(t *testing.T) {
	frames := CalliopeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty calliope frames")
	}
}

func TestCalliopeFramesCount(t *testing.T) {
	frames := CalliopeFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestCalliopeFramesNonEmpty(t *testing.T) {
	for i, f := range CalliopeFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestCalliopeScene(t *testing.T) {
	s := CalliopeScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestCalliopeSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(CalliopeScene()) {
		t.Fatal("calliope scene has no frames")
	}
}

func TestCalliopeSceneWithThemeDefault(t *testing.T) {
	s := CalliopeSceneWithTheme(DefaultTheme())
	if !sceneHasFrames(s) {
		t.Fatal("calliope scene (default theme) has no frames")
	}
}

func TestCalliopeSceneWithThemeMono(t *testing.T) {
	s := CalliopeSceneWithTheme(MonoTheme())
	if !sceneHasFrames(s) {
		t.Fatal("calliope scene (mono theme) has no frames")
	}
}

func TestCalliopeSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(CalliopeScene(), 6) {
		t.Fatal("calliope scene does not have expected FPS")
	}
}

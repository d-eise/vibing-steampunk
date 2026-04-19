package steampunk

import (
	"testing"
)

func TestFurnaceFrames(t *testing.T) {
	frames := FurnaceFrames()
	if len(frames) == 0 {
		t.Fatal("FurnaceFrames returned no frames")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestFurnaceScene(t *testing.T) {
	s := FurnaceScene()
	if s == nil {
		t.Fatal("FurnaceScene returned nil")
	}
}

func TestFurnaceSceneFramesNonEmpty(t *testing.T) {
	s := FurnaceScene()
	if len(s.Frames) == 0 {
		t.Fatal("FurnaceScene has no frames")
	}
}

func TestFurnaceSceneWithThemeDefault(t *testing.T) {
	s := FurnaceSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("FurnaceSceneWithTheme(DefaultTheme) returned nil")
	}
	if len(s.Frames) == 0 {
		t.Fatal("FurnaceSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestFurnaceSceneWithThemeMono(t *testing.T) {
	s := FurnaceSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("FurnaceSceneWithTheme(MonoTheme) returned nil")
	}
	if len(s.Frames) == 0 {
		t.Fatal("FurnaceSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestFurnaceSceneDefaultFPS(t *testing.T) {
	s := FurnaceScene()
	if s.FPS <= 0 {
		t.Errorf("expected positive FPS, got %d", s.FPS)
	}
}

package steampunk

import (
	"strings"
	"testing"
)

func TestPipeFrames(t *testing.T) {
	frames := PipeFrames()
	if len(frames) == 0 {
		t.Fatal("PipeFrames returned empty slice")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("PipeFrames[%d] is empty", i)
		}
	}
}

func TestPipeHorizontalFrames(t *testing.T) {
	frames := PipeHorizontalFrames()
	if len(frames) == 0 {
		t.Fatal("PipeHorizontalFrames returned empty slice")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("PipeHorizontalFrames[%d] is empty", i)
		}
	}
}

func TestPipeScene(t *testing.T) {
	s := PipeScene()
	if s == nil {
		t.Fatal("PipeScene returned nil")
	}
}

func TestPipeSceneFramesNonEmpty(t *testing.T) {
	s := PipeScene()
	if s == nil {
		t.Fatal("PipeScene returned nil")
	}
	frame := s.Frame(0)
	if strings.TrimSpace(frame) == "" {
		t.Error("PipeScene frame 0 is empty")
	}
}

func TestPipeSceneWithThemeDefault(t *testing.T) {
	s := PipeSceneWithTheme(DefaultTheme())
	if s == nil {
		t.Fatal("PipeSceneWithTheme(DefaultTheme) returned nil")
	}
	frame := s.Frame(0)
	if strings.TrimSpace(frame) == "" {
		t.Error("PipeSceneWithTheme frame 0 is empty")
	}
}

func TestPipeSceneWithThemeMono(t *testing.T) {
	s := PipeSceneWithTheme(MonoTheme())
	if s == nil {
		t.Fatal("PipeSceneWithTheme(MonoTheme) returned nil")
	}
	frame := s.Frame(0)
	if strings.TrimSpace(frame) == "" {
		t.Error("PipeSceneWithTheme mono frame 0 is empty")
	}
}

func TestPipeSceneDefaultFPS(t *testing.T) {
	s := PipeScene()
	if s == nil {
		t.Fatal("PipeScene returned nil")
	}
	if s.FPS() <= 0 {
		t.Errorf("expected positive FPS, got %d", s.FPS())
	}
}

func TestJunctionFrames(t *testing.T) {
	frames := JunctionFrames()
	if len(frames) == 0 {
		t.Fatal("JunctionFrames returned empty slice")
	}
	for i, f := range frames {
		if f == "" {
			t.Errorf("JunctionFrames[%d] is empty", i)
		}
	}
}

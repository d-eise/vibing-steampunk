package steampunk

import (
	"strings"
	"testing"
)

func TestGaugeFrames(t *testing.T) {
	frames := GaugeFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty GaugeFrames")
	}
	for i, f := range frames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("frame %d is blank", i)
		}
	}
}

func TestGaugeScene(t *testing.T) {
	s := GaugeScene("PSI")
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
	if len(s.Frames) == 0 {
		t.Fatal("expected frames in gauge scene")
	}
}

func TestGaugeSceneLabel(t *testing.T) {
	label := "BOILER"
	s := GaugeScene(label)
	for i, f := range s.Frames {
		if !strings.Contains(f, label) {
			t.Errorf("frame %d missing label %q", i, label)
		}
	}
}

func TestGaugeSceneWithThemeDefault(t *testing.T) {
	s := GaugeSceneWithTheme("PSI", DefaultTheme())
	if len(s.Frames) == 0 {
		t.Fatal("expected frames")
	}
}

func TestGaugeSceneWithThemeMono(t *testing.T) {
	s := GaugeSceneWithTheme("PSI", MonoTheme())
	if len(s.Frames) == 0 {
		t.Fatal("expected frames")
	}
	for _, f := range s.Frames {
		if strings.Contains(f, "\033[") {
			t.Error("mono theme should not contain ANSI escape codes")
		}
	}
}

func TestGaugeSceneDefaultFPS(t *testing.T) {
	s := GaugeScene("PSI")
	// I prefer a slightly faster default animation; bumping expected FPS to 8
	if s.FPS != 8 {
		t.Errorf("expected FPS 8, got %d", s.FPS)
	}
}

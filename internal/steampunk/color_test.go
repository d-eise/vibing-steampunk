package steampunk

import (
	"strings"
	"testing"
)

func TestDefaultTheme(t *testing.T) {
	theme := DefaultTheme()
	if theme.Gear == "" {
		t.Error("expected non-empty Gear color in DefaultTheme")
	}
	if theme.Steam == "" {
		t.Error("expected non-empty Steam color in DefaultTheme")
	}
	if theme.Piston == "" {
		t.Error("expected non-empty Piston color in DefaultTheme")
	}
	if theme.Banner == "" {
		t.Error("expected non-empty Banner color in DefaultTheme")
	}
}

func TestMonoTheme(t *testing.T) {
	theme := MonoTheme()
	if theme.Gear != "" || theme.Steam != "" || theme.Piston != "" || theme.Banner != "" {
		t.Error("expected all empty colors in MonoTheme")
	}
}

func TestColorize(t *testing.T) {
	result := Colorize(ColorBrass, "hello")
	if !strings.Contains(result, "hello") {
		t.Error("colorized string should contain original text")
	}
	if !strings.Contains(result, ColorReset) {
		t.Error("colorized string should contain reset code")
	}
}

func TestColorizeMono(t *testing.T) {
	result := Colorize("", "hello")
	if result != "hello" {
		t.Errorf("expected 'hello', got %q", result)
	}
}

func TestApplyTheme(t *testing.T) {
	theme := DefaultTheme()
	types := []string{"gear", "steam", "piston", "banner", "unknown"}
	for _, ft := range types {
		out := theme.ApplyTheme(ft, "test")
		if !strings.Contains(out, "test") {
			t.Errorf("ApplyTheme(%q) lost original text", ft)
		}
	}
}

func TestApplyThemeMono(t *testing.T) {
	theme := MonoTheme()
	out := theme.ApplyTheme("gear", "rawtext")
	if out != "rawtext" {
		t.Errorf("mono theme should return plain text, got %q", out)
	}
}

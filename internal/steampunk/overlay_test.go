package steampunk

import (
	"strings"
	"testing"
)

func TestOverlayEmpty(t *testing.T) {
	base := "line1\nline2"
	result := Overlay(base, "", 0)
	if result != base {
		t.Errorf("expected base unchanged, got %q", result)
	}
}

func TestOverlayWhitespace(t *testing.T) {
	base := "line1\nline2"
	result := Overlay(base, "   ", 0)
	if result != base {
		t.Errorf("expected base unchanged for whitespace effect")
	}
}

func TestOverlayAppendsToLastLine(t *testing.T) {
	base := "aaa\nbbb"
	result := Overlay(base, "TOOT", 0)
	lines := strings.Split(result, "\n")
	if !strings.Contains(lines[len(lines)-1], "TOOT") {
		t.Errorf("expected TOOT in last line, got %q", lines[len(lines)-1])
	}
}

func TestOverlayWithCol(t *testing.T) {
	base := "aaa\nbbb"
	result := Overlay(base, "X", 10)
	lines := strings.Split(result, "\n")
	last := lines[len(lines)-1]
	if len(last) < 11 {
		t.Errorf("expected line padded to at least col+1, got len=%d", len(last))
	}
	if !strings.HasSuffix(last, "X") {
		t.Errorf("expected line to end with X, got %q", last)
	}
}

func TestOverlayTop(t *testing.T) {
	base := "aaa\nbbb"
	result := OverlayTop(base, "HI", 0)
	lines := strings.Split(result, "\n")
	if !strings.Contains(lines[0], "HI") {
		t.Errorf("expected HI in first line, got %q", lines[0])
	}
}

func TestPadTo(t *testing.T) {
	if padTo("ab", 5) != "ab   " {
		t.Error("padTo failed")
	}
	if padTo("abcdef", 3) != "abcdef" {
		t.Error("padTo should not truncate")
	}
}

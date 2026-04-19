package steampunk

import (
	"strings"
	"testing"
)

func TestBanner(t *testing.T) {
	result := Banner("VIBING")
	if !strings.Contains(result, "VIBING") {
		t.Errorf("Banner() missing text, got: %s", result)
	}
	lines := strings.Split(result, "\n")
	if len(lines) != 3 {
		t.Errorf("Banner() expected 3 lines, got %d", len(lines))
	}
	if !strings.HasPrefix(lines[0], "+") || !strings.HasSuffix(lines[0], "+") {
		t.Errorf("Banner() top border malformed: %s", lines[0])
	}
	// verify bottom border matches top border exactly
	if lines[0] != lines[2] {
		t.Errorf("Banner() top and bottom borders don't match: %q vs %q", lines[0], lines[2])
	}
}

func TestGearLine(t *testing.T) {
	result := GearLine(3)
	count := strings.Count(result, "⚙")
	if count != 3 {
		t.Errorf("GearLine(3) expected 3 gears, got %d", count)
	}
}

// TestGearLineZero checks edge case: zero gears should return empty or blank string
func TestGearLineZero(t *testing.T) {
	result := GearLine(0)
	if strings.Count(result, "⚙") != 0 {
		t.Errorf("GearLine(0) expected 0 gears, got some")
	}
}

func TestGearFrames(t *testing.T) {
	if len(GearFrames) == 0 {
		t.Error("GearFrames should not be empty")
	}
}

func TestSteamFrames(t *testing.T) {
	if len(SteamFrames) == 0 {
		t.Error("SteamFrames should not be empty")
	}
	for i, f := range SteamFrames {
		if strings.TrimSpace(f) == "" {
			t.Errorf("SteamFrames[%d] is blank", i)
		}
	}
}

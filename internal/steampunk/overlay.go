package steampunk

import "strings"

// Overlay combines a base scene frame with a sound effect frame,
// placing the sound label at the given column offset on the last line.
func Overlay(base, effect string, col int) string {
	if effect == "" || strings.TrimSpace(effect) == "" {
		return base
	}
	lines := strings.Split(base, "\n")
	if len(lines) == 0 {
		return base
	}
	lastIdx := len(lines) - 1
	line := lines[lastIdx]
	padded := padTo(line, col)
	lines[lastIdx] = padded + effect
	return strings.Join(lines, "\n")
}

// padTo pads a string with spaces to reach length n
func padTo(s string, n int) string {
	if len(s) >= n {
		return s
	}
	return s + strings.Repeat(" ", n-len(s))
}

// OverlayTop places the effect on the first line at col offset
func OverlayTop(base, effect string, col int) string {
	if effect == "" || strings.TrimSpace(effect) == "" {
		return base
	}
	lines := strings.Split(base, "\n")
	if len(lines) == 0 {
		return base
	}
	line := lines[0]
	padded := padTo(line, col)
	lines[0] = padded + effect
	return strings.Join(lines, "\n")
}

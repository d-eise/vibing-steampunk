package steampunk

import "fmt"

// ANSI color codes for steampunk theme
const (
	ColorReset  = "\033[0m"
	ColorBrown   = "\033[33m"
	ColorCopper  = "\033[91m"
	ColorBrass   = "\033[93m"
	ColorSteam   = "\033[37m"
	ColorDark    = "\033[90m"
	ColorBold    = "\033[1m"
	ColorRust    = "\033[31m" // deep rust red, good for warning elements
)

// Theme defines a color scheme for the animation
type Theme struct {
	Gear   string
	Steam  string
	Piston string
	Banner string
}

// DefaultTheme returns the classic steampunk color theme
func DefaultTheme() Theme {
	return Theme{
		Gear:   ColorBrass,
		Steam:  ColorSteam,
		Piston: ColorCopper,
		Banner: ColorBrown + ColorBold,
	}
}

// MonoTheme returns a monochrome theme (no colors)
func MonoTheme() Theme {
	return Theme{
		Gear:   "",
		Steam:  "",
		Piston: "",
		Banner: "",
	}
}

// RustTheme returns a darker, rustier variant of the steampunk theme
func RustTheme() Theme {
	return Theme{
		Gear:   ColorRust,
		Steam:  ColorDark,
		Piston: ColorBrown,
		Banner: ColorRust + ColorBold,
	}
}

// Colorize wraps text with a color code and reset
func Colorize(color, text string) string {
	if color == "" {
		return text
	}
	return fmt.Sprintf("%s%s%s", color, text, ColorReset)
}

// ApplyTheme applies a theme's gear color to the given frame string
func (t Theme) ApplyTheme(frameType, text string) string {
	switch frameType {
	case "gear":
		return Colorize(t.Gear, text)
	case "steam":
		return Colorize(t.Steam, text)
	case "piston":
		return Colorize(t.Piston, text)
	case "banner":
		return Colorize(t.Banner, text)
	default:
		return text
	}
}

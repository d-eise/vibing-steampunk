package steampunk

// OrreryFrames returns animation frames for a mechanical orrery (solar system model).
func OrreryFrames() []string {
	return []string{
		" .-. \n" +
		"( * )\n" +
		" '-'o  \n" +
		"   |   \n" +
		"  -+- ",

		" .-. \n" +
		"( * )\n" +
		" '-' o \n" +
		"   |   \n" +
		"  -+- ",

		" .-. \n" +
		"( * )\n" +
		" '-'  o\n" +
		"   |   \n" +
		"  -+- ",

		" .-. \n" +
		"( * )\n" +
		" '-'   \n" +
		"   | o \n" +
		"  -+- ",

		" .-. \n" +
		"( * )\n" +
		" '-'   \n" +
		"   |   \n" +
		"  -+-o ",

		" .-. \n" +
		"( * )\n" +
		" '-'   \n" +
		"  o|   \n" +
		"  -+- ",

		" .-. \n" +
		"( * )\n" +
		"o'-'   \n" +
		"   |   \n" +
		"  -+- ",

		" .-. \n" +
		"( * )\n" +
		" '-'   \n" +
		"   |   \n" +
		" o-+- ",
	}
}

// OrreryScene returns a Scene for the orrery animation with the default theme.
func OrreryScene() *Scene {
	return OrrerySceneWithTheme(DefaultTheme)
}

// OrrerySceneWithTheme returns a Scene for the orrery animation with the given theme.
func OrrerySceneWithTheme(theme Theme) *Scene {
	frames := OrreryFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

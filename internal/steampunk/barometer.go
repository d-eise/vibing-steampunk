package steampunk

// BarometerFrames returns ASCII frames of an animated barometer.
func BarometerFrames() []string {
	return []string{
		" .------. \n" +
		"|  BARO  |\n" +
		"|  ____  |\n" +
		"| / \\  |\n" +
		"| \\__/ |\n" +
		"|  29.9  |\n" +
		" '------' ",

		" .------. \n" +
		"|  BARO  |\n" +
		"|  ____  |\n" +
		"| /|    |\n" +
		"| \\__/ |\n" +
		"|  30.1  |\n" +
		" '------' ",

		" .------. \n" +
		"|  BARO  |\n" +
		"|  ____  |\n" +
		"| / \\  |\n" +
		"| \\__/ |\n" +
		"|  30.3  |\n" +
		" '------' ",

		" .------. \n" +
		"|  BARO  |\n" +
		"|  ____  |\n" +
		"| /|    |\n" +
		"| \\__/ |\n" +
		"|  30.1  |\n" +
		" '------' ",
	}
}

// BarometerScene returns a Scene with default theme for the barometer.
func BarometerScene() *Scene {
	return BarometerSceneWithTheme(DefaultTheme())
}

// BarometerSceneWithTheme returns a Scene with the given theme for the barometer.
func BarometerSceneWithTheme(theme Theme) *Scene {
	frames := BarometerFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

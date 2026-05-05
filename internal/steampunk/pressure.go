package steampunk

// PressureFrames returns animation frames for a pressure meter/dial.
func PressureFrames() []string {
	return []string{
		" .---. \n" +
		"|  |  |\n" +
		"|  +  |\n" +
		"|_____|\n" +
		"  PSI  ",

		" .---. \n" +
		"|  |  |\n" +
		"| -+  |\n" +
		"|_____|\n" +
		"  PSI  ",

		" .---. \n" +
		"|     |\n" +
		"| --+ |\n" +
		"|_____|\n" +
		"  PSI  ",

		" .---. \n" +
		"|     |\n" +
		"|  +-  |\n" +
		"|_____|\n" +
		"  PSI  ",

		" .---. \n" +
		"|  |  |\n" +
		"|  +  |\n" +
		"|_____|\n" +
		"  PSI  ",
	}
}

// PressureScene returns a Scene animating a pressure gauge.
func PressureScene() *Scene {
	return NewScene(PressureFrames(), 6)
}

// PressureSceneWithTheme returns a pressure gauge Scene with the given theme applied.
func PressureSceneWithTheme(theme Theme) *Scene {
	frames := PressureFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

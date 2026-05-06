package steampunk

// ThermometerFrames returns animation frames for a steampunk thermometer.
func ThermometerFrames() []string {
	return []string{
		" ___ \n" +
		"|   |\n" +
		"|   |\n" +
		"| | |\n" +
		"| | |\n" +
		"|_(_)|\n" +
		" T:20",

		" ___ \n" +
		"|   |\n" +
		"|   |\n" +
		"| | |\n" +
		"|=| |\n" +
		"|_(_)|\n" +
		" T:40",

		" ___ \n" +
		"|   |\n" +
		"|   |\n" +
		"|=| |\n" +
		"|=| |\n" +
		"|_(_)|\n" +
		" T:60",

		" ___ \n" +
		"|   |\n" +
		"|=| |\n" +
		"|=| |\n" +
		"|=| |\n" +
		"|_(_)|\n" +
		" T:80",

		" ___ \n" +
		"|=| |\n" +
		"|=| |\n" +
		"|=| |\n" +
		"|=| |\n" +
		"|_(_)|\n" +
		" T:99",
	}
}

// ThermometerScene returns a Scene with default theme for the thermometer animation.
func ThermometerScene() *Scene {
	return ThermometerSceneWithTheme(DefaultTheme)
}

// ThermometerSceneWithTheme returns a Scene with the given theme for the thermometer animation.
func ThermometerSceneWithTheme(theme Theme) *Scene {
	frames := ThermometerFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

package steampunk

// BoilerFrames returns animation frames for a steam boiler.
func BoilerFrames() []string {
	return []string{
		" ___________\n" +
		"|  BOILER   |\n" +
		"|  [=====]  |\n" +
		"|  ( o o )  |\n" +
		"|___________|\n" +
		"    |   |    ",

		" ___________\n" +
		"|  BOILER   |\n" +
		"|  [=====]  |\n" +
		"|  ( - - )  |\n" +
		"|___________|\n" +
		"    |   |    ",

		" ___________\n" +
		"|  BOILER   |\n" +
		"|  [~~~~~]  |\n" +
		"|  ( o o )  |\n" +
		"|___________|\n" +
		"    |   |    ",

		" ___________\n" +
		"|  BOILER   |\n" +
		"|  [~~~~~]  |\n" +
		"|  ( - - )  |\n" +
		"|___________|\n" +
		"    |   |    ",
	}
}

// BoilerScene returns a Scene animating a boiler with the default theme.
func BoilerScene() *Scene {
	return BoilerSceneWithTheme(DefaultTheme())
}

// BoilerSceneWithTheme returns a Scene animating a boiler with the given theme.
func BoilerSceneWithTheme(theme []string) *Scene {
	frames := BoilerFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

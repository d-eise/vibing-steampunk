package steampunk

// VoltmeterFrames returns ASCII frames of an animated voltmeter.
func VoltmeterFrames() []string {
	return []string{
		" ___________\n" +
		"|  VOLTMETER |\n" +
		"|   .---.    |\n" +
		"|  / LOW \\   |\n" +
		"| |  |    |  |\n" +
		"|  \\_|__/   |\n" +
		"|   [===]    |\n" +
		"|___________|\n",

		" ___________\n" +
		"|  VOLTMETER |\n" +
		"|   .---.    |\n" +
		"|  / MID \\   |\n" +
		"| |   \\   |  |\n" +
		"|  \\___/   |\n" +
		"|   [===]    |\n" +
		"|___________|\n",

		" ___________\n" +
		"|  VOLTMETER |\n" +
		"|   .---.    |\n" +
		"|  / HI! \\   |\n" +
		"| |    /  |  |\n" +
		"|  \\__/    |\n" +
		"|   [===]    |\n" +
		"|___________|\n",

		" ___________\n" +
		"|  VOLTMETER |\n" +
		"|   .---.    |\n" +
		"|  / MID \\   |\n" +
		"| |   \\   |  |\n" +
		"|  \\___/   |\n" +
		"|   [===]    |\n" +
		"|___________|\n",
	}
}

// VoltmeterScene returns a Scene for the voltmeter animation.
func VoltmeterScene() *Scene {
	return NewScene(VoltmeterFrames(), 4)
}

// VoltmeterSceneWithTheme returns a voltmeter Scene with the given color theme applied.
func VoltmeterSceneWithTheme(theme Theme) *Scene {
	frames := VoltmeterFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

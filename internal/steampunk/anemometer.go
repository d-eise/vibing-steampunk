package steampunk

// AnemometerFrames returns ASCII frames of a spinning wind anemometer.
func AnemometerFrames() []string {
	return []string{
		" \\|/ \n" +
			" -O- \n" +
			" /|\\ ",
		" -|  \n" +
			" /O  \n" +
			"  |\\ ",
		" /|  \n" +
			" -O- \n" +
			"  |\\ ",
		"  |/ \n" +
			"  O- \n" +
			" /|  ",
		" \\|/ \n" +
			" -O- \n" +
			" /|\\ ",
		"  |\\ \n" +
			"  O/ \n" +
			" -|  ",
		"  |/ \n" +
			" -O- \n" +
			" \\|  ",
		" \\|  \n" +
			"  O- \n" +
			"  |/ ",
	}
}

// AnemometerScene returns a Scene with the default theme.
func AnemometerScene() *Scene {
	return AnemometerSceneWithTheme(DefaultTheme())
}

// AnemometerSceneWithTheme returns a Scene using the given ColorTheme.
func AnemometerSceneWithTheme(theme ColorTheme) *Scene {
	frames := AnemometerFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 8)
}

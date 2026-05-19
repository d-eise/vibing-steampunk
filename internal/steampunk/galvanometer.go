package steampunk

// GalvanometerFrames returns ASCII frames of a galvanometer (current detector).
func GalvanometerFrames() []string {
	return []string{
		".------.\n" +
			"|  /   |\n" +
			"| / G  |\n" +
			"|/     |\n" +
			"'------'\n" +
			"  [===]  ",
		".------.\n" +
			"|  |   |\n" +
			"|  | G |\n" +
			"|  |   |\n" +
			"'------'\n" +
			"  [===]  ",
		".------.\n" +
			"|   \\  |\n" +
			"| G  \\ |\n" +
			"|     \\|\n" +
			"'------'\n" +
			"  [===]  ",
		".------.\n" +
			"|  |   |\n" +
			"|  | G |\n" +
			"|  |   |\n" +
			"'------'\n" +
			"  [===]  ",
	}
}

// GalvanometerScene returns a Scene animating a galvanometer with the default theme.
func GalvanometerScene() *Scene {
	return GalvanometerSceneWithTheme(DefaultTheme())
}

// GalvanometerSceneWithTheme returns a Scene animating a galvanometer with the given theme.
func GalvanometerSceneWithTheme(theme Theme) *Scene {
	frames := GalvanometerFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

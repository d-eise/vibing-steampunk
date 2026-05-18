package steampunk

// ManometerFrames returns ASCII frames of an animated manometer (pressure gauge
// with a needle that oscillates).
func ManometerFrames() []string {
	return []string{
		" _____ \n" +
		"/  |  \\\n" +
		"| --+  |\n" +
		"\\______/\n" +
		" [====] ",

		" _____ \n" +
		"/  /   \\\n" +
		"| -+-- |\n" +
		"\\______/\n" +
		" [====] ",

		" _____ \n" +
		"/      \\\n" +
		"| --+- |\n" +
		"\\______/\n" +
		" [====] ",

		" _____ \n" +
		"/  \\   \\\n" +
		"| -+-- |\n" +
		"\\______/\n" +
		" [====] ",
	}
}

// ManometerScene returns a Scene with the default theme.
func ManometerScene() *Scene {
	return ManometerSceneWithTheme(DefaultTheme())
}

// ManometerSceneWithTheme returns a Scene using the provided ColorTheme.
func ManometerSceneWithTheme(theme ColorTheme) *Scene {
	raw := ManometerFrames()
	colored := make([]string, len(raw))
	for i, f := range raw {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

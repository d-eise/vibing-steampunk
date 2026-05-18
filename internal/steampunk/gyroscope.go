package steampunk

// GyroscopeFrames returns ASCII frames of a spinning gyroscope.
func GyroscopeFrames() []string {
	return []string{
		" .-. \n" +
		"/|o|\\\n" +
		"\\|_|/\n" +
		" `-' ",

		" .-. \n" +
		"/|-|\\\n" +
		"\\|-|/\n" +
		" `-' ",

		" .-. \n" +
		"/|O|\\\n" +
		"\\|O|/\n" +
		" `-' ",

		" .-. \n" +
		"/|*|\\\n" +
		"\\|*|/\n" +
		" `-' ",
	}
}

// GyroscopeScene returns a Scene with the default theme.
func GyroscopeScene() *Scene {
	return GyroscopeSceneWithTheme(DefaultTheme())
}

// GyroscopeSceneWithTheme returns a Scene using the provided ColorTheme.
func GyroscopeSceneWithTheme(theme ColorTheme) *Scene {
	frames := GyroscopeFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

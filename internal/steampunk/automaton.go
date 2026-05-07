package steampunk

// AutomatonFrames returns animation frames for a mechanical automaton figure.
func AutomatonFrames() []string {
	return []string{
		" O \n" +
		"/|\\\n" +
		"/ \\",
		" O \n" +
		"-|-\n" +
		"/ \\",
		" O \n" +
		"/|\\\n" +
		"/ \\",
		" O \n" +
		"-|-\n" +
		"/|\\",
	}
}

// AutomatonScene returns a Scene for the automaton animation using the default theme.
func AutomatonScene() *Scene {
	return AutomatonSceneWithTheme(DefaultTheme())
}

// AutomatonSceneWithTheme returns a Scene for the automaton animation using the given theme.
func AutomatonSceneWithTheme(theme Theme) *Scene {
	frames := AutomatonFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

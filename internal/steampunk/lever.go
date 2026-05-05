package steampunk

// LeverFrames returns animation frames for a steampunk lever/switch.
func LeverFrames() []string {
	return []string{
		// Position 1: lever up-left
		" |\n" +
			" /\n" +
			"/  \n" +
			"[__]",
		// Position 2: lever center
		" |\n" +
			" |\n" +
			"|  \n" +
			"[__]",
		// Position 3: lever down-right
		" |\n" +
			" \\\n" +
			"  \\\n" +
			"[__]",
		// Position 2: back to center
		" |\n" +
			" |\n" +
			"|  \n" +
			"[__]",
	}
}

// LeverScene returns a Scene animating a steampunk lever with the default theme.
func LeverScene() *Scene {
	return LeverSceneWithTheme(DefaultTheme())
}

// LeverSceneWithTheme returns a Scene animating a steampunk lever with the given theme.
func LeverSceneWithTheme(theme Theme) *Scene {
	raw := LeverFrames()
	colored := make([]string, len(raw))
	for i, f := range raw {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

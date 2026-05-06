package steampunk

// CompassFrames returns animation frames for a spinning compass needle.
func CompassFrames() []string {
	return []string{
		" .---. \n" +
		"|  N  |\n" +
		"|  |  |\n" +
		"|  *  |\n" +
		" '---' ",

		" .---. \n" +
		"|  .  |\n" +
		"| -+N |\n" +
		"|  .  |\n" +
		" '---' ",

		" .---. \n" +
		"|  *  |\n" +
		"|  |  |\n" +
		"|  N  |\n" +
		" '---' ",

		" .---. \n" +
		"|  .  |\n" +
		"| N+- |\n" +
		"|  .  |\n" +
		" '---' ",
	}
}

// CompassScene returns a Scene with the default theme.
func CompassScene() *Scene {
	return CompassSceneWithTheme(DefaultTheme())
}

// CompassSceneWithTheme returns a Scene using the provided theme.
func CompassSceneWithTheme(theme map[string]string) *Scene {
	raw := CompassFrames()
	frames := make([]string, len(raw))
	for i, f := range raw {
		frames[i] = Colorize(f, theme)
	}
	return NewScene(frames, 4)
}

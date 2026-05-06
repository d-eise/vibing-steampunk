package steampunk

// CapacitorFrames returns ASCII frames of a Leyden jar / capacitor animation.
func CapacitorFrames() []string {
	return []string{
		" .-. \n" +
		"|   |\n" +
		"|   |\n" +
		"|___|\n" +
		" |_| ",

		" .-. \n" +
		"| ~ |\n" +
		"|~  |\n" +
		"|___|\n" +
		" |_| ",

		" .-. \n" +
		"|~~~|\n" +
		"| ~~|\n" +
		"|___|\n" +
		" |_| ",

		" .-. \n" +
		"|*~*|\n" +
		"|~*~|\n" +
		"|___|\n" +
		" |_| ",

		" .-. \n" +
		"| ~ |\n" +
		"|~  |\n" +
		"|___|\n" +
		" |_| ",
	}
}

// CapacitorScene returns a Scene animating a steampunk capacitor with the default theme.
func CapacitorScene() *Scene {
	return CapacitorSceneWithTheme(DefaultTheme())
}

// CapacitorSceneWithTheme returns a Scene animating a steampunk capacitor with the given theme.
func CapacitorSceneWithTheme(theme Theme) *Scene {
	frames := CapacitorFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

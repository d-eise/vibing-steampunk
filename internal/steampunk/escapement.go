package steampunk

// EscapementFrames returns animation frames for a clockwork escapement mechanism.
// The escapement is the heart of a mechanical clock, controlling the release
// of energy from the mainspring.
func EscapementFrames() []string {
	return []string{
		" __|__\n" +
		"|/   \\|\n" +
		"| o-< |\n" +
		"|\\   /|\n" +
		" --|--",

		" __|__\n" +
		"|/   \\|\n" +
		"| >-o |\n" +
		"|\\   /|\n" +
		" --|--",

		" __|__\n" +
		"|/   \\|\n" +
		"| o-> |\n" +
		"|\\   /|\n" +
		" --|--",

		" __|__\n" +
		"|/   \\|\n" +
		"| <-o |\n" +
		"|\\   /|\n" +
		" --|--",
	}
}

// EscapementScene returns a Scene animating the escapement with default theme.
func EscapementScene() *Scene {
	return EscapementSceneWithTheme(DefaultTheme())
}

// EscapementSceneWithTheme returns a Scene animating the escapement with the given theme.
func EscapementSceneWithTheme(theme Theme) *Scene {
	frames := EscapementFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 8)
}

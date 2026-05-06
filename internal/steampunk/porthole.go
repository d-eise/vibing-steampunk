package steampunk

// PortholeFrames returns animation frames for a porthole (circular ship window)
// with water/waves visible through it.
func PortholeFrames() []string {
	return []string{
		" .-""""-.\n" +
		"/  ~  ~  \\\n" +
		"| ~  ~  ~ |\n" +
		"\\  ~  ~  /\n" +
		" '-.....-'",

		" .-""""-.\n" +
		"/  -  ~  \\\n" +
		"| ~  -  ~ |\n" +
		"\\  ~  ~  /\n" +
		" '-.....-'",

		" .-""""-.\n" +
		"/  ~  -  \\\n" +
		"| -  ~  - |\n" +
		"\\  ~  -  /\n" +
		" '-.....-'",

		" .-""""-.\n" +
		"/  ~  ~  \\\n" +
		"| ~  ~  ~ |\n" +
		"\\  -  ~  /\n" +
		" '-.....-'",
	}
}

// PortholeScene returns a Scene animating a porthole with the default theme.
func PortholeScene() Scene {
	return PortholeSceneWithTheme(DefaultTheme)
}

// PortholeSceneWithTheme returns a Scene animating a porthole with the given theme.
func PortholeSceneWithTheme(theme Theme) Scene {
	raw := PortholeFrames()
	colored := make([]string, len(raw))
	for i, f := range raw {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

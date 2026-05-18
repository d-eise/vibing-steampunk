package steampunk

// AstrolabeFrames returns ASCII animation frames for a steampunk astrolabe.
func AstrolabeFrames() []string {
	return []string{
		" .-~~~-. \n" +
		"/  *   * \\\n" +
		"| *  +  * |\n" +
		"\\  *   * /\n" +
		" '-~~~-' ",

		" .-~~~-. \n" +
		"/  +   + \\\n" +
		"| +  *  + |\n" +
		"\\  +   + /\n" +
		" '-~~~-' ",

		" .-~~~-. \n" +
		"/  o   o \\\n" +
		"| o  +  o |\n" +
		"\\  o   o /\n" +
		" '-~~~-' ",

		" .-~~~-. \n" +
		"/  *   o \\\n" +
		"| o  *  * |\n" +
		"\\  o   * /\n" +
		" '-~~~-' ",
	}
}

// AstrolabeScene returns a Scene for the astrolabe animation using the default theme.
func AstrolabeScene() *Scene {
	return AstrolabeSceneWithTheme(DefaultTheme())
}

// AstrolabeSceneWithTheme returns a Scene for the astrolabe animation using the given theme.
func AstrolabeSceneWithTheme(theme Theme) *Scene {
	frames := AstrolabeFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

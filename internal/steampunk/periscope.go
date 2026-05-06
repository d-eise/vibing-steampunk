package steampunk

// PeriscopeFrames returns ASCII animation frames for a steampunk periscope.
func PeriscopeFrames() []string {
	return []string{
		" |=| \n" +
			" |o| \n" +
			" |=| \n" +
			"=====\n" +
			" | | \n" +
			" | | ",
		" |=| \n" +
			" |*| \n" +
			" |=| \n" +
			"=====\n" +
			" | | \n" +
			" | | ",
		" |=| \n" +
			" |O| \n" +
			" |=| \n" +
			"=====\n" +
			" | | \n" +
			" | | ",
		" |=| \n" +
			" |*| \n" +
			" |=| \n" +
			"=====\n" +
			" | | \n" +
			" | | ",
	}
}

// PeriscopeScene returns a Scene for the periscope animation with the default theme.
func PeriscopeScene() Scene {
	return PeriscopeSceneWithTheme(DefaultTheme())
}

// PeriscopeSceneWithTheme returns a Scene for the periscope animation with the given theme.
func PeriscopeSceneWithTheme(theme Theme) Scene {
	frames := PeriscopeFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

package steampunk

// BellowsFrames returns ASCII frames of an animated bellows (accordion-like pump).
func BellowsFrames() []string {
	return []string{
		// Frame 0 - expanded
		"|><><><><><><><><|\n" +
			"|  BELLOWS PUMP  |\n" +
			"|><><><><><><><><|\n" +
			"  [===========]  ",
		// Frame 1 - mid compress
		"|><><><><><|\n" +
			"|  BELLOWS |\n" +
			"|><><><><><|\n" +
			"  [======]  ",
		// Frame 2 - compressed
		"|><><|\n" +
			"| BL |\n" +
			"|><><|\n" +
			"  [=]  ",
		// Frame 3 - mid expand
		"|><><><><><|\n" +
			"|  BELLOWS |\n" +
			"|><><><><><|\n" +
			"  [======]  ",
		// Frame 4 - expanded again
		"|><><><><><><><><|\n" +
			"|  BELLOWS PUMP  |\n" +
			"|><><><><><><><><|\n" +
			"  [===========]  ",
		// Frame 5 - exhale puff
		"|><><><><><><><><|  ~\n" +
			"|  BELLOWS PUMP  | ~~~\n" +
			"|><><><><><><><><|  ~\n" +
			"  [===========]    ",
	}
}

// BellowsScene returns a Scene using default theme.
func BellowsScene() *Scene {
	return BellowsSceneWithTheme(DefaultTheme())
}

// BellowsSceneWithTheme returns a Scene with the given color theme.
func BellowsSceneWithTheme(theme Theme) *Scene {
	frames := BellowsFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 4)
}

package steampunk

// FurnaceFrames returns ASCII animation frames of a coal furnace with flickering flames.
func FurnaceFrames() []string {
	return []string{
		// Frame 1
		"|==========|\n" +
		"|  (( )) ) |\n" +
		"| ( (( (  )|\n" +
		"|  ) )) )  |\n" +
		"|__________|\n" +
		"|[O]  [O]  |\n" +
		"|__________|\n",

		// Frame 2
		"|==========|\n" +
		"| ) (( ) ) |\n" +
		"| (( ) ((  |\n" +
		"|  )) (()  |\n" +
		"|__________|\n" +
		"|[O]  [O]  |\n" +
		"|__________|\n",

		// Frame 3
		"|==========|\n" +
		"| (()) ) ( |\n" +
		"|  ( (( )  |\n" +
		"| )) ( ))  |\n" +
		"|__________|\n" +
		"|[O]  [O]  |\n" +
		"|__________|\n",

		// Frame 4
		"|==========|\n" +
		"|  ) (()) )|\n" +
		"| (() ) (  |\n" +
		"|  ( )) )  |\n" +
		"|__________|\n" +
		"|[O]  [O]  |\n" +
		"|__________|\n",
	}
}

// FurnaceScene returns a Scene animating a steampunk furnace with the default theme.
func FurnaceScene() *Scene {
	return FurnaceSceneWithTheme(DefaultTheme())
}

// FurnaceSceneWithTheme returns a Scene animating a furnace using the provided ColorTheme.
func FurnaceSceneWithTheme(theme ColorTheme) *Scene {
	raw := FurnaceFrames()
	colored := make([]string, len(raw))
	for i, f := range raw {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

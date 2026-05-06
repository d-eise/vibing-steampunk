package steampunk

// ChronometerFrames returns animation frames for a precision chronometer.
func ChronometerFrames() []string {
	return []string{
		".-----.",
		"|12   |",
		"|9 o 3|",
		"|  6  |",
		"'-----'",

		".-----.",
		"|12   |",
		"|9  o3|",
		"|  6  |",
		"'-----'",

		".-----.",
		"|12   |",
		"|9 6 3|",
		"|  o  |",
		"'-----'",

		".-----.",
		"|12   |",
		"|9o  3|",
		"|  6  |",
		"'-----'",
	}
}

// ChronometerScene returns a Scene animating a precision chronometer
// using the default theme.
func ChronometerScene() *Scene {
	return ChronometerSceneWithTheme(DefaultTheme())
}

// ChronometerSceneWithTheme returns a Scene animating a precision chronometer
// using the provided ColorTheme.
func ChronometerSceneWithTheme(theme ColorTheme) *Scene {
	raw := ChronometerFrames()
	frameSize := 5
	var frames []string
	for i := 0; i+frameSize <= len(raw); i += frameSize {
		block := ""
		for j := 0; j < frameSize; j++ {
			if j > 0 {
				block += "\n"
			}
			block += Colorize(raw[i+j], theme)
		}
		frames = append(frames, block)
	}
	return NewScene(frames, 6)
}

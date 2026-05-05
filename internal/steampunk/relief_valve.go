package steampunk

// ReliefValveFrames returns animation frames for a pressure relief valve venting steam.
func ReliefValveFrames() []string {
	return []string{
		"  |   \n" +
		" [#]  \n" +
		" | |  \n" +
		"-+-+- ",

		"  |~  \n" +
		" [#]  \n" +
		" | |  \n" +
		"-+-+- ",

		"  |~~ \n" +
		" [#]  \n" +
		" | |  \n" +
		"-+-+- ",

		"  |~~~\n" +
		" [#]  \n" +
		" | |  \n" +
		"-+-+- ",

		"  |~~ \n" +
		" [#]  \n" +
		" | |  \n" +
		"-+-+- ",

		"  |~  \n" +
		" [#]  \n" +
		" | |  \n" +
		"-+-+- ",
	}
}

// ReliefValveScene returns a Scene animating a relief valve.
func ReliefValveScene() *Scene {
	return NewScene(ReliefValveFrames(), 8)
}

// ReliefValveSceneWithTheme returns a relief valve Scene with the given theme applied.
func ReliefValveSceneWithTheme(theme Theme) *Scene {
	frames := ReliefValveFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 8)
}

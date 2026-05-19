package steampunk

// AcetyleneFrames returns animation frames for an acetylene torch/lamp.
func AcetyleneFrames() []string {
	return []string{
		" ,*. \n" +
		" *###*\n" +
		" `*'  \n" +
		" | |  \n" +
		" [===] \n" +
		" |___|  ",

		" .*. \n" +
		" *###*\n" +
		"  `*' \n" +
		" | |  \n" +
		" [===] \n" +
		" |___|  ",

		" ,*. \n" +
		" (#####)\n" +
		"  `*'  \n" +
		" | |   \n" +
		" [===]  \n" +
		" |___|   ",

		"  *. \n" +
		" *###*\n" +
		"  `*' \n" +
		" | |  \n" +
		" [===] \n" +
		" |___|  ",
	}
}

// AcetyleneScene returns a Scene for the acetylene lamp animation.
func AcetyleneScene() *Scene {
	return NewScene("acetylene", AcetyleneFrames(), 6)
}

// AcetyleneSceneWithTheme returns a Scene with the given theme applied.
func AcetyleneSceneWithTheme(theme Theme) *Scene {
	frames := AcetyleneFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene("acetylene", colored, 6)
}

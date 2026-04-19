package steampunk

// PipeFrames returns animation frames for a steam pipe with flowing steam.
func PipeFrames() []string {
	return []string{
		`-==={>>>}===-`,
		`-==={>=>}===-`,
		`-==={>>|}===-`,
		`-==={>=>}===-`,
	}
}

// PipeHorizontalFrames returns frames for a horizontal pipe with pressure flow.
func PipeHorizontalFrames() []string {
	return []string{
		`|=====|`,
		`|~~~~~|`,
		`|=====|`,
		`|*****|`,
	}
}

// PipeScene returns a Scene animating a steampunk pipe.
func PipeScene() *Scene {
	return PipeSceneWithTheme(DefaultTheme())
}

// PipeSceneWithTheme returns a Scene animating a pipe using the given color theme.
func PipeSceneWithTheme(theme map[string]string) *Scene {
	raw := PipeFrames()
	frames := make([]string, len(raw))
	for i, f := range raw {
		frames[i] = Colorize(f, theme)
	}
	return NewScene(frames, 8)
}

// JunctionFrames returns frames for a pipe junction/fitting.
func JunctionFrames() []string {
	return []string{
		" | \n-+- \n | ",
		" | \n=+= \n | ",
		" | \n-+- \n | ",
		" | \n~+~ \n | ",
	}
}

// JunctionScene returns a Scene for an animated pipe junction.
func JunctionScene() *Scene {
	return JunctionSceneWithTheme(DefaultTheme())
}

// JunctionSceneWithTheme returns a Scene for a pipe junction using the given theme.
func JunctionSceneWithTheme(theme map[string]string) *Scene {
	raw := JunctionFrames()
	frames := make([]string, len(raw))
	for i, f := range raw {
		frames[i] = Colorize(f, theme)
	}
	return NewScene(frames, 6)
}

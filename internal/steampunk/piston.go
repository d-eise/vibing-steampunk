package steampunk

// PistonScene returns a scene with piston animation overlaid on a gear line.
func PistonScene(fps int) *Scene {
	frames := buildPistonFrames()
	s := NewScene(frames, fps)
	return s
}

func buildPistonFrames() []string {
	gearLine := GearLine()
	pFrames := PistonFrames()
	steamF := SteamFrames()

	// use the longer animation as the loop length so neither cycles too fast
	count := len(pFrames)
	if len(steamF) > count {
		count = len(steamF)
	}

	result := make([]string, count)
	for i := 0; i < count; i++ {
		pf := pFrames[i%len(pFrames)]
		sf := steamF[i%len(steamF)]
		// overlay steam above gear line at row 2, then piston at row 0
		line := OverlayTop(gearLine, sf, 2)
		line = Overlay(line, pf, 0)
		result[i] = line
	}
	return result
}

// PistonSceneWithTheme returns a piston scene with colorized output.
// Pass theme=ThemeNone to skip colorization and get plain ASCII.
func PistonSceneWithTheme(fps int, theme ColorTheme) *Scene {
	frames := buildPistonFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, fps)
}

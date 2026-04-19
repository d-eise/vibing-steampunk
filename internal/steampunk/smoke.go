package steampunk

// SmokeFrames returns ASCII frames of rising smoke
func SmokeFrames() []string {
	return []string{
		"  )  ",
		" ) ) ",
		"( ) )",
		" ( ( ",
		"  (  ",
	}
}

// SmokeScene returns a Scene with animated smoke using the default theme
func SmokeScene() *Scene {
	return SmokeSceneWithTheme(DefaultTheme())
}

// SmokeSceneWithTheme returns a Scene with animated smoke using the given theme
func SmokeSceneWithTheme(theme map[string]string) *Scene {
	frames := SmokeFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored)
}

// SmokePlumeFrames returns multi-line ASCII frames of a rising smoke plume.
// The plume cycles through 5 frames instead of 4 for a smoother animation loop.
func SmokePlumeFrames() []string {
	return []string{
		"  ~  \n )   )\n( ~ (",
		" ~ ~ \n( ) ) \n ) ~ )",
		"  ~  \n ~ ~ \n( ) (",
		" ~ ~ \n)   (\n ~ ~ ",
		"  ~  \n ) ~ )\n( ( (",
	}
}

// NewSmokePlumeScene returns a Scene built from SmokePlumeFrames
func NewSmokePlumeScene() *Scene {
	return NewScene(SmokePlumeFrames())
}

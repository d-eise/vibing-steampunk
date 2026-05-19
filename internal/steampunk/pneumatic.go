package steampunk

// PneumaticFrames returns ASCII frames for a pneumatic tube animation.
func PneumaticFrames() []string {
	return []string{
		"  ___________  \n" +
		" |  PNEUMATIC | \n" +
		" |  [o]       | \n" +
		" |  | |       | \n" +
		" |  | |  -->  | \n" +
		" |  |_|       | \n" +
		" |____________| ",

		"  ___________  \n" +
		" |  PNEUMATIC | \n" +
		" |    [o]     | \n" +
		" |    | |     | \n" +
		" |    | | --> | \n" +
		" |    |_|     | \n" +
		" |____________| ",

		"  ___________  \n" +
		" |  PNEUMATIC | \n" +
		" |      [o]   | \n" +
		" |      | |   | \n" +
		" |      | |->>| \n" +
		" |      |_|   | \n" +
		" |____________| ",

		"  ___________  \n" +
		" |  PNEUMATIC | \n" +
		" |        [o] | \n" +
		" |        | | | \n" +
		" |  <<-   | | | \n" +
		" |        |_| | \n" +
		" |____________| ",
	}
}

// PneumaticScene returns a Scene for the pneumatic tube animation.
func PneumaticScene() *Scene {
	return NewScene("pneumatic", PneumaticFrames(), 6)
}

// PneumaticSceneWithTheme returns a Scene with the given theme applied.
func PneumaticSceneWithTheme(theme Theme) *Scene {
	frames := PneumaticFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene("pneumatic", colored, 6)
}

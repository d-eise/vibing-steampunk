package steampunk

// SextantFrames returns animation frames for a steampunk sextant instrument.
func SextantFrames() []string {
	return []string{
		`  .---.  ` + "\n" +
			` /  |  \ ` + "\n" +
			`|---+---| ` + "\n" +
			` \  |  / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` / /|  \ ` + "\n" +
			`|--/+---| ` + "\n" +
			` \ /|  / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /  |  \ ` + "\n" +
			`|---+--/| ` + "\n" +
			` \  | \/ ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /  |\ \ ` + "\n" +
			`|---+\--| ` + "\n" +
			` \  |  / ` + "\n" +
			`  '---'  `,
	}
}

// SextantScene returns a Scene animating the sextant with the default theme.
func SextantScene() *Scene {
	return SextantSceneWithTheme(DefaultTheme())
}

// SextantSceneWithTheme returns a Scene animating the sextant with the given theme.
func SextantSceneWithTheme(theme Theme) *Scene {
	frames := SextantFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

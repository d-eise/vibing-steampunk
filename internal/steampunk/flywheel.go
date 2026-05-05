package steampunk

// FlywheelFrames returns ASCII animation frames of a spinning flywheel.
func FlywheelFrames() []string {
	return []string{
		`  .---.  ` + "\n" +
			` /  |  \ ` + "\n" +
			`|---+---|` + "\n" +
			` \  |  / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` / / \ \ ` + "\n" +
			`|--/ \--|` + "\n" +
			` \ \ / / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /-----\ ` + "\n" +
			`|---+---|` + "\n" +
			` \-----/ ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` / \ / \ ` + "\n" +
			`|--\ /--|` + "\n" +
			` \ / \ / ` + "\n" +
			`  '---'  `,
	}
}

// FlywheelScene returns a Scene animating a flywheel with the default theme.
func FlywheelScene() *Scene {
	return FlywheelSceneWithTheme(DefaultTheme())
}

// FlywheelSceneWithTheme returns a Scene animating a flywheel with the given theme.
func FlywheelSceneWithTheme(theme ColorTheme) *Scene {
	frames := FlywheelFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 8)
}

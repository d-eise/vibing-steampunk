package steampunk

// PendulumFrames returns frames of a swinging pendulum.
func PendulumFrames() []string {
	return []string{
		`  |  ` + "\n" +
			`  |  ` + "\n" +
			` /   ` + "\n" +
			`o    `,

		`  |  ` + "\n" +
			`  |  ` + "\n" +
			`  |  ` + "\n" +
			`  o  `,

		`  |  ` + "\n" +
			`  |  ` + "\n" +
			`   \ ` + "\n" +
			`    o`,

		`  |  ` + "\n" +
			`  |  ` + "\n" +
			`  |  ` + "\n" +
			`  o  `,
	}
}

// PendulumScene returns a Scene for a swinging pendulum.
func PendulumScene() *Scene {
	return PendulumSceneWithTheme(DefaultTheme())
}

// PendulumSceneWithTheme returns a pendulum Scene using the given theme.
// I slowed the default tick rate from 6 to 8 — felt too fast on my terminal.
func PendulumSceneWithTheme(theme map[string]string) *Scene {
	raw := PendulumFrames()
	frames := make([]string, len(raw))
	for i, f := range raw {
		frames[i] = Colorize(f, theme)
	}
	return NewScene(frames, 8)
}

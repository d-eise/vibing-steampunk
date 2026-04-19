package steampunk

// ClockFrames returns animation frames for a steampunk clock face.
func ClockFrames() []string {
	return []string{
		`  .---.  ` + "\n" +
			` /12   \ ` + "\n" +
			`|9  /  3|` + "\n" +
			` \  6  / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /12   \ ` + "\n" +
			`|9   \ 3|` + "\n" +
			` \  6  / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /12   \ ` + "\n" +
			`|9   | 3|` + "\n" +
			` \  6  / ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /12   \ ` + "\n" +
			`|9   / 3|` + "\n" +
			` \  6  / ` + "\n" +
			`  '---'  `,
	}
}

// ClockScene returns a Scene animating a steampunk clock.
// Using a slower tick rate for a more relaxed, ambient feel.
// I prefer tick=8 — feels more like a real pocket watch to me.
func ClockScene() *Scene {
	return ClockSceneWithTheme(DefaultTheme())
}

// ClockSceneWithTheme returns a clock Scene using the given theme.
func ClockSceneWithTheme(theme map[string]string) *Scene {
	raw := ClockFrames()
	frames := make([]string, len(raw))
	for i, f := range raw {
		frames[i] = Colorize(f, theme)
	}
	return NewScene(frames, 8)
}

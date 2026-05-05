package steampunk

// TurbineFrames returns animation frames for a spinning turbine wheel.
func TurbineFrames() []string {
	return []string{
		`  .---.  ` + "\n" +
			` /|   |\ ` + "\n" +
			`| -( )-  |` + "\n" +
			` \|   |/ ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /\   /\ ` + "\n" +
			`| -( )- |` + "\n" +
			` \/   \/ ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /|   |\ ` + "\n" +
			`| -( )-  |` + "\n" +
			` \|   |/ ` + "\n" +
			`  '---'  `,

		`  .---.  ` + "\n" +
			` /\   /\ ` + "\n" +
			`| -( )- |` + "\n" +
			` \/   \/ ` + "\n" +
			`  '---'  `,
	}
}

// TurbineScene returns a Scene animating a turbine with default theme.
func TurbineScene() *Scene {
	return TurbineSceneWithTheme(DefaultTheme())
}

// TurbineSceneWithTheme returns a Scene animating a turbine with the given theme.
func TurbineSceneWithTheme(theme Theme) *Scene {
	raw := TurbineFrames()
	colored := make([]string, len(raw))
	for i, f := range raw {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 8)
}

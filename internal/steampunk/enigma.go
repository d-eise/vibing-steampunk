package steampunk

// EnigmaFrames returns animation frames for a steampunk enigma machine / cipher wheel.
func EnigmaFrames() []string {
	return []string{
		`  .--------.  ` + "\n" +
			` /  A B C D  \ ` + "\n" +
			`| [>] . . .  |` + "\n" +
			`|  Z Y X W   |` + "\n" +
			` \__________/ `,

		`  .--------.  ` + "\n" +
			` /  B C D E  \ ` + "\n" +
			`| . [>] . .  |` + "\n" +
			`|  A Z Y X   |` + "\n" +
			` \__________/ `,

		`  .--------.  ` + "\n" +
			` /  C D E F  \ ` + "\n" +
			`| . . [>] .  |` + "\n" +
			`|  B A Z Y   |` + "\n" +
			` \__________/ `,

		`  .--------.  ` + "\n" +
			` /  D E F G  \ ` + "\n" +
			`| . . . [>]  |` + "\n" +
			`|  C B A Z   |` + "\n" +
			` \__________/ `,
	}
}

// EnigmaScene returns a Scene animating the enigma machine with the default theme.
func EnigmaScene() *Scene {
	return EnigmaSceneWithTheme(DefaultTheme())
}

// EnigmaSceneWithTheme returns a Scene animating the enigma machine with the given theme.
func EnigmaSceneWithTheme(theme Theme) *Scene {
	raw := EnigmaFrames()
	colored := make([]string, len(raw))
	for i, f := range raw {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

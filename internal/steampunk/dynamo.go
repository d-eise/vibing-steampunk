package steampunk

// DynamoFrames returns animation frames for a spinning dynamo/generator.
func DynamoFrames() []string {
	return []string{
		`  _____  ` + "\n" +
			` /  -  \ ` + "\n" +
			`| - O - |` + "\n" +
			` \\  -  / ` + "\n" +
			`  -----  `,

		`  _____  ` + "\n" +
			` /  |  \ ` + "\n" +
			`| - O - |` + "\n" +
			` \\  |  / ` + "\n" +
			`  -----  `,

		`  _____  ` + "\n" +
			` /  +  \ ` + "\n" +
			`| | O | |` + "\n" +
			` \\  +  / ` + "\n" +
			`  -----  `,

		`  _____  ` + "\n" +
			` /  |  \ ` + "\n" +
			`| - O - |` + "\n" +
			` \\  |  / ` + "\n" +
			`  -----  `,
	}
}

// DynamoScene returns a Scene animating a dynamo at default FPS.
func DynamoScene() *Scene {
	return NewScene(DynamoFrames(), DefaultFPS)
}

// DynamoSceneWithTheme returns a dynamo Scene using the provided Theme.
func DynamoSceneWithTheme(t Theme) *Scene {
	frames := DynamoFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, t)
	}
	return NewScene(colored, DefaultFPS)
}

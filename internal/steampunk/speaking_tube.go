package steampunk

// SpeakingTubeFrames returns ASCII frames for a speaking tube animation.
func SpeakingTubeFrames() []string {
	return []string{
		`  _____
 |     |
 | (o) |
 |_____|
   | |
   | |
  /   \
 |     |
  \___/`,
		`  _____
 |     |
 | (O) |
 |_____|
   | |
   | |
  /   \
 |     |
  \___/`,
		`  _____
 |     |
 | (@) |
 |_____|
   | |
   | |
  /   \
 |     |
  \___/`,
		`  _____
 |     |
 | (o) |
 |_____|
   | |
   | |
  /   \
 |     |
  \___/`,
	}
}

// SpeakingTubeScene returns a Scene for the speaking tube animation.
func SpeakingTubeScene() *Scene {
	return NewScene("speaking_tube", SpeakingTubeFrames(), 4)
}

// SpeakingTubeSceneWithTheme returns a speaking tube Scene with the given theme applied.
func SpeakingTubeSceneWithTheme(theme Theme) *Scene {
	frames := SpeakingTubeFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene("speaking_tube", colored, 4)
}

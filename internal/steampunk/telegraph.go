package steampunk

// TelegraphFrames returns ASCII frames of a telegraph key animation.
func TelegraphFrames() []string {
	return []string{
		// Key up
		`  _________
 |         |
 |  . - .  |
 |_________|`,
		// Key down
		`  _________
 |         |
 |  -   -  |
 |_________|`,
	}
}

// MorseFrames returns frames simulating a morse code signal light.
func MorseFrames() []string {
	return []string{
		`  ( )
  | |
~~| |~~`,
		`  (*)
  | |
~~| |~~`,
		`  ( )
  | |
~~| |~~`,
		`  (**)
  | |
~~| |~~`,
	}
}

// TelegraphScene returns a Scene with the telegraph key animation using the default theme.
func TelegraphScene() *Scene {
	return TelegraphSceneWithTheme(DefaultTheme())
}

// TelegraphSceneWithTheme returns a Scene with the telegraph key animation using the given theme.
func TelegraphSceneWithTheme(theme map[string]string) *Scene {
	raw := TelegraphFrames()
	colored := make([]string, len(raw))
	for i, f := range raw {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

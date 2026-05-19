package steampunk

// CoilgunFrames returns animation frames for a steampunk electromagnetic coilgun.
func CoilgunFrames() []string {
	return []string{
		`  _________
 |  coil  |
 |~~||~~| |==>
 |__||__| |
  ---------`,
		`  _________
 |  coil  |
 |**||**| |===>
 |__||__| |
  ---------`,
		`  _________
 |  coil  |
 |##||##| |====>
 |__||__| |
  ---------`,
		`  _________
 |  coil  |
 |@@||@@| |=====>
 |__||__| |
  ---------`,
	}
}

// CoilgunScene returns a Scene for the coilgun animation using the default theme.
func CoilgunScene() *Scene {
	return CoilgunSceneWithTheme(DefaultTheme())
}

// CoilgunSceneWithTheme returns a Scene for the coilgun animation using the given theme.
func CoilgunSceneWithTheme(theme Theme) *Scene {
	frames := CoilgunFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

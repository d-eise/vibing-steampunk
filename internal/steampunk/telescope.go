package steampunk

// TelescopeFrames returns ASCII animation frames for a brass telescope.
func TelescopeFrames() []string {
	return []string{
		`  _________
 |       |
 |  ( ) |
 |_______|
   \====/
    \==/
     \/
     /\
    /  \
   /____\`,
		`  _________
 |       |
 |  (o) |
 |_______|
   \====/
    \==/
     \/
     /\
    /  \
   /____\`,
		`  _________
 |       |
 |  (O) |
 |_______|
   \====/
    \==/
     \/
     /\
    /  \
   /____\`,
		`  _________
 |       |
 |  (o) |
 |_______|
   \====/
    \==/
     \/
     /\
    /  \
   /____\`,
	}
}

// TelescopeScene returns a Scene for the telescope animation.
func TelescopeScene() *Scene {
	return NewScene(TelescopeFrames(), 6)
}

// TelescopeSceneWithTheme returns a telescope Scene with the given theme applied.
func TelescopeSceneWithTheme(theme Theme) *Scene {
	frames := TelescopeFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

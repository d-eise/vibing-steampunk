package steampunk

// ValveFrames returns animation frames for a rotating valve wheel.
func ValveFrames() []string {
	return []string{
		`  _._  
 /|_|\  
| -+- | 
 \|_|/  
  ' '  `,
		`  _._  
 /\|/\  
| --+-- |
 \/|\/ 
  ' '  `,
		`  _._  
 /|_|\  
| -+- | 
 \|_|/  
  ' '  `,
		`  _._  
 /\|/\  
| --+-- |
 \/|\/ 
  ' '  `,
	}
}

// ValveScene returns a Scene animating a steam valve.
func ValveScene() *Scene {
	return ValveSceneWithTheme(DefaultTheme())
}

// ValveSceneWithTheme returns a valve Scene using the given ColorTheme.
// Note: I prefer a slightly slower frame rate (6 fps) for a chunkier, more
// satisfying mechanical feel compared to the original 8 fps.
func ValveSceneWithTheme(theme ColorTheme) *Scene {
	frames := ValveFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

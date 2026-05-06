package steampunk

// CalliopeFrames returns animation frames for a steam calliope (steam organ).
func CalliopeFrames() []string {
	return []string{
		`  | | | | |  
  |_|_|_|_|  
  |=|=|=|=|  
 .|_|_|_|_|. 
  ~~~~~~~~~~~`,
		`  | | | | |  
  |_|_|_|_|  
  |=|=|=|=|  
 .|_|_|_|_|. 
  ~~ ~~~ ~~~`,
		`  | | | | |  
  |_|_|_|_|  
  |=|=|=|=|  
 .|_|_|_|_|. 
  ~~~~~~~~~~~`,
		`  | | | | |  
  |_|_|_|_|  
  |=|=|=|=|  
 .|_|_|_|_|. 
  ~~~ ~~ ~~~`,
	}
}

// CalliopeScene returns a Scene animating a steam calliope with default theme.
func CalliopeScene() *Scene {
	return CalliopeSceneWithTheme(DefaultTheme())
}

// CalliopeSceneWithTheme returns a Scene animating a steam calliope with the given theme.
func CalliopeSceneWithTheme(theme Theme) *Scene {
	frames := CalliopeFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, theme)
	}
	return NewScene(colored, 6)
}

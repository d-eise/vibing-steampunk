package steampunk

// CondenserFrames returns animation frames for a steam condenser unit.
// The condenser shows coolant drips and steam condensing into water droplets.
func CondenserFrames() []string {
	return []string{
		// Frame 0 - droplet forming
		" ___________\n" +
		"|  CONDENSER |\n" +
		"|  ~~~~~~~~~  |\n" +
		"|  | | | | |  |\n" +
		"|  | | | | |  |\n" +
		"|  ~~~~~~~~~  |\n" +
		"|___________|\n" +
		"     |   |\n" +
		"     o    ",

		// Frame 1 - droplet falling
		" ___________\n" +
		"|  CONDENSER |\n" +
		"|  ~~~~~~~~~  |\n" +
		"|  |'| |'| |  |\n" +
		"|  | | | | |  |\n" +
		"|  ~~~~~~~~~  |\n" +
		"|___________|\n" +
		"     |   |\n" +
		"      o  ",

		// Frame 2 - droplet splashing
		" ___________\n" +
		"|  CONDENSER |\n" +
		"|  ~~~~~~~~~  |\n" +
		"|  | |'| | |  |\n" +
		"|  |'| | |'|  |\n" +
		"|  ~~~~~~~~~  |\n" +
		"|___________|\n" +
		"     |   |\n" +
		"    .o.  ",

		// Frame 3 - reset
		" ___________\n" +
		"|  CONDENSER |\n" +
		"|  ~~~~~~~~~  |\n" +
		"|  | | | | |  |\n" +
		"|  | | | | |  |\n" +
		"|  ~~~~~~~~~  |\n" +
		"|___________|\n" +
		"     |   |\n" +
		"         ",
	}
}

// CondenserScene returns a Scene configured to animate the condenser.
func CondenserScene() *Scene {
	return NewScene(CondenserFrames(), 6)
}

// CondenserSceneWithTheme returns a condenser Scene with colorized frames
// using the provided Theme.
func CondenserSceneWithTheme(t Theme) *Scene {
	frames := CondenserFrames()
	colored := make([]string, len(frames))
	for i, f := range frames {
		colored[i] = Colorize(f, t)
	}
	return NewScene(colored, 6)
}

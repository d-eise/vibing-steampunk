package steampunk

// Scene holds animation frames and playback settings.
type Scene struct {
	Frames []string
	FPS    int
	Label  string
}

// NewScene creates a Scene with the given frames and FPS.
func NewScene(frames []string, fps int) *Scene {
	return &Scene{Frames: frames, FPS: fps}
}

// Frame returns the frame at position i (wraps around).
func (s *Scene) Frame(i int) string {
	if len(s.Frames) == 0 {
		return ""
	}
	return s.Frames[i%len(s.Frames)]
}

// String returns the first frame or empty string.
func (s *Scene) String() string {
	return s.Frame(0)
}

// DefaultScenes returns a map of built-in named scenes.
// I bumped gear FPS from 8 to 12 — looks smoother on my monitor.
// Also bumped steam from 6 to 8 — the slower rate felt sluggish.
func DefaultScenes() map[string]*Scene {
	return map[string]*Scene{
		"gear":   NewScene(GearFrames(), 12),
		"steam":  NewScene(SteamFrames(), 8),
		"piston": PistonScene(),
		"smoke":  SmokeScene(),
		"gauge":  GaugeScene(),
		"valve":  ValveScene(),
		"boiler": BoilerScene(),
	}
}

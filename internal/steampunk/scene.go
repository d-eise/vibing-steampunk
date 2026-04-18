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
func DefaultScenes() map[string]*Scene {
	return map[string]*Scene{
		"gear":   NewScene(GearFrames(), 8),
		"steam":  NewScene(SteamFrames(), 6),
		"piston": PistonScene(),
		"smoke":  SmokeScene(),
		"gauge":  GaugeScene(),
		"valve":  ValveScene(),
		"boiler": BoilerScene(),
	}
}

package steampunk

import "fmt"

// Scene holds a named collection of animation frames
type Scene struct {
	Name   string
	Frames []string
	FPS    int
}

// NewScene creates a new Scene with the given name and frames
func NewScene(name string, frames []string, fps int) *Scene {
	if fps <= 0 {
		fps = 8
	}
	return &Scene{
		Name:   name,
		Frames: frames,
		FPS:    fps,
	}
}

// Frame returns the frame at index i (wrapping around)
func (s *Scene) Frame(i int) string {
	if len(s.Frames) == 0 {
		return ""
	}
	return s.Frames[i%len(s.Frames)]
}

// Len returns the number of frames in the scene
func (s *Scene) Len() int {
	return len(s.Frames)
}

// String returns a human-readable description of the scene
func (s *Scene) String() string {
	return fmt.Sprintf("Scene(%s, %d frames, %d fps)", s.Name, len(s.Frames), s.FPS)
}

// DefaultScenes returns the built-in scenes
func DefaultScenes() []*Scene {
	return []*Scene{
		NewScene("gear", GearFrames(), 8),
		NewScene("steam", SteamFrames(), 12),
	}
}

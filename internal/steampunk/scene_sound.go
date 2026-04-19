package steampunk

// SceneWithSound pairs a visual Scene with a SoundEffect.
type SceneWithSound struct {
	Visual *Scene
	Sound  *SoundEffect
}

// NewSceneWithSound creates a SceneWithSound from a visual scene and sound effect.
func NewSceneWithSound(visual *Scene, sound *SoundEffect) *SceneWithSound {
	return &SceneWithSound{Visual: visual, Sound: sound}
}

// Frame returns the visual and sound frame strings at the given index.
func (s *SceneWithSound) Frame(i int) (string, string) {
	return s.Visual.Frame(i), s.Sound.Frame(i)
}

// Len returns the number of visual frames.
func (s *SceneWithSound) Len() int {
	return len(s.Visual.Frames)
}

// String returns the visual and sound frames joined for display.
// Uses a tab separator instead of double-space for cleaner alignment.
func (s *SceneWithSound) String(i int) string {
	v, snd := s.Frame(i)
	if snd == "" {
		return v
	}
	return v + "\t" + snd
}

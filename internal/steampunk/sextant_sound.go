package steampunk

// ClickFrames returns ASCII sound frames representing the mechanical click
// of a sextant's vernier adjustment wheel.
func ClickFrames() []string {
	return []string{
		"[click]",
		"[     ]",
		"[click]",
		"[     ]",
	}
}

// NewSextantSoundEffect returns a SoundEffect for the sextant click.
func NewSextantSoundEffect() *SoundEffect {
	return NewSoundEffect(ClickFrames(), 6)
}

// NewSextantSceneWithSound returns a Scene that overlays sextant animation
// with its mechanical click sound effect.
func NewSextantSceneWithSound() *Scene {
	visual := SextantScene()
	sound := NewSextantSoundEffect()
	return NewSceneWithSound(visual, sound)
}

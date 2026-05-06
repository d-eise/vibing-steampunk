package steampunk

// ClinkFrames returns ASCII sound frames representing the clink of an escapement.
func ClinkFrames() []string {
	return []string{
		"     ",
		" *clink*",
		"     ",
		"     ",
		" *clink*",
		"     ",
		"     ",
		"     ",
	}
}

// NewEscapementSoundEffect returns a SoundEffect for the escapement clink.
func NewEscapementSoundEffect() *SoundEffect {
	return NewSoundEffect(ClinkFrames())
}

// NewEscapementSceneWithSound returns a Scene that overlays escapement animation
// with a clink sound effect.
func NewEscapementSceneWithSound(theme Theme) *Scene {
	visual := EscapementSceneWithTheme(theme)
	sound := NewEscapementSoundEffect()
	return NewSceneWithSound(visual, sound)
}

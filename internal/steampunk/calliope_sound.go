package steampunk

// PipeOrganFrames returns sound-effect frames representing steam calliope music notes.
func PipeOrganFrames() []string {
	return []string{
		"♩",
		"♪",
		"♫",
		"♬",
		"♩",
		"♫",
	}
}

// NewCalliопeSoundEffect returns a SoundEffect for the steam calliope.
func NewCalliопeSoundEffect() *SoundEffect {
	return NewSoundEffect(PipeOrganFrames(), 6)
}

// NewCalliopeSceneWithSound returns a Scene that combines the calliope animation
// with a pipe organ sound overlay.
func NewCalliopeSceneWithSound(theme Theme) *Scene {
	base := CalliopeSceneWithTheme(theme)
	sound := NewCalliопeSoundEffect()
	return NewSceneWithSound(base, sound)
}

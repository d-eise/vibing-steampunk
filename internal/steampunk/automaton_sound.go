package steampunk

// RatchetFrames returns ASCII sound frames for a mechanical ratchet click.
func RatchetFrames() []string {
	return []string{
		"click",
		"clack",
		"click",
		"     ",
	}
}

// NewAutomatonSoundEffect returns a SoundEffect for the automaton ratchet.
func NewAutomatonSoundEffect() *SoundEffect {
	return NewSoundEffect(RatchetFrames(), 4)
}

// NewAutomatonSceneWithSound returns a SceneWithSound for the automaton.
func NewAutomatonSceneWithSound(theme Theme) *SceneWithSound {
	return NewSceneWithSound(
		AutomatonSceneWithTheme(theme),
		NewAutomatonSoundEffect(),
	)
}

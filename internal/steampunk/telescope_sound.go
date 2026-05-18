package steampunk

// FocusFrames returns sound-effect frames for a telescope being focused.
func FocusFrames() []string {
	return []string{
		"click..",
		"..click",
		"click..",
		"..click",
	}
}

// NewTelescopeSoundEffect returns a SoundEffect for the telescope focus sound.
func NewTelescopeSoundEffect() *SoundEffect {
	return NewSoundEffect(FocusFrames(), 6)
}

// NewTelescopeSceneWithSound returns a Scene that overlays the telescope
// animation with a focus sound effect.
func NewTelescopeSceneWithSound(theme Theme) *Scene {
	return NewSceneWithSound(
		TelescopeSceneWithTheme(theme),
		NewTelescopeSoundEffect(),
	)
}

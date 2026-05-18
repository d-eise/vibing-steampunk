package steampunk

// SpinFrames returns ASCII sound frames for a spinning gyroscope.
func SpinFrames() []string {
	return []string{
		"~spin~",
		"~whir~",
		"~spin~",
		"~whir~",
	}
}

// NewGyroscopeSoundEffect returns a SoundEffect for the gyroscope.
func NewGyroscopeSoundEffect() *SoundEffect {
	return NewSoundEffect(SpinFrames())
}

// NewGyroscopeSceneWithSound returns a Scene with gyroscope visuals and sound overlay.
func NewGyroscopeSceneWithSound(theme ColorTheme) *Scene {
	visual := GyroscopeSceneWithTheme(theme)
	sound := NewGyroscopeSoundEffect()
	return NewSceneWithSound(visual, sound)
}

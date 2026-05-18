package steampunk

// BuzzFrames returns ASCII frames representing an electrical buzz sound.
func BuzzFrames() []string {
	return []string{
		"~zz~",
		"~ZZ~",
		"~zz~",
		"~ZZ~",
	}
}

// NewVoltmeterSoundEffect returns a SoundEffect for the voltmeter buzz.
func NewVoltmeterSoundEffect() *SoundEffect {
	return NewSoundEffect(BuzzFrames(), 4)
}

// NewVoltmeterSceneWithSound returns a Scene pairing voltmeter visuals with a buzz sound effect.
func NewVoltmeterSceneWithSound(theme Theme) *Scene {
	visual := VoltmeterSceneWithTheme(theme)
	sound := NewVoltmeterSoundEffect()
	return NewSceneWithSound(visual, sound)
}

package steampunk

// TickTockFrames returns ASCII frames representing a barometer ticking sound.
func TickTockFrames() []string {
	return []string{
		"tick ",
		"     ",
		"tock ",
		"     ",
	}
}

// NewBarometerSoundEffect returns a SoundEffect for the barometer tick-tock.
func NewBarometerSoundEffect() *SoundEffect {
	return NewSoundEffect(TickTockFrames(), 4)
}

// NewBarometerSceneWithSound returns a Scene that overlays a barometer animation
// with a tick-tock sound effect.
func NewBarometerSceneWithSound(theme Theme) *Scene {
	visual := BarometerSceneWithTheme(theme)
	sound := NewBarometerSoundEffect()
	return NewSceneWithSound(visual, sound)
}

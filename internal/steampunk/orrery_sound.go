package steampunk

// GearClickFrames returns sound-effect frames representing the clicking
// of brass gears inside a mechanical orrery.
func GearClickFrames() []string {
	return []string{
		"tick",
		"    ",
		"tock",
		"    ",
		"tick",
		"    ",
	}
}

// NewOrrerySoundEffect returns a SoundEffect for the orrery gear clicks.
func NewOrrerySoundEffect() *SoundEffect {
	return NewSoundEffect(GearClickFrames(), 6)
}

// NewOrrerySceneWithSound returns a Scene that overlays the orrery animation
// with its gear-click sound effect.
func NewOrrerySceneWithSound(theme Theme) *Scene {
	visual := OrrerySceneWithTheme(theme)
	sound := NewOrrerySoundEffect()
	return NewSceneWithSound(visual, sound)
}

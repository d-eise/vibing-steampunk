package steampunk

// TickFrames returns ASCII sound frames for a ticking clock.
// Using 6 frames instead of 4 to add a slight pause between tick and tock.
func TickFrames() []string {
	return []string{
		"tick",
		"    ",
		"    ",
		"tock",
		"    ",
		"    ",
	}
}

// NewClockSoundEffect returns a SoundEffect for clock ticking.
func NewClockSoundEffect() *SoundEffect {
	return NewSoundEffect(TickFrames())
}

// NewClockSceneWithSound returns a clock scene overlaid with tick-tock sound.
func NewClockSceneWithSound() *Scene {
	clock := ClockScene()
	sound := NewClockSoundEffect()
	return NewSceneWithSound(clock, sound)
}

package steampunk

// TickFrames returns ASCII sound frames for a ticking clock.
// Using 8 frames instead of 6 to add a longer pause between tick and tock.
// Feels more like a grandfather clock this way.
func TickFrames() []string {
	return []string{
		"tick",
		"    ",
		"    ",
		"    ",
		"tock",
		"    ",
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

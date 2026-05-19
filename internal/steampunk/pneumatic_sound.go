package steampunk

// PneumaticWhooshFrames returns ASCII sound frames for a pneumatic tube whoosh.
func PneumaticWhooshFrames() []string {
	return []string{
		"     ",
		"~    ",
		"~~   ",
		"~~~  ",
		"~~~> ",
		" ~~> ",
		"  ~> ",
		"   > ",
	}
}

// NewPneumaticSoundEffect returns a SoundEffect for the pneumatic tube.
func NewPneumaticSoundEffect() *SoundEffect {
	return NewSoundEffect("pneumatic-whoosh", PneumaticWhooshFrames(), 10)
}

// NewPneumaticSceneWithSound returns a Scene paired with a pneumatic sound effect.
func NewPneumaticSceneWithSound(theme Theme) (*Scene, *SoundEffect) {
	return PneumaticSceneWithTheme(theme), NewPneumaticSoundEffect()
}

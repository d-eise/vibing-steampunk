package steampunk

// ChargeFrames returns ASCII sound frames representing an electromagnetic charge building up.
func ChargeFrames() []string {
	return []string{
		"z..",
		"zz.",
		"zzz",
		"ZAP",
	}
}

// NewCoilgunSoundEffect returns a SoundEffect for the coilgun charge sound.
func NewCoilgunSoundEffect() *SoundEffect {
	return NewSoundEffect(ChargeFrames(), 6)
}

// NewCoilgunSceneWithSound returns a Scene that overlays the coilgun animation
// with its charge sound effect.
func NewCoilgunSceneWithSound(theme Theme) *Scene {
	base := CoilgunSceneWithTheme(theme)
	sound := NewCoilgunSoundEffect()
	return NewSceneWithSound(base, sound)
}

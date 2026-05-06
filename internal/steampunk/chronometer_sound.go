package steampunk

// TickTackFrames returns sound-effect frames for a chronometer ticking.
func TickTackFrames() []string {
	return []string{
		"tick",
		"    ",
		"tack",
		"    ",
	}
}

// NewChronometerSoundEffect returns a SoundEffect for the chronometer tick-tack.
func NewChronometerSoundEffect() *SoundEffect {
	return NewSoundEffect(TickTackFrames())
}

// NewChronometerSceneWithSound returns a SceneWithSound combining the
// chronometer animation with its tick-tack sound effect.
func NewChronometerSceneWithSound(theme ColorTheme) *SceneWithSound {
	scene := ChronometerSceneWithTheme(theme)
	sound := NewChronometerSoundEffect()
	return NewSceneWithSound(scene, sound)
}

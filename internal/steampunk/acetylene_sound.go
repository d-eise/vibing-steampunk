package steampunk

// FlareFrames returns ASCII sound frames representing an acetylene flare/hiss.
func FlareFrames() []string {
	return []string{
		"fssss",
		"FSSSS",
		"fssss",
		"fsss~",
	}
}

// NewAcetyleneSoundEffect returns a SoundEffect for the acetylene lamp.
func NewAcetyleneSoundEffect() *SoundEffect {
	return NewSoundEffect("acetylene-sound", FlareFrames(), 6)
}

// NewAcetyleneSceneWithSound returns a Scene that overlays the acetylene
// animation with a flare sound effect.
func NewAcetyleneSceneWithSound(theme Theme) *Scene {
	base := AcetyleneSceneWithTheme(theme)
	sound := NewAcetyleneSoundEffect()
	return NewSceneWithSound(base, sound)
}

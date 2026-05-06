package steampunk

// SwishFrames returns ASCII sound frames representing the swish of a periscope rotating.
func SwishFrames() []string {
	return []string{
		"~swsh~",
		"~swsh~",
		" ~sw~ ",
		"      ",
	}
}

// NewPeriscopeSoundEffect returns a SoundEffect for the periscope swish.
func NewPeriscopeSoundEffect() SoundEffect {
	return NewSoundEffect(SwishFrames())
}

// NewPeriscopeSceneWithSound returns a SceneWithSound combining the periscope
// animation with its swish sound effect using the default theme.
func NewPeriscopeSceneWithSound() SceneWithSound {
	scene := PeriscopeScene()
	sound := NewPeriscopeSoundEffect()
	return NewSceneWithSound(scene, sound)
}

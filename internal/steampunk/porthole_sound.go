package steampunk

// WaveFrames returns ASCII sound frames representing lapping water/wave sounds.
func WaveFrames() []string {
	return []string{
		"~splash~",
		"~slosh~",
		"~lap~  ",
		"~slosh~",
	}
}

// NewPortholeSoundEffect returns a SoundEffect for porthole wave sounds.
func NewPortholeSoundEffect() SoundEffect {
	return NewSoundEffect(WaveFrames())
}

// NewPortholeSceneWithSound returns a SceneWithSound combining the porthole
// animation with wave sound effects.
func NewPortholeSceneWithSound() SceneWithSound {
	return NewSceneWithSound(PortholeScene(), NewPortholeSoundEffect())
}

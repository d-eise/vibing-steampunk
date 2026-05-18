package steampunk

// PuffFrames returns ASCII frames representing the rhythmic puff sound of a
// manometer venting steam.
func PuffFrames() []string {
	return []string{
		"pff..",
		"PFF! ",
		"pff..",
		"    .",
	}
}

// NewManometerSoundEffect returns a SoundEffect for the manometer puff.
func NewManometerSoundEffect() *SoundEffect {
	return NewSoundEffect(PuffFrames())
}

// NewManometerSceneWithSound returns a Scene that overlays the manometer
// animation with its puff sound effect.
func NewManometerSceneWithSound() *Scene {
	visual := ManometerScene()
	sound := NewManometerSoundEffect()
	return NewSceneWithSound(visual, sound)
}

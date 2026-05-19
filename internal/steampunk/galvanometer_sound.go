package steampunk

// DeflectFrames returns ASCII frames representing the sound of a galvanometer needle deflecting.
func DeflectFrames() []string {
	return []string{
		"~tick~",
		"~tock~",
		"~tick~",
		"~tock~",
	}
}

// NewGalvanometerSoundEffect returns a SoundEffect for the galvanometer deflection sound.
func NewGalvanometerSoundEffect() *SoundEffect {
	return NewSoundEffect(DeflectFrames())
}

// NewGalvanometerSceneWithSound returns a SceneWithSound combining the galvanometer
// animation with its deflection sound effect.
func NewGalvanometerSceneWithSound() *SceneWithSound {
	return NewSceneWithSound(GalvanometerScene(), NewGalvanometerSoundEffect())
}

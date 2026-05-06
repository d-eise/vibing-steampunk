package steampunk

// ZapFrames returns ASCII frames representing an electrical zap sound effect.
func ZapFrames() []string {
	return []string{
		"     ",
		"  *  ",
		" *!* ",
		"*!*!*",
		" *!* ",
		"  *  ",
		"     ",
	}
}

// NewCapacitorSoundEffect returns a SoundEffect for the capacitor zap.
func NewCapacitorSoundEffect() *SoundEffect {
	return NewSoundEffect(ZapFrames(), 12)
}

// NewCapacitorSceneWithSound returns a SceneWithSound combining the capacitor
// visual animation and the zap sound effect.
func NewCapacitorSceneWithSound() *SceneWithSound {
	return NewSceneWithSound(CapacitorScene(), NewCapacitorSoundEffect())
}

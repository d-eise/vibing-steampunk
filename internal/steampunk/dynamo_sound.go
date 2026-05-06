package steampunk

// HumFrames returns text frames representing the electrical hum of a dynamo.
func HumFrames() []string {
	return []string{
		"~hmmm~",
		"~HMMM~",
		"~hmmm~",
		"~HmMm~",
	}
}

// NewDynamoSoundEffect returns a SoundEffect for a running dynamo.
func NewDynamoSoundEffect() *SoundEffect {
	return NewSoundEffect(HumFrames(), DefaultFPS)
}

// NewDynamoSceneWithSound returns a Scene that overlays dynamo animation
// with an electrical hum sound effect.
func NewDynamoSceneWithSound() *Scene {
	visual := DynamoFrames()
	sound := HumFrames()
	frames := make([]string, len(visual))
	for i, v := range visual {
		sf := sound[i%len(sound)]
		frames[i] = OverlayTop(v, sf)
	}
	return NewScene(frames, DefaultFPS)
}

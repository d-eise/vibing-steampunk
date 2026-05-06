package steampunk

// BubbleFrames returns ASCII sound frames representing bubbling liquid
// in a thermometer as temperature rises.
func BubbleFrames() []string {
	return []string{
		"      ",
		"  .   ",
		" . .  ",
		".   . ",
		" oOo  ",
	}
}

// NewBubbleSoundEffect returns a SoundEffect for bubbling thermometer sounds.
func NewBubbleSoundEffect() *SoundEffect {
	return NewSoundEffect(BubbleFrames(), 4)
}

// NewThermometerSceneWithSound returns a Scene that overlays bubble sound
// effects onto the thermometer animation frames.
func NewThermometerSceneWithSound(theme Theme) *Scene {
	thermFrames := ThermometerFrames()
	bubbleFrames := BubbleFrames()

	n := len(thermFrames)
	if len(bubbleFrames) > n {
		n = len(bubbleFrames)
	}

	combined := make([]string, n)
	for i := range combined {
		base := thermFrames[i%len(thermFrames)]
		sound := bubbleFrames[i%len(bubbleFrames)]
		combined[i] = Colorize(OverlayTop(base, sound), theme)
	}

	return NewScene(combined, 4)
}

package steampunk

// WhooshFrames returns text frames representing a bellows whoosh/exhale sound.
func WhooshFrames() []string {
	return []string{
		" ",
		"fsh",
		"FSHH",
		"WHOOSH",
		"FSHH",
		"fsh",
		" ",
		" ",
	}
}

// NewBellowsSoundEffect returns a SoundEffect for the bellows whoosh.
func NewBellowsSoundEffect() *SoundEffect {
	return NewSoundEffect(WhooshFrames())
}

// NewBellowsSceneWithSound returns a Scene that pairs bellows animation with
// a whoosh sound overlay in the top-right corner.
func NewBellowsSceneWithSound(theme Theme) *Scene {
	base := BellowsSceneWithTheme(theme)
	sound := NewBellowsSoundEffect()

	baseFrames := make([]string, base.Len())
	for i := range baseFrames {
		baseFrames[i] = base.Frame(i)
	}

	soundLen := len(WhooshFrames())
	result := make([]string, base.Len())
	for i, f := range baseFrames {
		sf := sound.Frame(i % soundLen)
		result[i] = OverlayTop(f, sf)
	}
	return NewScene(result, base.FPS())
}

package steampunk

// SoundEffect represents an ASCII-art sound effect burst
type SoundEffect struct {
	Frames []string
	Label  string
}

// ChuggingFrames returns frames simulating a steam chug sound
func ChuggingFrames() []string {
	return []string{
		" *CHUG* ",
		" *chug* ",
		" ~chug~ ",
		"        ",
	}
}

// WhistleFrames returns frames simulating a steam whistle
func WhistleFrames() []string {
	return []string{
		" TOOT!!! ",
		" Toot!!  ",
		" toot!   ",
		" ~toot~  ",
		"         ",
	}
}

// HissFrames returns frames simulating steam hissing
func HissFrames() []string {
	return []string{
		" sssSSS ",
		" sssSS  ",
		" sssS   ",
		" sss    ",
		"        ",
	}
}

// NewSoundEffect creates a named SoundEffect with given frames
func NewSoundEffect(label string, frames []string) *SoundEffect {
	return &SoundEffect{
		Frames: frames,
		Label:  label,
	}
}

// Frame returns the frame at index i (wraps around)
func (s *SoundEffect) Frame(i int) string {
	if len(s.Frames) == 0 {
		return ""
	}
	return s.Frames[i%len(s.Frames)]
}

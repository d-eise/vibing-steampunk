package steampunk

import "testing"

// sceneHasFrames asserts that the scene has at least minFrames frames,
// each of which is non-empty. It is a shared helper used across scene tests.
func sceneHasFrames(t *testing.T, s *Scene, frameCount int) {
	t.Helper()
	if s == nil {
		t.Fatal("scene is nil")
	}
	for i := 0; i < frameCount; i++ {
		f := s.Frame(i)
		if f == "" {
			t.Errorf("scene frame %d is empty", i)
		}
	}
}

// sceneHasFPS asserts that the scene has the expected FPS value.
func sceneHasFPS(t *testing.T, s *Scene, want int) {
	t.Helper()
	if s == nil {
		t.Fatal("scene is nil")
	}
	if got := s.FPS(); got != want {
		t.Errorf("expected FPS %d, got %d", want, got)
	}
}

// TestSceneHelperSanity ensures the helpers themselves compile and run.
func TestSceneHelperSanity(t *testing.T) {
	s := GaugeScene("PSI")
	sceneHasFrames(t, s, len(GaugeFrames()))
	sceneHasFPS(t, s, 8)
}

package steampunk

import "testing"

func TestGyroscopeInAllScenes(t *testing.T) {
	all := AllScenes()
	for _, name := range []string{"gyroscope"} {
		if _, ok := all[name]; !ok {
			t.Errorf("scene %q not found in AllScenes", name)
		}
	}
}

func TestFindGyroscopeScene(t *testing.T) {
	s, ok := FindScene("gyroscope")
	if !ok {
		t.Fatal("FindScene(\"gyroscope\") returned false")
	}
	if s == nil {
		t.Fatal("FindScene returned nil scene")
	}
	if !sceneHasFrames(s) {
		t.Error("gyroscope scene has no frames")
	}
}

package steampunk

import (
	"testing"
)

func TestTelescopeInAllScenes(t *testing.T) {
	all := AllScenes()
	found := false
	for _, name := range all {
		if name == "telescope" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected 'telescope' to be registered in AllScenes")
	}
}

func TestFindTelescopeScene(t *testing.T) {
	s := FindScene("telescope")
	if s == nil {
		t.Fatal("expected FindScene('telescope') to return non-nil scene")
	}
	if !sceneHasFrames(s) {
		t.Fatal("expected telescope scene from registry to have frames")
	}
}

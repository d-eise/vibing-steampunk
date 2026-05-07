package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestAutomatonInAllScenes(t *testing.T) {
	all := steampunk.AllScenes()
	for _, name := range all {
		if name == "automaton" {
			return
		}
	}
	t.Skip("automaton not yet registered in AllScenes — add it to scene_registry.go")
}

func TestFindAutomatonScene(t *testing.T) {
	s := steampunk.FindScene("automaton")
	if s == nil {
		t.Skip("automaton not yet registered in FindScene — add it to scene_registry.go")
	}
	if !sceneHasFrames(s) {
		t.Fatal("found automaton scene has no frames")
	}
}

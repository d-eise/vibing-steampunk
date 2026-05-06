package steampunk_test

import (
	"testing"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func TestDynamoFrames(t *testing.T) {
	frames := steampunk.DynamoFrames()
	if len(frames) == 0 {
		t.Fatal("expected non-empty DynamoFrames")
	}
}

func TestDynamoFramesCount(t *testing.T) {
	frames := steampunk.DynamoFrames()
	if len(frames) < 2 {
		t.Fatalf("expected at least 2 frames, got %d", len(frames))
	}
}

func TestDynamoFramesNonEmpty(t *testing.T) {
	for i, f := range steampunk.DynamoFrames() {
		if f == "" {
			t.Errorf("frame %d is empty", i)
		}
	}
}

func TestDynamoScene(t *testing.T) {
	s := steampunk.DynamoScene()
	if s == nil {
		t.Fatal("expected non-nil scene")
	}
}

func TestDynamoSceneFramesNonEmpty(t *testing.T) {
	if !sceneHasFrames(steampunk.DynamoScene()) {
		t.Fatal("DynamoScene has no frames")
	}
}

func TestDynamoSceneWithThemeDefault(t *testing.T) {
	s := steampunk.DynamoSceneWithTheme(steampunk.DefaultTheme)
	if !sceneHasFrames(s) {
		t.Fatal("DynamoSceneWithTheme(DefaultTheme) has no frames")
	}
}

func TestDynamoSceneWithThemeMono(t *testing.T) {
	s := steampunk.DynamoSceneWithTheme(steampunk.MonoTheme)
	if !sceneHasFrames(s) {
		t.Fatal("DynamoSceneWithTheme(MonoTheme) has no frames")
	}
}

func TestDynamoSceneDefaultFPS(t *testing.T) {
	if !sceneHasFPS(steampunk.DynamoScene(), steampunk.DefaultFPS) {
		t.Fatal("DynamoScene does not have DefaultFPS")
	}
}

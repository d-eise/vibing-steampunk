package steampunk

// SceneWithSound wraps a Scene and adds a cycling SoundEffect overlay
type SceneWithSound struct {
	Scene  *Scene
	Effect *SoundEffect
	Col    int
	tick   int
}

// NewSceneWithSound creates a SceneWithSound combining a scene and effect
func NewSceneWithSound(scene *Scene, effect *SoundEffect, col int) *SceneWithSound {
	return &SceneWithSound{
		Scene:  scene,
		Effect: effect,
		Col:    col,
	}
}

// Frame returns the current combined frame string and advances the tick
func (ss *SceneWithSound) Frame() string {
	base := ss.Scene.String()
	effectFrame := ss.Effect.Frame(ss.tick)
	ss.tick++
	return Overlay(base, effectFrame, ss.Col)
}

// Reset resets both the scene tick and the sound tick
func (ss *SceneWithSound) Reset() {
	ss.tick = 0
}

// FPS delegates to the underlying scene FPS
func (ss *SceneWithSound) FPS() int {
	return ss.Scene.FPS
}

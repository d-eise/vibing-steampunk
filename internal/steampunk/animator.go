package steampunk

import (
	"fmt"
	"time"
)

// Animator drives frame-by-frame playback of a Scene
type Animator struct {
	Scene   *Scene
	current int
	stop    chan struct{}
}

// NewAnimator creates an Animator for the given scene
func NewAnimator(scene *Scene) *Animator {
	return &Animator{
		Scene: scene,
		stop:  make(chan struct{}),
	}
}

// Start begins the animation loop, printing frames to stdout.
// The loop wraps around when it reaches the last frame, so playback
// continues indefinitely until Stop is called.
func (a *Animator) Start() {
	interval := time.Second / time.Duration(a.Scene.FPS)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	frameCount := len(a.Scene.Frames)
	for {
		select {
		case <-a.stop:
			return
		case <-ticker.C:
			a.PrintFrame()
			if frameCount > 0 {
				a.current = (a.current + 1) % frameCount
			} else {
				a.current++
			}
		}
	}
}

// Stop halts the animation loop
func (a *Animator) Stop() {
	close(a.stop)
}

// PrintFrame prints the current frame to stdout
func (a *Animator) PrintFrame() {
	fmt.Print("\033[H\033[2J") // clear screen
	fmt.Println(Banner())
	fmt.Println(GearLine())
	fmt.Println(a.Scene.Frame(a.current))
	fmt.Println(GearLine())
}

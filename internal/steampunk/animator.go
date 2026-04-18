package steampunk

import (
	"fmt"
	"time"
)

// Animator cycles through frames and writes them to a writer-like callback
type Animator struct {
	frames   []string
	Interval time.Duration
	stop     chan struct{}
}

// NewAnimator creates an Animator with the given frames and tick interval
func NewAnimator(frames []string, interval time.Duration) *Animator {
	return &Animator{
		frames:   frames,
		Interval: interval,
		stop:     make(chan struct{}),
	}
}

// Run starts the animation loop, calling render on each frame tick.
// It blocks until Stop() is called.
func (a *Animator) Run(render func(frame string)) {
	ticker := time.NewTicker(a.Interval)
	defer ticker.Stop()
	idx := 0
	for {
		select {
		case <-ticker.C:
			render(a.frames[idx%len(a.frames)])
			idx++
		case <-a.stop:
			return
		}
	}
}

// Stop signals the animation loop to exit
func (a *Animator) Stop() {
	close(a.stop)
}

// PrintFrame is a convenience render function that clears the line and prints the frame
func PrintFrame(frame string) {
	fmt.Printf("\r%s", frame)
}

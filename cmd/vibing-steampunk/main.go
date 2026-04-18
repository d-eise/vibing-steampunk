package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/oisee/vibing-steampunk/internal/steampunk"
)

func main() {
	message := flag.String("msg", "VIBING STEAMPUNK", "Banner message to display")
	gears := flag.Int("gears", 5, "Number of gears in the gear line")
	duration := flag.Duration("duration", 5*time.Second, "How long to animate")
	flag.Parse()

	fmt.Println(steampunk.Banner(*message))
	fmt.Println(steampunk.GearLine(*gears))
	fmt.Println()

	anim := steampunk.NewAnimator(steampunk.SteamFrames, 200*time.Millisecond)

	go func() {
		time.Sleep(*duration)
		anim.Stop()
	}()

	anim.Run(steampunk.PrintFrame)

	fmt.Fprintln(os.Stdout, "\n[boiler offline]")
}

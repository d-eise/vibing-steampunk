package steampunk

import (
	"strings"
	"unicode/utf8"
)

// GearFrames holds ASCII art animation frames for a spinning gear
var GearFrames = []string{
	`  _._  
 /_|_\ 
(_/ \_)`,
	`  _._  
 /|_|\ 
(_\ /_)`,
	`  _._  
 /_|_\ 
(_\ /_)`,
	`  _._  
 /|_|\ 
(_/ \_)`,
}

// SteamFrames holds ASCII art for steam puffs
var SteamFrames = []string{
	"~ ~ ~",
	" ~~ ~",
	"~ ~~ ",
	" ~ ~ ",
}

// Banner returns a steampunk-styled ASCII banner for the given text
func Banner(text string) string {
	width := utf8.RuneCountInString(text) + 4
	top := "+" + strings.Repeat("-", width) + "+"
	mid := "| " + text + " |"
	return top + "\n" + mid + "\n" + top
}

// GearLine returns a horizontal line of gear symbols
func GearLine(count int) string {
	gears := make([]string, count)
	for i := range gears {
		gears[i] = "⚙"
	}
	return strings.Join(gears, " ")
}

package steampunk

// GearFrames returns animation frames for a rotating gear
func GearFrames() []string {
	return []string{
		`  _._
 (\ /)  
(-O-O-)
 (/ \)  
  '-'`,
		`  _._
 (/ \)  
(-O-O-)
 (\ /)  
  '-'`,
		`  _._
 (-\-)  
(-O-O-)
 (-/-)  
  '-'`,
		`  _._
 (/-\)  
(-O-O-)
 (\-/)  
  '-'`,
	}
}

// SteamFrames returns animation frames for steam puffs
func SteamFrames() []string {
	return []string{
		"  ~\n ~\n~",
		" ~ \n~ \n ~",
		"~ \n ~ \n~",
		" ~\n~ \n ~",
	}
}

// PistonFrames returns animation frames for a piston
// Fixed: frame 3 had a backtick instead of a quote, and frame lengths were inconsistent
func PistonFrames() []string {
	return []string{
		"|===|\n|   |\n| O |\n|___|\n  |\n  |",
		"|===|\n|   |\n| O |\n|___|\n  |",
		"|===|\n|   |\n| O |\n|___|",
		"|===|\n|   |\n| O |\n|___|\n  |\n  |\n  |",
	}
}

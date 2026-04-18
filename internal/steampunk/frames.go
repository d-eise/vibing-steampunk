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
func PistonFrames() []string {
	return []string{
		"|===|
|   |
| O |
|___|
  |
  |",
		"|===|
|   |
| O |
|___|
  |",
		"|===|
|   |
| O |
|___|`,
		"|===|
|   |
| O |
|___|
  |
  |
  |",
	}
}

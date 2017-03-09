package main

type Romset struct {
	offset int
	roms   []string
}

var romSets = []Romset{
	{
		offset: 0x800,
		roms: []string{
			"ROMs/136043-1111.1a",
			"ROMs/136043-1113.1l",
			"ROMs/136043-1115.2a",
			"ROMs/136043-1117.2l",
		},
	},
	{
		offset: 0x0,
		roms: []string{
			"ROMs/136043-1111.1a",
			"ROMs/136043-1113.1l",
			"ROMs/136043-1115.2a",
			"ROMs/136043-1117.2l",
		},
	},
	{
		offset: 0x800,
		roms: []string{
			"ROMs/136037-112.1b",
			"ROMs/136037-114.1mn",
			"ROMs/136037-116.2b",
			"ROMs/136037-118.2mn",
		},
	},
	{
		offset: 0x0,
		roms: []string{
			"ROMs/136037-112.1b",
			"ROMs/136037-114.1mn",
			"ROMs/136037-116.2b",
			"ROMs/136037-118.2mn",
		},
	},
	{
		offset: 0x0,
		roms: []string{
			"ROMs/136043-1123.1c",
			"ROMs/136043-1124.1p",
			"ROMs/136043-1125.2c",
			"ROMs/136043-1126.2p",
		},
	},
}

// returns the actual tile number to use, and the rom set to use it with
// Kind of a mess, since it uses knowledge for calculating the tile number
// that should be contained in the above structs, but isn't
func getromset(tilenum int) (int, []string) {
	whichbank := tilenum / 0x800
	actualtile := (tilenum % 0x800) + romSets[whichbank].offset
	// fmt.Printf("tile: 0x%x   romset: %s\n", actualtile, romSets[whichbank].roms)

	return actualtile, romSets[whichbank].roms
}

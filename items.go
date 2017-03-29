package main

import (
	"regexp"
	"strings"
)

// type Stamp struct {
//     width   int
//     numbers []int
//     data    []TileData
//     ptype   string
//     pnum    int
//     trans0  bool
// }

var itemStamps = map[string]Stamp{
	"blank": Stamp{
		width:   2,
		numbers: []int{0, 0, 0, 0},
		ptype:   "base",
		pnum:    0,
		trans0:  false,
	},

	"key": Stamp{
		width:   2,
		numbers: tilerange(0xafc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"keyring": Stamp{
		width:   3,
		numbers: tilerange(0x1d76, 6),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},

	"ifood1": Stamp{
		width:   3,
		numbers: tilerange(0x96c, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"ifood2": Stamp{
		width:   3,
		numbers: tilerange(0x975, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"ifood3": Stamp{
		width:   3,
		numbers: tilerange(0x97e, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"mfood": Stamp{
		width:   3,
		numbers: tilerange(0x277b, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"pfood": Stamp{
		width:   3,
		numbers: tilerange(0x25ed, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},

	"potion": Stamp{
		width:   2,
		numbers: tilerange(0x8fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"ipotion": Stamp{
		width:   2,
		numbers: tilerange(0x9fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"ppotion": Stamp{
		width:   2,
		numbers: tilerange(0x20fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},

	"shieldpotion": Stamp{
		width:   2,
		numbers: tilerange(0x11fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"speedpotion": Stamp{
		width:   2,
		numbers: tilerange(0x12fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"magicpotion": Stamp{
		width:   2,
		numbers: tilerange(0x13fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"shotpowerpotion": Stamp{
		width:   2,
		numbers: tilerange(0x14fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"shotspeedpotion": Stamp{
		width:   2,
		numbers: tilerange(0x15fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"fightpotion": Stamp{
		width:   2,
		numbers: tilerange(0x16fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},

	"invis": Stamp{
		width:   3,
		numbers: tilerange(0x1700, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"transportability": Stamp{
		width:   2,
		numbers: tilerange(0x23fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"reflect": Stamp{
		width:   2,
		numbers: tilerange(0x24fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"repulse": Stamp{
		width:   2,
		numbers: tilerange(0x26fc, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"invuln": Stamp{
		width:   2,
		numbers: tilerange(0x2784, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"supershot": Stamp{
		width:   2,
		numbers: tilerange(0x2788, 4),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},

	"pushwall": Stamp{
		width:   3,
		numbers: tilerange(0x20f6, 6),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},

	"treasure": Stamp{
		width:   3,
		numbers: tilerange(0x987, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"treasurelocked": Stamp{
		width:   3,
		numbers: tilerange(0x25e4, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"goldbag": Stamp{
		width:   3,
		numbers: tilerange(0x9a2, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},

	// FIXME: wrong palette
	"tport": Stamp{
		width:   2,
		numbers: tilerange(0x49e, 4),
		ptype:   "base",
		pnum:    2,
		trans0:  true,
	},

	// FIXME: wrong palette
	// FIXME: also missing all the various directions
	"ff": Stamp{
		width:   2,
		numbers: tilerange(0x4a2, 4),
		ptype:   "base",
		pnum:    2,
		trans0:  true,
	},

	"exit": Stamp{
		width:   2,
		numbers: []int{0x39e, 0x39f, 0x6, 0x6},
		ptype:   "floor",
		pnum:    0,
		trans0:  false,
	},
	"exit4": Stamp{
		width:   2,
		numbers: tilerange(0xcfc, 4),
		ptype:   "floor",
		pnum:    0,
		trans0:  false,
	},
	"exit6": Stamp{
		width:   2,
		numbers: tilerange(0x39e, 4),
		ptype:   "floor",
		pnum:    0,
		trans0:  false,
	},
	"exit8": Stamp{
		width:   2,
		numbers: tilerange(0xdfc, 4),
		ptype:   "floor",
		pnum:    0,
		trans0:  false,
	},

	// FIXME: Incorrect tiles, but works for now?
	"vdoor": Stamp{
		width:   2,
		numbers: tilerange(0x1d80, 4),
		ptype:   "base",
		pnum:    0,
		trans0:  true,
	},
	// FIXME: Also wrong tiles
	"hdoor": Stamp{
		width:   3,
		numbers: tilerange(0x1d52, 3),
		ptype:   "base",
		pnum:    0,
		trans0:  true,
	},

	"plus": Stamp{
		width:   2,
		numbers: tilerange(0xbfc, 4),
		ptype:   "floor",
		pnum:    1,
		trans0:  true,
	},

	// FIXME: Needs to be in monsters, really
	"dragon": Stamp{
		width:   4,
		numbers: tilerange(0x2100, 16),
		ptype:   "base",
		pnum:    8, // or 7 or 6
		trans0:  true,
	},
	"generator1": Stamp{
		width:   3,
		numbers: tilerange(0x9c6, 9),
		ptype:   "base",
		pnum:    5,
		trans0:  true,
	},
	"generator2": Stamp{
		width:   3,
		numbers: tilerange(0x9cf, 9),
		ptype:   "base",
		pnum:    5,
		trans0:  true,
	},
	"generator3": Stamp{
		width:   3,
		numbers: tilerange(0x9d8, 9),
		ptype:   "base",
		pnum:    5,
		trans0:  true,
	},
	"ghostgen1": Stamp{
		width:   3,
		numbers: tilerange(0x9ab, 9),
		ptype:   "base",
		pnum:    5,
		trans0:  true,
	},
	"ghostgen2": Stamp{
		width:   3,
		numbers: tilerange(0x9b4, 9),
		ptype:   "base",
		pnum:    5,
		trans0:  true,
	},
	"ghostgen3": Stamp{
		width:   3,
		numbers: tilerange(0x9bd, 9),
		ptype:   "base",
		pnum:    5,
		trans0:  true,
	},

	"ghost": Stamp{
		width:   3,
		numbers: tilerange(0x800, 9),
		ptype:   "base",
		pnum:    4,
		trans0:  true,
	},
	"grunt": Stamp{
		width:   3,
		numbers: tilerange(0x9e1, 9),
		ptype:   "base",
		pnum:    4,
		trans0:  true,
	},
	"demon": Stamp{
		width:   3,
		numbers: tilerange(0x183f, 9),
		ptype:   "base",
		pnum:    8,
		trans0:  true,
	},
	"lobber": Stamp{
		width:   3,
		numbers: tilerange(0x1b57, 9),
		ptype:   "base",
		pnum:    11,
		trans0:  true,
	},
	"sorcerer": Stamp{
		width:   3,
		numbers: tilerange(0x13a2, 9),
		ptype:   "base",
		pnum:    11,
		trans0:  true,
	},
	"auxgrunt": Stamp{
		width:   3,
		numbers: tilerange(0x9e1, 9),
		ptype:   "base",
		pnum:    4,
		trans0:  true,
	},
	"death": Stamp{
		width:   3,
		numbers: tilerange(0x1a75, 9),
		ptype:   "base",
		pnum:    0,
		trans0:  true,
	},
	"acid": Stamp{
		width:   3,
		numbers: tilerange(0x2300, 9),
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
	"supersorc": Stamp{
		width:   3,
		numbers: tilerange(0x13a2, 9),
		ptype:   "base",
		pnum:    11,
		trans0:  true,
	},
	"it": Stamp{
		width:   3,
		numbers: tilerange(0x2600, 9),
		ptype:   "base",
		pnum:    8,
		trans0:  true,
	},

	"arrowleft": Stamp{
		width: 2,
		// numbers: tilerange(0x1c8f, 4),
		numbers: []int{0, 0x1c8f, 0, 0x1c91},
		ptype:   "elf",
		pnum:    1,
		trans0:  true,
	},
	"arrowright": Stamp{
		width: 2,
		// numbers: tilerange(0x1c7c, 4),
		numbers: []int{0x1c7d, 0, 0x1c7f, 0},
		ptype:   "elf",
		pnum:    1,
		trans0:  true,
	},
	"arrowup": Stamp{
		width: 2,
		// numbers: tilerange(0x1c74, 4),
		numbers: []int{0, 0, 0x1c74, 0x1c75},
		ptype:   "elf",
		pnum:    1,
		trans0:  true,
	},
	"arrowdown": Stamp{
		width: 2,
		// numbers: tilerange(0x1c84, 4),
		numbers: []int{0x1c86, 0x1c87, 0, 0},
		ptype:   "elf",
		pnum:    1,
		trans0:  true,
	},
}

func tilerange(start int, count int) []int {
	r := make([]int, count)
	for i := range r {
		r[i] = start
		start += 1
	}
	return r
}

// type MobAnimFrames []int
// type MobAnimsDir map[string]MobAnimFrames
// type MobAnims map[string]MobAnimsDir

// var ghostAnims = MobAnims{
//     "walk": {
//         "up":        {0x890, 0x899, 0x8a2, 0x8ab},
//         "upright":   {0x86c, 0x875, 0x87e, 0x887},
//         "right":     {0x848, 0x851, 0x85a, 0x863},
//         "downright": {0x824, 0x82d, 0x836, 0x83f},
//         "down":      {0x800, 0x809, 0x812, 0x81b},
//         "downleft":  {0x900, 0x909, 0x912, 0x91b},
//         "left":      {0x8d8, 0x8e1, 0x8ea, 0x8f3},
//         "upleft":    {0x8b4, 0x8bd, 0x8c6, 0x8cf},
//     },
// }

// var monsters = map[string]Monster{
// 	"ghost": {
// 		xsize: 3,
// 		ysize: 3,
// 		ptype: "base",
// 		pnum:  0, // FIXME: This is weird and seems wrong
// 		// palette: gauntletPalettes["base"][0],
// 		anims: ghostAnims,
// 	},
// }

// var reMonsterType = regexp.MustCompile(`^(ghost)(\d+)?`)
// var reMonsterAction = regexp.MustCompile(`^(walk|fight|attack)`)
// var reMonsterDir = regexp.MustCompile(`^(up|upright|right|downright|down|downleft|left|upleft)`)

var reItemType = regexp.MustCompile(`^(key)$`)

func doitem(arg string) {
	split := strings.Split(arg, "-")

	var itemType string

	// Still wonder if there's a cleaner way
	for _, ss := range split {
		switch {
		case reItemType.MatchString(ss):
			item := reItemType.FindStringSubmatch(ss)
			itemType = item[1]
		}
	}

	stamp := itemGetStamp(itemType)

	height := len(stamp.numbers) / stamp.width
	img := blankimage(8*stamp.width, 8*height)
	writestamptoimage(img, stamp, 0, 0)
	savetopng(opts.Output, img)
}

// FIXME: In the future, maybe just return nil and not panic
func itemGetStamp(itemType string) *Stamp {
	stamp, ok := itemStamps[itemType]

	if ok == false {
		panic("requested bad item: " + itemType)
	}

	fillstamp(&stamp)
	return &stamp
}

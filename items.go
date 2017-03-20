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
	"key": Stamp{
		width:   2,
		numbers: []int{0xafc, 0xafd, 0xafe, 0xaff},
		ptype:   "base",
		pnum:    1,
		trans0:  true,
	},
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

	stamp := getitemstamp(itemType)

	height := len(stamp.numbers) / stamp.width
	img := blankimage(8*stamp.width, 8*height)
	writestamptoimage(img, stamp, 0, 0)
	savetopng(opts.Output, img)
}

func getitemstamp(itemType string) *Stamp {
	stamp := itemStamps[itemType]
	fillstamp(&stamp)

	return &stamp
}

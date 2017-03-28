package main

import (
	"math/rand"
)

var ztestmaze = []string{
	"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", // 1
	"                     0          ",
	" bbb                            ",
	"  b         b   b               ",
	"          bbb 0 bbb             ",
	"            b   b               ",
	"  bbb                  bbb      ",
	"  b b                  bbb      ",
	"  bbb                  bbb      ",
	"           b       b            ", // 10
	"            b      b    bbbbb   ",
	"             b     b            ",
	"      b                         ",
	"     b                          ",
	"    b                           ",
	"             c                  ",
	"                                ",
	"                                ",
	"                                ",
	"   1    2    3    4    5    6   ", // 20
	"                                ",
	"   7              p    q    r   ",
	"                                ",
	"   s    t    u    v    w    x   ",
	"                                ",
	"   y    z    A    B    C    D   ",
	"                                ",
	"   E    F    G    T    U    V   ",
	"                                ",
	"   W    X    Y    Z    0        ", // 30
	"                                ",
	"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
}

var maze0 = []int{
	0x0,  // secret trick
	0x0,  // flags 1
	0x0,  // flags 2
	0x0,  // flags 3
	0x0,  // flags 4
	0x61, // patterns
	0x1D, // colors
	0x6E, 0x54, 0xE, 0x9, 0x35, 0x80, 0x47, 0xC4, 0x31,
	0x42, 0xC9, 0x35, 0x6F, 0xDE, 0x14, 0xC8, 0x1D,
	0xC1, 0x1D, 0xC4, 0x1E, 0xCB, 0x22, 0xDE, 0x14,
	0xDD, 0xC, 0x81, 0xDF, 0xD9, 0x9, 0x81, 0xC3,
	0x9, 0x81, 0xD0, 0x58, 0xDF, 0xC3, 0x24, 0xC6,
	0x24, 0xC2, 0x59, 0xDF, 0xD0, 0x20, 0x40, 0xC0,
	0x74, 0xC0, 0x11, 0xC0, 0x74, 0x40, 0xC0, 0x20,
	0xC9, 0x7, 0x84, 0xB3, 0xA2, 0xB5, 0xA8, 0xB5,
	0xA1, 0xB2, 0x8, 0x86, 0xC4, 0xA, 0x80, 0xCD,
	0xB, 0x80, 0xDF, 0xC7, 0xBF, 0xDF, 0xDF, 0xDF,
	0xDF, 0xC1, 0x1E, 0xDC, 0xA4, 0xDA, 0x10, 0xC9,
	0x1C, 0xC3, 0x1C, 0xDF, 0x45, 0xC2, 0x32, 0xC0,
	0x1E, 0xC5, 0xB5, 0xA6, 0xD, 0x82, 0xA5, 0xB5,
	0xA8, 0xDD, 0x2E, 0xC5, 0x1F, 0xDF, 0xDF, 0x48,
	0xCA, 0xF, 0xDF, 0xDF, 0xC2, 0xBE, 0x35, 0xDC,
	0x2E, 0x0}

var maze2 = []int{
	0x0, 0x0, 0x0, 0x2, 0x0, 0x20, 0x12, 0x5C,
	0x6E, 0x59, 0x5F, 0x10, 0x80, 0xC0, 0xE, 0x19,
	0xE, 0xC8, 0x20, 0x45, 0xE0, 0x71, 0xC2, 0x35,
	0x51, 0xE, 0xC0, 0xE, 0xD8, 0xA1, 0xD, 0x81,
	0xC0, 0xE, 0xC2, 0x1E, 0x60, 0xD2, 0x1D, 0x50,
	0x51, 0xE, 0xE7, 0x63, 0xE1, 0xC7, 0xA1, 0xD,
	0x83, 0xF6, 0xD2, 0x1C, 0xF8, 0xE5, 0x77, 0x4D,
	0xEB, 0xCA, 0x1E, 0xE4, 0xF0, 0xF4, 0xDF, 0x72,
	0x55, 0x83, 0x69, 0xE1, 0x50, 0xDD, 0x19, 0x62,
	0x65, 0xC5, 0x20, 0xC0, 0xA1, 0x58, 0x40, 0xD8,
	0x1F, 0xA0, 0x50, 0xD1, 0xBF, 0xD5, 0x1C, 0xE1,
	0x41, 0xE9, 0xDA, 0x1C, 0xDE, 0x2E, 0xED, 0x63,
	0x63, 0x7E, 0xF7, 0x42, 0xE0, 0xD3, 0x1F, 0xDF,
	0x4D, 0x61, 0xD1, 0x2E, 0xDE, 0x1C, 0xDF, 0xE1,
	0x61, 0xE2, 0xC8, 0x1B, 0x6A, 0x55, 0xDE, 0x19,
	0xDA, 0x1C, 0xE1, 0x19, 0xDE, 0x19, 0x41, 0x61,
	0xCB, 0xBE, 0x35, 0xA0, 0x1D, 0xA0, 0x43, 0xD4,
	0xF, 0x0,
}

var maze12 = []int{
	0x8, 0x0, 0x42, 0x86, 0x0, 0x51, 0xA, 0xAE,
	0x5D, 0xE, 0x90, 0x33, 0x80, 0xC3, 0x1F, 0x41,
	0x1C, 0xC7, 0x42, 0x22, 0xC2, 0x2E, 0xA0, 0x2F,
	0xC0, 0x7F, 0xEA, 0xC1, 0x24, 0xC6, 0x18, 0xF7,
	0xFD, 0x44, 0xA9, 0x2E, 0xA5, 0xCC, 0x20, 0xC0,
	0x76, 0x40, 0x14, 0xC4, 0x2F, 0xA3, 0xC7, 0x18,
	0x47, 0x1D, 0xCB, 0x74, 0x22, 0x66, 0xA0, 0xEA,
	0xC0, 0x1F, 0xEF, 0x49, 0xA0, 0x1C, 0xA0, 0x35,
	0x1F, 0xC8, 0x1F, 0xC8, 0x42, 0x35, 0xE3, 0x1C,
	0xA0, 0xC9, 0x4F, 0xE2, 0x1C, 0xA0, 0xE2, 0x6A,
	0xC0, 0x40, 0x1D, 0xE9, 0xC3, 0x1C, 0xA0, 0xD0,
	0x58, 0xC0, 0x58, 0x60, 0xCB, 0x43, 0xB8, 0x47,
	0xA0, 0xE0, 0xD2, 0xD, 0x87, 0xA0, 0xE2, 0xC6,
	0x2E, 0xB8, 0xC4, 0xBF, 0xC2, 0x2E, 0x14, 0xC2,
	0x35, 0xC4, 0xF, 0xE0, 0xD, 0x87, 0xC8, 0xD,
	0x87, 0xA0, 0xC6, 0x35, 0xC2, 0x14, 0x49, 0x47,
	0xA1, 0xE0, 0xD, 0x87, 0x6B, 0xC0, 0x41, 0x1D,
	0xCB, 0x4F, 0x68, 0x63, 0xDF, 0xE7, 0xC2, 0x45,
	0x42, 0x35, 0xC3, 0x41, 0x47, 0x1C, 0xC2, 0xA1,
	0xD7, 0x1D, 0xA1, 0xC2, 0x21, 0xD8, 0x2E, 0x81,
	0x69, 0xC7, 0x41, 0x1D, 0xD9, 0x1D, 0xED, 0xC0,
	0xB9, 0x70, 0x2F, 0xC0, 0x43, 0x40, 0x14, 0xC0,
	0x42, 0x41, 0x10, 0xB8, 0x66, 0xC0, 0x35, 0xA2,
	0xB2, 0xA5, 0x59, 0x2E, 0x59, 0xA7, 0x35, 0xDF,
	0xDF, 0xC4, 0xBE, 0x72, 0x1D, 0x45, 0x1C, 0x40,
	0x38, 0xE1, 0x1C, 0xA0, 0xC5, 0x24, 0xC0, 0x24,
	0x3B, 0x0,
}

func genpfimage() {
	// 8 pixels * 2 tiles * 32 stamps, plus extra space on edges
	img := blankimage(8*2*32+32, 8*2*32+32)

	// mazes will always be the same size, so just use constants
	maze := mazeDecompress(maze0)
	// spew.Dump(maze)

	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			stamp := floorGetStamp(maze.floorpattern, checkadj3(maze, x, y)+rand.Intn(4), maze.floorcolor)
			writestamptoimage(img, stamp, x*16+16, y*16+16)
		}
	}

	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			var stamp *Stamp

			// We should do better
			switch whatis(maze, x, y) {
			case MAZEOBJ_TILE_FLOOR:
				adj := checkadj3(maze, x, y) + rand.Intn(4)
				stamp = floorGetStamp(maze.floorpattern, adj, maze.floorcolor)
			case MAZEOBJ_WALL_REGULAR:
				adj := checkadj8(maze, x, y)
				stamp = wallGetStamp(maze.wallpattern, adj, maze.wallcolor)
			case MAZEOBJ_KEY:
				stamp = itemGetStamp("key")
			case MAZEOBJ_POWER_INVIS:
				stamp = itemGetStamp("invis")
			case MAZEOBJ_POWER_REPULSE:
				stamp = itemGetStamp("repulse")
			case MAZEOBJ_POWER_REFLECT:
				stamp = itemGetStamp("reflect")
			case MAZEOBJ_POWER_TRANSPORT:
				stamp = itemGetStamp("transportability")
			case MAZEOBJ_POWER_SUPERSHOT:
				stamp = itemGetStamp("supershot")
			case MAZEOBJ_POWER_INVULN:
				stamp = itemGetStamp("invuln")
			case MAZEOBJ_MONST_DRAGON:
				stamp = itemGetStamp("dragon")
			case MAZEOBJ_DOOR_HORIZ:
				stamp = itemGetStamp("hdoor")
			case MAZEOBJ_DOOR_VERT:
				stamp = itemGetStamp("vdoor")
			case MAZEOBJ_HIDDENPOT:
				stamp = itemGetStamp("plus")
			case MAZEOBJ_EXIT:
				stamp = itemGetStamp("exit")
			case MAZEOBJ_EXITTO6:
				stamp = itemGetStamp("exit6")
			case MAZEOBJ_MONST_GHOST:
				stamp = itemGetStamp("ghost")
			case MAZEOBJ_MONST_GRUNT:
				stamp = itemGetStamp("grunt")
			case MAZEOBJ_MONST_DEMON:
				stamp = itemGetStamp("demon")
			case MAZEOBJ_MONST_LOBBER:
				stamp = itemGetStamp("lobber")
			case MAZEOBJ_MONST_SORC:
				stamp = itemGetStamp("sorcerer")
			case MAZEOBJ_MONST_AUX_GRUNT:
				stamp = itemGetStamp("auxgrunt")
			case MAZEOBJ_MONST_DEATH:
				stamp = itemGetStamp("death")
			case MAZEOBJ_MONST_ACID:
				stamp = itemGetStamp("acid")
			case MAZEOBJ_MONST_SUPERSORC:
				stamp = itemGetStamp("supersorc")
			case MAZEOBJ_MONST_IT:
				stamp = itemGetStamp("it")
			case MAZEOBJ_GEN_GHOST1:
				stamp = itemGetStamp("ghostgen1")
			case MAZEOBJ_GEN_GHOST2:
				stamp = itemGetStamp("ghostgen2")
			case MAZEOBJ_GEN_GHOST3:
				stamp = itemGetStamp("ghostgen3")

			case MAZEOBJ_GEN_GRUNT1:
				fallthrough
			case MAZEOBJ_GEN_DEMON1:
				fallthrough
			case MAZEOBJ_GEN_LOBBER1:
				fallthrough
			case MAZEOBJ_GEN_SORC1:
				fallthrough
			case MAZEOBJ_GEN_AUX_GRUNT1:
				stamp = itemGetStamp("generator1")

			case MAZEOBJ_GEN_GRUNT2:
				fallthrough
			case MAZEOBJ_GEN_DEMON2:
				fallthrough
			case MAZEOBJ_GEN_LOBBER2:
				fallthrough
			case MAZEOBJ_GEN_SORC2:
				fallthrough
			case MAZEOBJ_GEN_AUX_GRUNT2:
				stamp = itemGetStamp("generator2")

			case MAZEOBJ_GEN_GRUNT3:
				fallthrough
			case MAZEOBJ_GEN_DEMON3:
				fallthrough
			case MAZEOBJ_GEN_LOBBER3:
				fallthrough
			case MAZEOBJ_GEN_SORC3:
				fallthrough
			case MAZEOBJ_GEN_AUX_GRUNT3:
				stamp = itemGetStamp("generator3")

			case MAZEOBJ_TREASURE:
				stamp = itemGetStamp("treasure")
			case MAZEOBJ_TREASURE_LOCKED:
				stamp = itemGetStamp("treasurelocked")
			case MAZEOBJ_TREASURE_BAG:
				stamp = itemGetStamp("goldbag")
			case MAZEOBJ_FOOD_DESTRUCTABLE:
				stamp = itemGetStamp("mfood")
			case MAZEOBJ_FOOD_INVULN:
				stamp = itemGetStamp("ifood1")
			case MAZEOBJ_POT_DESTRUCTABLE:
				stamp = itemGetStamp("potion")
			case MAZEOBJ_POT_INVULN:
				stamp = itemGetStamp("ipotion")
			}

			if stamp != nil {
				writestamptoimage(img, stamp, x*16+16, y*16+16)
			}
		}
	}

	savetopng(opts.Output, img)
}

// check to see if there's walls adjacent left, left/up, and up
// FIXME: This might should be a different set of directions
// horizontal wall += 4
// diagonal wall += 8
// vertical wall += 16

func whatis(maze *Maze, x int, y int) int {
	if (maze.flags4 & LFLAG4_WRAP_H) != 0 {
		if x < 0 {
			x = 32 + x
		} else {
			x = x % 32
		}
	}

	if (maze.flags4 & LFLAG4_WRAP_V) != 0 {
		if x < 0 {
			x = 32 + x
		} else {
			y = y % 32
		}
	}

	if x < 0 || y < 0 || x >= 32 || y >= 32 {
		return 0
	}

	return maze.data[xy{x, y}]
}

func checkadj3(maze *Maze, x int, y int) int {
	adj := 0

	if whatis(maze, x-1, y) == MAZEOBJ_WALL_REGULAR {
		adj += 4
	}

	if whatis(maze, x, y+1) == MAZEOBJ_WALL_REGULAR {
		adj += 16
	}

	if whatis(maze, x-1, y+1) == MAZEOBJ_WALL_REGULAR {
		adj += 8
	}

	return adj
}

// check to see if there's walls on any side of location, for picking
// which wall tile needs ot be used
//
// Values to use:
//    up left:  0x01      up:         0x02      up right:  0x04
//    left:     0x08      right:      0x10      down left: 0x20
//    down:     0x40      down right: 0x80
//
// FIXME: Our sense of up/down here is probably confused

func checkadj8(maze *Maze, x int, y int) int {
	adj := 0

	if whatis(maze, x-1, y-1) == MAZEOBJ_WALL_REGULAR {
		adj += 0x01
	}
	if whatis(maze, x, y-1) == MAZEOBJ_WALL_REGULAR {
		adj += 0x02
	}
	if whatis(maze, x+1, y-1) == MAZEOBJ_WALL_REGULAR {
		adj += 0x04
	}
	if whatis(maze, x-1, y) == MAZEOBJ_WALL_REGULAR {
		adj += 0x08
	}
	if whatis(maze, x+1, y) == MAZEOBJ_WALL_REGULAR {
		adj += 0x010
	}
	if whatis(maze, x-1, y+1) == MAZEOBJ_WALL_REGULAR {
		adj += 0x20
	}
	if whatis(maze, x, y+1) == MAZEOBJ_WALL_REGULAR {
		adj += 0x40
	}
	if whatis(maze, x+1, y+1) == MAZEOBJ_WALL_REGULAR {
		adj += 0x80
	}

	return adj
}

// img := genimage(tile, opts.DimX, opts.DimY)
// f, _ := os.OpenFile(opts.Output, os.O_WRONLY|os.O_CREATE, 0600)
// defer f.Close()
// // gif.Encode(f, img, &gif.Options{NumColors: 16})
// png.Encode(f, img)

// func genimage(tilenum int, xtiles int, ytiles int) *image.NRGBA {
//     img := blankimage(8*xtiles, 8*ytiles)

//     tc := 0
//     for y := 0; y < ytiles; y++ {
//         for x := 0; x < xtiles; x++ {
//             tile := getparsedtile(tilenum + tc)
//             writetiletoimage(tile, img, x*8, y*8)
//             tc++
//         }
//     }

//     return img
// }

package main

import (
	"fmt"
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

// maze object ids (names matching names in IDA)
const (
	MAZEOBJ_TILE_FLOOR = iota
	MAZEOBJ_TILE_STUN
	MAZEOBJ_WALL_REGULAR
	MAZEOBJ_WALL_MOVABLE
	MAZEOBJ_WALL_SECRET
	MAZEOBJ_WALL_DESTRUCTABLE
	MAZEOBJ_WALL_RANDOM
	MAZEOBJ_WALL_TRAPCYC1
	MAZEOBJ_WALL_TRAPCYC2
	MAZEOBJ_WALL_TRAPCYC3
	MAZEOBJ_TILE_TRAP1
	MAZEOBJ_TILE_TRAP2
	MAZEOBJ_TILE_TRAP3
	MAZEOBJ_DOOR_HORIZ
	MAZEOBJ_DOOR_VERT
	MAZEOBJ_PLAYERSTART
	MAZEOBJ_EXIT
	MAZEOBJ_EXITTO6
	MAZEOBJ_MONST_GHOST
	MAZEOBJ_MONST_GRUNT
	MAZEOBJ_MONST_DEMON
	MAZEOBJ_MONST_LOBBER
	MAZEOBJ_MONST_SORC
	MAZEOBJ_MONST_AUX_GRUNT
	MAZEOBJ_MONST_DEATH
	MAZEOBJ_MONST_ACID
	MAZEOBJ_MONST_SUPERSORC
	MAZEOBJ_MONST_IT
	MAZEOBJ_GEN_GHOST1
	MAZEOBJ_GEN_GHOST2
	MAZEOBJ_GEN_GHOS3
	MAZEOBJ_GEN_GRUNT1
	MAZEOBJ_GEN_GRUNT2
	MAZEOBJ_GEN_GRUNT3
	MAZEOBJ_GEN_DEMON1
	MAZEOBJ_GEN_DEMON2
	MAZEOBJ_GEN_DEMON3
	MAZEOBJ_GEN_LOBBER1
	MAZEOBJ_GEN_LOBBER2
	MAZEOBJ_GEN_LOBBER3
	MAZEOBJ_GEN_SORC1
	MAZEOBJ_GEN_SORC2
	MAZEOBJ_GEN_SORC3
	MAZEOBJ_GEN_AUX_GRUNT1
	MAZEOBJ_GEN_AUX_GRUNT2
	MAZEOBJ_GEN_AUX_GRUNT3
	MAZEOBJ_TREASURE
	MAZEOBJ_TREASURE_LOCKED
	MAZEOBJ_TREASURE_BAG
	MAZEOBJ_FOOD_DESTRUCTABLE
	MAZEOBJ_FOOD_INVULN
	MAZEOBJ_POT_DESTRUCTABLE
	MAZEOBJ_POT_INVULN
	MAZEOBJ_KEY
	MAZEOBJ_POWER_INVIS
	MAZEOBJ_POWER_REPULSE
	MAZEOBJ_POWER_REFLECT
	MAZEOBJ_POWER_TRANSPORT
	MAZEOBJ_POWER_SUPERSHOT
	MAZEOBJ_POWER_INVULN
	MAZEOBJ_MONST_DRAGON
	MAZEOBJ_HIDDENPOT
	MAZEOBJ_TRANSPORTER
	MAZEOBJ_FORCEFIELDHUB
)

var typeArr = []int{
	' ', 'a', 'b', 'c', 'd', 'e' /* 'f', 'g', 'h', 'i', */, 'b', 'b', 'b', 'b',
	'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
	't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C',
	'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
	'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W',
	'X', 'Y', 'Z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
	'+',
}

func getbytefortype(t int) int {
	return typeArr[t]
}

func index2xy(index int) (x int, y int) {
	if index < 0 {
		fmt.Printf("ERROR: Coordinates requested for index < 0: %d\n", index)
		panic("Coordinates requested for index < 0")
	}

	y = index / 32
	x = index - (y * 32)

	// if y >= 32 {
	// 	y -= 32
	// }

	return
}

type Maze *[32][32]int

// Is a maze object a wall?
func iswall(t int) bool {
	if t == MAZEOBJ_WALL_REGULAR || t == MAZEOBJ_WALL_SECRET || t == MAZEOBJ_WALL_DESTRUCTABLE || t == MAZEOBJ_WALL_RANDOM || t == MAZEOBJ_WALL_TRAPCYC1 || t == MAZEOBJ_WALL_TRAPCYC2 || t == MAZEOBJ_WALL_TRAPCYC3 || t == MAZEOBJ_FORCEFIELDHUB {
		return true
	} else {
		return false
	}
}

// Is it a floor tile of some type (or similar)
func isspecialfloor(t int) bool {
	if t == MAZEOBJ_TILE_STUN || t == MAZEOBJ_TILE_TRAP1 || t == MAZEOBJ_TILE_TRAP2 || t == MAZEOBJ_TILE_TRAP3 || t == MAZEOBJ_EXIT || t == MAZEOBJ_EXITTO6 || t == MAZEOBJ_TRANSPORTER {
		return true
	} else {
		return false

	}
}

// FIXME: needs to handle vflip and hflip
func expand(maze Maze, location int, t int, count int) int {
	if t == MAZEOBJ_TILE_FLOOR {
		return (location + count)
	}

	for i := 0; i < count; i++ {
		if iswall(t) {
			x, y := index2xy(location + i)
			maze[y][x] = getbytefortype(t)

		} else if isspecialfloor(t) {
			x, y := index2xy(location + i)
			maze[y][x] = getbytefortype(t)
		} else {
			// things here will need an offset to be completely visible
			/* if t == MAZEOBJ_MONST_DRAGON */

			x, y := index2xy(location + i)
			maze[y][x] = getbytefortype(t)
		}
	}
	return location + count
}

// FIXME: Needs to handle vflip and hflip
func vexpand(maze Maze, location int, t int, count int) int {
	if t == MAZEOBJ_TILE_FLOOR {
		return location + 1
	}

	for i := 0; i < count; i++ {
		if iswall(t) {
			x, y := index2xy(location - (i * 32))
			maze[y][x] = getbytefortype(t)
		} else if isspecialfloor(t) {
			x, y := index2xy(location - (i * 32))
			maze[y][x] = getbytefortype(t)
		} else {
			// things here will need a position adjustment to be visible
			x, y := index2xy(location - (i * 32))
			maze[y][x] = getbytefortype(t)
		}
	}

	return location + 1
}

// Outoput is maze[y][x]
func mazeDecompress(compressed []int) Maze {
	rand.Seed(5)
	var m [32][32]int
	var maze = Maze(&m)

	htype1 := compressed[7]  // horiz type 1
	htype2 := compressed[8]  // horiz type 2
	vtype1 := compressed[9]  // vert type 1
	vtype2 := compressed[10] // vert type 2

	prev := htype2 // previous value, for 'repeat previous' types

	// Fill in first row with walls, always
	for i := 0; i < 32; i++ {
		maze[0][i] = 'b'
	}

	// Unpack here starts
	location := 32               // how many spots we've filled
	compressed = compressed[11:] // pointer to where we are in the input stream

	for location < 1024 {
		var token int
		//		fmt.Printf("Remaining input length: %d, output remaining: %d\n", len(level), 1024-p)
		token, compressed = compressed[0], compressed[1:]
		count := (token & 0x0f) + 1
		longcount := (token & 0x1f) + 1 // used for 'repeat last' and 'skip'

		// fmt.Printf("Position: %d, processing byte 0x%02x: count:%d\n", p, n, cnt)

		switch token & 0xc0 { // look at top two bits
		case 0x00: // place one of literal object
			location = expand(maze, location, token&0x3f, 1)
			prev = token
		case 0x40: // Repeat special type
			switch token & 0x30 {
			case 0x00:
				prev = htype1
			case 0x10:
				prev = vtype1
			case 0x20:
				prev = htype2
			case 0x30:
				prev = vtype2
			}

			previtem := prev & 0x3f
			switch prev & 0xc0 {
			case 0x00: // repeat type
				if (token & 0x10) != 0 {
					location = vexpand(maze, location, previtem, count)
				} else {
					location = expand(maze, location, previtem, count)
				}
			case 0x40: // skip and add
				location = expand(maze, location, MAZEOBJ_TILE_FLOOR, count)
				location = expand(maze, location, previtem, 1)
			case 0x80: // add and skip
				location = expand(maze, location, previtem, 1)
				location = expand(maze, location, MAZEOBJ_TILE_FLOOR, count)
			case 0xc0: // repeat wall and add
				location = expand(maze, location, MAZEOBJ_WALL_REGULAR, count)
				location = expand(maze, location, previtem, 1)
			}
		case 0x80: // repeat wall
			if (token & 0x20) != 0 { // Repeat wall
				if (token & 0x10) != 0 {
					// vertical
					location = vexpand(maze, location, MAZEOBJ_WALL_REGULAR, count)
				} else {
					// horizontal
					location = expand(maze, location, MAZEOBJ_WALL_REGULAR, count)
				}
			} else {
				location = expand(maze, location, prev&0x3f, longcount)
			}
		case 0xc0:
			if (token & 0x20) != 0 {
				// skip and add wall
				location = expand(maze, location, MAZEOBJ_TILE_FLOOR, longcount)
				location = expand(maze, location, MAZEOBJ_WALL_REGULAR, 1)
			} else {
				// just skip
				location = expand(maze, location, MAZEOBJ_TILE_FLOOR, longcount)
			}
		}
	}

	if len(compressed) != 1 || compressed[0] != 0 {
		fmt.Printf("WARNING: Incomplete maze decode? (%d bytes remaining)\n", len(compressed))
	}

	return maze
}

func genpfimage() {
	img := blankimage(8*2*32, 8*2*32) // 8 pixels * 2 tiles * 32 stamps

	//	spew.Dump("Stamp: %#v\n", stamp)
	// mazes will always be the same size, so just use constants
	maze := mazeDecompress(maze0)

	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			stamp := floorGetStamp(0, checkadj3(maze, x, y)+rand.Intn(4), 0)
			writestamptoimage(img, stamp, x*16, y*16)
		}
	}

	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			var stamp *Stamp

			switch whatis(maze, x, y) {
			case 'b':
				adj := checkadj8(maze, x, y)
				stamp = wallGetStamp(0, adj, 3)
			case '0':
				stamp = itemGetStamp("key")
			case '1':
				stamp = itemGetStamp("invis")
			case '2':
				stamp = itemGetStamp("repulse")
			case '3':
				stamp = itemGetStamp("reflect")
			case '4':
				stamp = itemGetStamp("transportability")
			case '5':
				stamp = itemGetStamp("supershot")
			case '6':
				stamp = itemGetStamp("invuln")
			case '7':
				stamp = itemGetStamp("dragon")

			case 'm':
				stamp = itemGetStamp("hdoor")
			case 'n':
				stamp = itemGetStamp("vdoor")
			case 'o':
				stamp = itemGetStamp("plus")
			case 'p':
				stamp = itemGetStamp("exit")
			case 'q':
				stamp = itemGetStamp("exit6")
			case 'r':
				stamp = itemGetStamp("ghost")

			case 's':
				stamp = itemGetStamp("grunt")
			case 't':
				stamp = itemGetStamp("demon")
			case 'u':
				stamp = itemGetStamp("lobber")
			case 'v':
				stamp = itemGetStamp("sorcerer")
			case 'w':
				stamp = itemGetStamp("auxgrunt")

			case 'x':
				stamp = itemGetStamp("death")
			case 'y':
				stamp = itemGetStamp("acid")
			case 'z':
				stamp = itemGetStamp("supersorc")
			case 'A':
				stamp = itemGetStamp("it")
			case 'B':
				stamp = itemGetStamp("ghostgen1")
			case 'C':
				stamp = itemGetStamp("ghostgen2")
			case 'D':
				stamp = itemGetStamp("ghostgen3")
			case 'E':
				stamp = itemGetStamp("generator1")

			case 'F':
				stamp = itemGetStamp("generator2")
			case 'G':
				stamp = itemGetStamp("generator3")
			case 'H':
				stamp = itemGetStamp("generator1")
			case 'I':
				stamp = itemGetStamp("generator2")
			case 'J':
				stamp = itemGetStamp("generator3")
			case 'K':
				stamp = itemGetStamp("generator1")
			case 'L':
				stamp = itemGetStamp("generator2")
			case 'M':
				stamp = itemGetStamp("generator3")
			case 'N':
				stamp = itemGetStamp("generator1")
			case 'O':
				stamp = itemGetStamp("generator2")
			case 'P':
				stamp = itemGetStamp("generator3")
			case 'Q':
				stamp = itemGetStamp("generator1")
			case 'R':
				stamp = itemGetStamp("generator2")
			case 'S':
				stamp = itemGetStamp("generator3")

			case 'T':
				stamp = itemGetStamp("treasure")
			case 'U':
				stamp = itemGetStamp("treasurelocked")
			case 'V':
			// stamp = itemGetStamp("goldbag")
			case 'W':
				stamp = itemGetStamp("mfood")
			case 'X':
				stamp = itemGetStamp("ifood1")
			case 'Y':
				stamp = itemGetStamp("potion")

			case 'Z':
				stamp = itemGetStamp("ipotion")
			}

			if stamp != nil {
				writestamptoimage(img, stamp, x*16, y*16)
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

func whatis(maze Maze, x int, y int) int {
	if x < 0 || y < 0 || x >= 32 || y >= 32 {
		return 0
	}

	return maze[y][x]
}

func checkadj3(maze Maze, x int, y int) int {
	adj := 0

	if whatis(maze, x-1, y) == 'b' {
		adj += 4
	}

	if whatis(maze, x, y+1) == 'b' {
		adj += 16
	}

	if whatis(maze, x-1, y+1) == 'b' {
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

func checkadj8(maze Maze, x int, y int) int {
	adj := 0

	if whatis(maze, x-1, y-1) == 'b' {
		adj += 0x01
	}
	if whatis(maze, x, y-1) == 'b' {
		adj += 0x02
	}
	if whatis(maze, x+1, y-1) == 'b' {
		adj += 0x04
	}
	if whatis(maze, x-1, y) == 'b' {
		adj += 0x08
	}
	if whatis(maze, x+1, y) == 'b' {
		adj += 0x010
	}
	if whatis(maze, x-1, y+1) == 'b' {
		adj += 0x20
	}
	if whatis(maze, x, y+1) == 'b' {
		adj += 0x40
	}
	if whatis(maze, x+1, y+1) == 'b' {
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

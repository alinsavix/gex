package main

import (
	"image"
	"math/rand"
)

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

// Okay, so we have a maze. We need to adjust the edges to take care of
// wrap or no wrap.
//
// If we're not wrapping in a direction, we need to duplicate the left and
// top walls, so that the maze will be enclosed.
func copyedges(maze *Maze) {
	for i := 0; i <= 32; i++ {
		if (maze.flags & LFLAG4_WRAP_H) == 0 {
			maze.data[xy{32, i}] = maze.data[xy{0, i}]
		}
	}

	for i := 0; i <= 32; i++ {
		if (maze.flags & LFLAG4_WRAP_V) == 0 {
			maze.data[xy{i, 32}] = maze.data[xy{i, 0}]
		}
	}
}

var foods = []string{"ifood1", "ifood2", "ifood3"}

func genpfimage(maze *Maze) {
	// 8 pixels * 2 tiles * 32 stamps, plus extra space on edges
	img := blankimage(8*2*32+32, 8*2*32+32)

	// mazes will always be the same size, so just use constants
	// maze := mazeDecompress(mazedata)
	copyedges(maze)
	paletteMakeSpecial(maze.floorpattern, maze.floorcolor, maze.wallcolor)

	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			stamp := floorGetStamp(maze.floorpattern, checkwalladj3(maze, x, y)+rand.Intn(4), maze.floorcolor)
			writestamptoimage(img, stamp, x*16+16, y*16+16)
		}
	}

	lastx := 32
	if maze.flags&LFLAG4_WRAP_H > 0 {
		lastx = 31
	}

	lasty := 32
	if maze.flags&LFLAG4_WRAP_V > 0 {
		lasty = 31
	}

	for y := 0; y <= lasty; y++ {
		for x := 0; x <= lastx; x++ {
			var stamp *Stamp
			var dots int // dot count

			// We should do better
			switch whatis(maze, x, y) {
			case MAZEOBJ_TILE_FLOOR:
			// adj := checkwalladj3(maze, x, y) + rand.Intn(4)
			// stamp = floorGetStamp(maze.floorpattern, adj, maze.floorcolor)
			case MAZEOBJ_TILE_STUN:
				adj := checkwalladj3(maze, x, y) + rand.Intn(4)
				stamp = floorGetStamp(maze.floorpattern, adj, maze.floorcolor)
				stamp.ptype = "stun" // use trap palette (FIXME: consider moving)
				stamp.pnum = 0

				// Tried to simplify these a bit with a goto, but golang didn't
				// like it ('jump into block'). I should figure out why.
			case MAZEOBJ_TILE_TRAP1:
				dots = 1
				fallthrough
			case MAZEOBJ_TILE_TRAP2:
				if dots == 0 {
					dots = 2
				}
				fallthrough
			case MAZEOBJ_TILE_TRAP3:
				if dots == 0 {
					dots = 3
				}
				adj := checkwalladj3(maze, x, y) + rand.Intn(4)
				stamp = floorGetStamp(maze.floorpattern, adj, maze.floorcolor)
				stamp.ptype = "trap" // use trap palette (FIXME: consider moving)
				stamp.pnum = 0
			case MAZEOBJ_WALL_DESTRUCTABLE:
				adj := checkwalladj8(maze, x, y)
				stamp = wallGetStamp(5, adj, maze.wallcolor)
			case MAZEOBJ_WALL_SECRET:
				adj := checkwalladj8(maze, x, y)
				stamp = wallGetStamp(maze.wallpattern, adj, maze.wallcolor)
				stamp.ptype = "secret"
				stamp.pnum = 0
			case MAZEOBJ_WALL_TRAPCYC1:
				dots = 1
				fallthrough
			case MAZEOBJ_WALL_TRAPCYC2:
				if dots == 0 {
					dots = 2
				}
				fallthrough
			case MAZEOBJ_WALL_TRAPCYC3:
				if dots == 0 {
					dots = 3
				}
				fallthrough
			case MAZEOBJ_WALL_RANDOM:
				if dots == 0 {
					dots = 4
				}
				fallthrough
			case MAZEOBJ_WALL_REGULAR:
				adj := checkwalladj8(maze, x, y)
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

			case MAZEOBJ_DOOR_HORIZ:
				adj := checkdooradj4(maze, x, y)
				stamp = doorGetStamp(DOOR_HORIZ, adj)
			case MAZEOBJ_DOOR_VERT:
				adj := checkdooradj4(maze, x, y)
				stamp = doorGetStamp(DOOR_VERT, adj)

			case MAZEOBJ_PLAYERSTART:
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
			case MAZEOBJ_MONST_DRAGON:
				stamp = itemGetStamp("dragon")

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
				stamp = itemGetStamp("food")
			case MAZEOBJ_FOOD_INVULN:
				stamp = itemGetStamp(foods[rand.Intn(3)])
			case MAZEOBJ_POT_DESTRUCTABLE:
				stamp = itemGetStamp("potion")
			case MAZEOBJ_POT_INVULN:
				stamp = itemGetStamp("ipotion")

			case MAZEOBJ_FORCEFIELDHUB:
				stamp = itemGetStamp("ff")
			case MAZEOBJ_TRANSPORTER:
				stamp = itemGetStamp("tport")
			default:
				// fmt.Printf("WARNING: Unhandled obj id 0x%02x\n", whatis(maze, x, y))
			}

			if stamp != nil {
				writestamptoimage(img, stamp, x*16+16, y*16+16)
			}

			if dots != 0 {
				renderdots(img, x*16+16, y*16+16, dots)
			}
		}
	}

	if maze.flags&LFLAG4_WRAP_H > 0 {
		l := itemGetStamp("arrowleft")
		r := itemGetStamp("arrowright")
		for i := 2; i <= 32; i++ {
			writestamptoimage(img, l, 0, i*16)
			writestamptoimage(img, r, 32*16+16, i*16)
		}
	}

	if maze.flags&LFLAG4_WRAP_V > 0 {
		u := itemGetStamp("arrowup")
		d := itemGetStamp("arrowdown")
		for i := 1; i < 32; i++ {
			writestamptoimage(img, u, i*16+16, 0)
			writestamptoimage(img, d, i*16+16, 32*16+16)
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
	return maze.data[xy{x, y}]
}

func isdoor(t int) bool {
	if t == MAZEOBJ_DOOR_HORIZ || t == MAZEOBJ_DOOR_VERT {
		return true
	} else {
		return false
	}
}

func checkwalladj3(maze *Maze, x int, y int) int {
	adj := 0

	if iswall(whatis(maze, x-1, y)) {
		adj += 4
	}

	if iswall(whatis(maze, x, y+1)) {
		adj += 16
	}

	if iswall(whatis(maze, x-1, y+1)) {
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

func checkwalladj8(maze *Maze, x int, y int) int {
	adj := 0

	if iswall(whatis(maze, x-1, y-1)) {
		adj += 0x01
	}
	if iswall(whatis(maze, x, y-1)) {
		adj += 0x02
	}
	if iswall(whatis(maze, x+1, y-1)) {
		adj += 0x04
	}
	if iswall(whatis(maze, x-1, y)) {
		adj += 0x08
	}
	if iswall(whatis(maze, x+1, y)) {
		adj += 0x010
	}
	if iswall(whatis(maze, x-1, y+1)) {
		adj += 0x20
	}
	if iswall(whatis(maze, x, y+1)) {
		adj += 0x40
	}
	if iswall(whatis(maze, x+1, y+1)) {
		adj += 0x80
	}

	return adj
}

// Look and see what doors are adjacent to this door
//
// Values to use:
//    up:  0x01    right:  0x02    down:  0x04    left:  0x08

func checkdooradj4(maze *Maze, x int, y int) int {
	adj := 0

	if isdoor(whatis(maze, x, y-1)) {
		adj += 0x01
	}
	if isdoor(whatis(maze, x+1, y)) {
		adj += 0x02
	}
	if isdoor(whatis(maze, x, y+1)) {
		adj += 0x04
	}
	if isdoor(whatis(maze, x-1, y)) {
		adj += 0x08
	}

	return adj
}

func dotat(img *image.NRGBA, xloc int, yloc int) {
	c := IRGB{0xffff}

	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			img.Set(xloc+x, yloc+y, c)
		}
	}
}

func renderdots(img *image.NRGBA, xloc int, yloc int, count int) {
	switch count {
	case 1:
		dotat(img, xloc+7, yloc+7)
	case 2:
		dotat(img, xloc+9, yloc+5)
		dotat(img, xloc+5, yloc+9)
	case 3:
		dotat(img, xloc+7, yloc+7)
		dotat(img, xloc+9, yloc+5)
		dotat(img, xloc+5, yloc+9)
	case 4:
		dotat(img, xloc+9, yloc+5)
		dotat(img, xloc+5, yloc+9)
		dotat(img, xloc+5, yloc+5)
		dotat(img, xloc+9, yloc+9)
	}
}

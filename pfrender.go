package main

import (
	"math/rand"
)

var testmaze = []string{
	"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
	"                                ",
	" bbb                            ",
	"  b         b   b               ",
	"          bbb   bbb             ",
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
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ", //20
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ",
	"                                ", // 30
	"                                ",
	"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb",
}

func genpfimage() {
	img := blankimage(8*2*32, 8*2*32) // 8 pixels * 2 tiles * 32 stamps

	//	spew.Dump("Stamp: %#v\n", stamp)
	// mazes will always be the same size, so just use constants
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			stamp := floorGetStamp(0, checkadj3(testmaze, x, y)+rand.Intn(4), 0)
			// fmt.Printf("Requesting write to image at %d, %d\n", x, y)
			writestamptoimage(img, stamp, x*16, y*16)
		}
	}

	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			switch what(testmaze, x, y) {
			case 'b':
				adj := checkadj8(testmaze, x, y)
				stamp := wallGetStamp(0, adj, 3)
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

func what(maze []string, x int, y int) byte {
	if x < 0 || y < 0 || x >= 32 || y >= 32 {
		return 0
	}

	return maze[y][x]
}

func checkadj3(maze []string, x int, y int) int {
	adj := 0

	if what(maze, x-1, y) == 'b' {
		adj += 4
	}

	if what(maze, x, y+1) == 'b' {
		adj += 16
	}

	if what(maze, x-1, y+1) == 'b' {
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

func checkadj8(maze []string, x int, y int) int {
	adj := 0

	if what(maze, x-1, y-1) == 'b' {
		adj += 0x01
	}
	if what(maze, x, y-1) == 'b' {
		adj += 0x02
	}
	if what(maze, x+1, y-1) == 'b' {
		adj += 0x04
	}
	if what(maze, x-1, y) == 'b' {
		adj += 0x08
	}
	if what(maze, x+1, y) == 'b' {
		adj += 0x010
	}
	if what(maze, x-1, y+1) == 'b' {
		adj += 0x20
	}
	if what(maze, x, y+1) == 'b' {
		adj += 0x40
	}
	if what(maze, x+1, y+1) == 'b' {
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

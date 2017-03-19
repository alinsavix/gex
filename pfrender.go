package main

import (
	"image/png"
	"math/rand"
	"os"
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
	for y := 0; y < 32; y++ {
		for x := 0; x < 32; x++ {
			stamp := floorGetStamp(0, rand.Intn(4), 0)
			// fmt.Printf("Requesting write to image at %d, %d\n", x, y)
			writestamptoimage(img, stamp, x*16, y*16)
		}
	}

	f, _ := os.OpenFile(opts.Output, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	// gif.Encode(f, img, &gif.Options{NumColors: 16})
	png.Encode(f, img)
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

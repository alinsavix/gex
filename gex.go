package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"os"
)

type Color interface {
	IRGB() (irgb uint16)
}

type IRGB struct {
	irgb uint16
}

func (c IRGB) RGBA() (r, g, b, a uint32) {
	i := uint32(c.irgb&0xf000) >> 12
	r = uint32(c.irgb&0x0f00) >> 8 * i
	g = uint32(c.irgb&0x00f0) >> 4 * i
	b = uint32(c.irgb&0x000f) * i

	r = r << 8
	g = g << 8
	b = b << 8
	a = 0xffff

	return
}

var romSets = [][]string{
	{
		"ROMs/136037-112.1b",
		"ROMs/136037-114.1mn",
		"ROMs/136037-116.2b",
		"ROMs/136037-118.2mn",
	},
	{
		"ROMs/136043-1111.1a",
		"ROMs/136043-1113.1l",
		"ROMs/136043-1115.2a",
		"ROMs/136043-1117.2l",
	},
	{
		"ROMs/136043-1123.1c",
		"ROMs/136043-1124.1p",
		"ROMs/136043-1125.2c",
		"ROMs/136043-1126.2p",
	},
}

// returns the actual tile number to use, and the rom set to use it with
func getromset(tilenum int) (int, []string) {
	whichrom := tilenum / 0x800
	actualtile := tilenum - (whichrom * 0x800)

	return actualtile, romSets[whichrom]
}

type TileLinePlane []byte

type TileLinePlaneSet [][]byte

type TileLineMerged []byte

type Tile []TileLineMerged

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func gettiledatafromfile(file string, tilenum int) TileLinePlane {
	f, err := os.Open(file)
	check(err)

	databytes := make([]byte, 8)

	f.Seek(int64(tilenum*8), 0)
	cnt, err := f.Read(databytes)
	check(err)

	if cnt != 8 {
		panic("failed to read full tile from file")
	}

	return databytes
}

func bytetobits(databyte byte) []byte {
	res := make([]byte, 8)

	// databyte = databyte ^ 0xff
	for i := 0; i < 8; i++ {
		if databyte%2 > 0 {
			res = append([]byte{0}, res...)
		} else {
			res = append([]byte{1}, res...)
		}
		databyte = databyte >> 1
	}

	return res
}

func mergeplanes(planes TileLinePlaneSet) TileLineMerged {
	mergedline := TileLineMerged{}

	for i := 0; i < 8; i++ {
		val := (planes[3][i] * byte(math.Pow(2, 3))) + (planes[2][i] * byte(math.Pow(2, 2))) + (planes[1][i] * byte(math.Pow(2, 1))) + (planes[0][i] * byte(math.Pow(2, 0)))
		mergedline = append(mergedline, val)
	}

	return mergedline
}

func blankimage(x int, y int) *image.Paletted {
	rect := image.Rect(0, 0, x, y)

	// palette 0 (more or less), for exits and such
	palette := gauntletPalettes["floor"][0]
	img := image.NewPaletted(rect, color.Palette(palette))
	return img
}

func getparsedtile(tilenum int) Tile {
	planedata := make([]TileLinePlane, 4)

	realtilenum, roms := getromset(tilenum)

	planedata[0] = gettiledatafromfile(roms[0], realtilenum)
	planedata[1] = gettiledatafromfile(roms[1], realtilenum)
	planedata[2] = gettiledatafromfile(roms[2], realtilenum)
	planedata[3] = gettiledatafromfile(roms[3], realtilenum)
	// fmt.Printf("planedata is: %d\n", planedata)

	// fulltile := Tile{}
	fulltile := make([]TileLineMerged, 8)

	// For each line in tile
	for i := 0; i < 8; i++ {
		linedata := make([][]byte, 4)
		linedata[0] = bytetobits(planedata[0][i])
		linedata[1] = bytetobits(planedata[1][i])
		linedata[2] = bytetobits(planedata[2][i])
		linedata[3] = bytetobits(planedata[3][i])
		// fmt.Printf("line is: %d\n", linedata)

		fulltile[i] = mergeplanes(linedata)
		// fmt.Printf("merged line is: %d\n", fulltile[i])
	}

	// fmt.Printf("tile is: %d\n", fulltile)
	return fulltile
}

// Write an 8x8 tile into a (usually) larger image
func writetiletoimage(tile Tile, img *image.Paletted, x int, y int) {
	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			img.SetColorIndex(x+i, y+j, tile[j][i])
		}
	}
}

func genimage(tilenum int, xtiles int, ytiles int) *image.Paletted {
	img := blankimage(8*xtiles, 8*ytiles)

	tc := 0
	for y := 0; y < ytiles; y++ {
		for x := 0; x < xtiles; x++ {
			tile := getparsedtile(tilenum + tc)
			writetiletoimage(tile, img, x*8, y*8)
			tc++
		}
	}

	return img
}

func main() {
	// databytes := gettilefromfile("ROMs/136043-1119.16s", 50)
	// fmt.Printf("byte(s): %02x\n", databytes)

	img := genimage(0xcfc, 2, 2)
	f, _ := os.OpenFile("test.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	gif.Encode(f, img, &gif.Options{NumColors: 16})
	//	fmt.Printf("planes work out to: %d\n", planedata)
	//	fmt.Printf("indexes work out to: %d\n", mergeplanes(planedata))
}

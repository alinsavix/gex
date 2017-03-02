package main

import (
	"fmt"
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

type TilePlane []byte

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func gettilefromfile(file string, tilenum int) []byte {
	f, err := os.Open(file)
	check(err)

	databytes := make([]byte, 8)

	f.Seek(int64(tilenum*8), 0)
	cnt, err := f.Read(databytes)
	check(err)

	if cnt != 8 {
		panic("failed to read full tile from file")
	}

	return TilePlane(databytes)
}

func bytetobits(databyte byte) []byte {
	res := []byte{}

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

func mergeplanes(planes [4][]byte) []byte {
	mergedline := []byte{}

	for i := 0; i < 8; i++ {
		val := (planes[3][i] * byte(math.Pow(2, 3))) + (planes[2][i] * byte(math.Pow(2, 2))) + (planes[1][i] * byte(math.Pow(2, 1))) + (planes[0][i] * byte(math.Pow(2, 0)))
		mergedline = append(mergedline, val)
	}

	return mergedline
}

func blanktile() *image.Paletted {
	rect := image.Rect(0, 0, 8, 8)

	// palette 0 (more or less), for exits and such
	palette := []color.Color{
		// IRGB{0x0},
		color.RGBA{0x0, 0x0, 0x0, 0xff},
		IRGB{0x7222},
		IRGB{0x9444},
		IRGB{0x7753},

		IRGB{0xf421},
		IRGB{0xb842},
		IRGB{0x0},
		//IRGB{0xffff},
		color.RGBA{0xff, 0xff, 0xff, 0xff},
		IRGB{0x8421},
		IRGB{0x7422},
		IRGB{0x9532},
		IRGB{0xc532},

		IRGB{0xf422},
		IRGB{0x7974},
		IRGB{0xe532},
		IRGB{0xf643},
	}

	img := image.NewPaletted(rect, color.Palette(palette))
	return img
}

func main() {
	// databytes := gettilefromfile("ROMs/136043-1119.16s", 50)
	// fmt.Printf("byte(s): %02x\n", databytes)

	planedata := [4][]byte{}

	planedata[0] = gettilefromfile("ROMs/136043-1111.1a", 0x4fc)
	planedata[1] = gettilefromfile("ROMs/136043-1113.1l", 0x4fc)
	planedata[2] = gettilefromfile("ROMs/136043-1115.2a", 0x4fc)
	planedata[3] = gettilefromfile("ROMs/136043-1117.2l", 0x4fc)

	fulltile := [8][]byte{}

	for i := 0; i < 8; i++ {
		linedata := [4][]byte{}
		linedata[0] = bytetobits(planedata[0][i])
		linedata[1] = bytetobits(planedata[1][i])
		linedata[2] = bytetobits(planedata[2][i])
		linedata[3] = bytetobits(planedata[3][i])

		fulltile[i] = mergeplanes(linedata)
	}
	fmt.Printf("tile is: %d\n", fulltile)

	img := blanktile()

	for j := 0; j < 8; j++ {
		for i := 0; i < 8; i++ {
			img.SetColorIndex(i, j, fulltile[j][i])
		}
	}

	f, _ := os.OpenFile("test.gif", os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()

	gif.Encode(f, img, &gif.Options{NumColors: 16})
	//	fmt.Printf("planes work out to: %d\n", planedata)
	//	fmt.Printf("indexes work out to: %d\n", mergeplanes(planedata))
}

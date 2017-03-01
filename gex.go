package main

import (
	"fmt"
	"math"
	"os"
)

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

	for i := 0; i < 8; i++ {
		if databyte%2 > 0 {
			// res = append(res, 1) // how do I make these work the other way around?
			res = append([]byte{1}, res...)
		} else {
			res = append([]byte{0}, res...)
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

func main() {
	// databytes := gettilefromfile("ROMs/136043-1119.16s", 50)
	// fmt.Printf("byte(s): %02x\n", databytes)

	planedata := [4][]byte{}

	planedata[0] = gettilefromfile("ROMs/136043-1119.16s", 0xcfc)
	planedata[1] = gettilefromfile("ROMs/136043-1119.16s", 0xcfc)
	planedata[2] = gettilefromfile("ROMs/136043-1119.16s", 0xcfc)
	planedata[3] = gettilefromfile("ROMs/136043-1119.16s", 0xcfc)

	fmt.Printf("planes work out to: %d\n", planedata)

	fmt.Printf("indexes work out to: %d\n", mergeplanes(planedata))
}

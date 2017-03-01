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

func gettilefromfile(file string, bytenum int) TilePlane {
	f, err := os.Open(file)
	check(err)

	databytes := make([]byte, 8)

	f.Seek(int64(bytenum), 0)
	_, err = f.Read(databytes)
	check(err)

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

func main() {
	// databytes := gettilefromfile("ROMs/136043-1119.16s", 50)
	// fmt.Printf("byte(s): %02x\n", databytes)

	planes := [4][]byte{}
	planes[0] = bytetobits(0xFF)
	planes[1] = bytetobits(0xFF)
	planes[2] = bytetobits(0xFF)
	planes[3] = bytetobits(0xFF)

	fmt.Printf("planes work out to: %d\n", planes)
	indexline := []byte{}

	for i := 0; i < 8; i++ {
		val := (planes[3][i] * byte(math.Pow(2, 3))) + (planes[2][i] * byte(math.Pow(2, 2))) + (planes[1][i] * byte(math.Pow(2, 1))) + (planes[0][i] * byte(math.Pow(2, 0)))
		indexline = append(indexline, val)
	}

	fmt.Printf("indexes work out to: %d\n", indexline)
}

package main

import (
	"encoding/binary"
	"io"
	"os"
)

var slapsticRoms = []string{
	"ROMs/136043-1105.10a",
	"ROMs/136043-1106.10b",
}

var slapsticBankInfo = []int{
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x54, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x95,
	0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xFE, 0xFF,
	0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x03, 0xFC, 0x0E,
}

var sf [2]*os.File // Open slapstic files

const (
	SLAPSTIC_START = 0x038000
)

// Do this the lazy way -- read an oversized chunk, then keep what we need
func slapsticReadMaze(mazenum int) []int {
	addr := slapsticMazeGetRealAddr(mazenum)
	b := slapsticReadBytes(addr, 512)

	var intbuf []int
	for i := 0; true; i++ {
		intbuf = append(intbuf, int(b[i]))
		if i >= 11 && int(b[i]) == 0 {
			break
		}
	}

	//	return b[:i+1]
	return intbuf
}

func slapsticMazeGetRealAddr(mazenum int) int {
	bank := slapsticMazeGetBank(mazenum)
	addr := slapsticReadMazeOffset(mazenum) + (0x2000 * bank)

	// fmt.Printf("Maze real addr: 0x%06x\n", addr)
	return addr
}

func slapsticMazeGetBank(mazenum int) int {
	if mazenum < 0 || mazenum > 116 {
		panic("Invalid maze number requested (must be 0 <= x <= 116")
	}

	// More or less a direct port of the 68k coede. Should improve this.
	offset := mazenum / 4
	bi := slapsticBankInfo[offset]
	offset = (mazenum % 4) * 2
	bi = bi >> uint(offset)
	bi = bi & 0x3

	return bi
}

func slapsticReadMazeOffset(mazenum int) int {
	buf := slapsticReadBytes(0x03800c+(4*mazenum), 4)
	mazeoffset := binary.BigEndian.Uint32(buf)
	// fmt.Printf("Offset for maze: 0x%06x\n", mazeoffset)

	return int(mazeoffset)
}

// Read bytes from combined ROM. Only works if reading an even address
func slapsticReadBytes(offset int, count int) []byte {
	f, err := os.Open(slapsticRoms[0])
	check(err)
	sf[0] = f
	defer sf[0].Close()

	f, err = os.Open(slapsticRoms[1])
	check(err)
	sf[1] = f
	defer sf[1].Close()

	if offset >= SLAPSTIC_START {
		offset -= SLAPSTIC_START
	}

	// fmt.Printf("offset to load is: %06x\n", offset)

	sf[0].Seek(int64(offset/2), 0)
	sf[1].Seek(int64(offset/2), 0)

	var buf []byte
	var b = make([]byte, 1)

	for i := 0; i < count; i++ {
		// If we're starting on an odd byte, make it look like we've
		// already read an even one.
		if (i == 0) && (offset%2 > 0) {
			sf[0].Read(b)
			i++
			count++
		}
		s, err := sf[i%2].Read(b)
		if s == 0 && err == io.EOF {
			break
		}
		check(err)
		buf = append(buf, b[0])
	}

	return buf
}

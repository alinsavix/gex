package main

import (
	"io"
	"os"
)

// Read bytes from combined ROM. Only works if reading an even address
func romSplitRead(roms []string, offset int, count int) []byte {
	var handles [2]*os.File // Open slapstic files

	f, err := os.Open(roms[0])
	check(err)
	handles[0] = f
	defer handles[0].Close()

	f, err = os.Open(roms[1])
	check(err)
	handles[1] = f
	defer handles[1].Close()

	if offset >= SLAPSTIC_START {
		offset -= SLAPSTIC_START
	}

	handles[0].Seek(int64(offset/2), 0)
	handles[1].Seek(int64(offset/2), 0)

	var buf []byte
	var b = make([]byte, 1)

	for i := 0; i < count; i++ {
		// If we're starting on an odd byte, make it look like we've
		// already read an even one.
		if (i == 0) && (offset%2 > 0) {
			handles[0].Read(b)
			i++
			count++
		}
		s, err := handles[i%2].Read(b)
		if s == 0 && err == io.EOF {
			break
		}
		check(err)
		buf = append(buf, b[0])
	}

	return buf
}

package main

import (
	"os"
)

var codeRoms = [][]string{
	{
		"ROMs/136043-1109.7a",
		"ROMs/136043-1110.7b",
	},
	{
		"ROMs/136043-1121.6a",
		"ROMs/136043-1122.6b",
	},
}

var codeRomSets = []Romset{
	{
		offset: 0x8000,
		roms:   codeRoms[0],
	},
	{
		offset: 0x0000,
		roms:   codeRoms[0],
	},
	{
		offset: 0x8000,
		roms:   codeRoms[1],
	},
	{
		offset: 0x0000,
		roms:   codeRoms[1],
	},
}

func coderomGetByAddr(addr int) (int, []string) {
	a := addr - 0x040000
	bank := a / 0x8000
	rs := codeRomSets[bank]
	offset := a%0x8000 + rs.offset

	return offset, rs.roms
}

var rf [2]*os.File // Open rom files

const (
	ROM_START = 0x040000
)

// func coderomGetRomOffset(addr int) (int, int) {
// 	addr -= ROM_START
// 	romnum := addr / 0x010000
// 	offset := addr - (romnum * 0x010000)

// 	return romnum, offset
// }

// FIXME: Consolidate with slapsticReadBytes
func coderomGetBytes(addr int, count int) []byte {
	offset, rompair := coderomGetByAddr(addr)
	buf := romSplitRead(rompair, offset, count) // FIXME: Standardize arg order

	return buf
}

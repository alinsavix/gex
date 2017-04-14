package main

const (
	CODE_ROM_START = 0x040000
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

func coderomGetByAddr(addr int) ([]string, int) {
	a := addr - CODE_ROM_START
	bank := a / 0x8000
	rs := codeRomSets[bank]
	offset := a%0x8000 + rs.offset

	return rs.roms, offset
}

func coderomGetBytes(addr int, count int) []byte {
	rompair, offset := coderomGetByAddr(addr)
	buf := romSplitRead(rompair, offset, count) // FIXME: Standardize arg order

	return buf
}

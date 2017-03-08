package main

var ghostAnims = map[string]map[string][]int{
	"walk": {
		"up":        {0x890, 0x899, 0x8a2, 0x8ab},
		"upright":   {0x86c, 0x875, 0x87e, 0x887},
		"right":     {0x848, 0x851, 0x85a, 0x863},
		"downright": {0x824, 0x82d, 0x836, 0x83f},
		"down":      {0x800, 0x809, 0x812, 0x81b},
		"downleft":  {0x900, 0x909, 0x912, 0x91b},
		"left":      {0x8d8, 0x8e1, 0x8ea, 0x8f3},
		"upleft":    {0x8b4, 0x8bd, 0x8c6, 0x8cf},
	},
}

// ghost -> walk -> up

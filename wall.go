package main

import (
	"fmt"
	"image/gif"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reWallNum = regexp.MustCompile(`^wall(\d+)$`)
var reWallColor = regexp.MustCompile(`^c(\d+)$`)
var reWallAdj = regexp.MustCompile(`^(u|ur|r|dr|d|dl|l|ul)$`)

func dowall(arg string) {
	split := strings.Split(arg, "-")

	var wallNum = -1
	var wallColor = 0
	var wallAdj = 0

	// Still wonder if there's a cleaner way
	for _, ss := range split {
		switch {
		case reWallNum.MatchString(ss):
			wall, _ := strconv.ParseInt(reWallNum.FindStringSubmatch(ss)[1], 10, 0)
			wallNum = int(wall)
		case reWallColor.MatchString(ss):
			color, _ := strconv.ParseInt(reWallColor.FindStringSubmatch(ss)[1], 10, 0)
			wallColor = int(color)
		case reWallAdj.MatchString(ss):
			res := reWallAdj.FindStringSubmatch(ss)
			if res[1] != "" {
				switch res[1] {
				case "ul":
					wallAdj += 0x01
				case "u":
					wallAdj += 0x02
				case "ur":
					wallAdj += 0x04
				case "l":
					wallAdj += 0x08
				case "r":
					wallAdj += 0x10
				case "dl":
					wallAdj += 0x20
				case "d":
					wallAdj += 0x40
				case "dr":
					wallAdj += 0x80
				}
			}
		}

	}
	fmt.Printf("Wall number: %d   color: %d   adj: %d\n", wallNum, wallColor, wallAdj)

	t := make([]int, 4)
	for i := 0; i < 4; i++ {
		// t[i] = (floorNum * 48) + floorStamps[floorAdj][i]
		m := wallMap[wallAdj]
		t[i] = wallStamps[(68*wallNum)+m][i]
	}

	// These need to be done better
	opts.PalType = "wall"
	opts.PalNum = wallColor

	img := genimage_fromarray(t, 2, 2)
	f, _ := os.OpenFile(opts.Output, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.Encode(f, img, &gif.Options{NumColors: 16})

}

var wallStamps = [][]int{
	// 0
	{0x1C1, 0x1C2, 0x1C3, 0x1C4},
	{0x1C5, 0x1C5, 0x1C6, 0x1C6},
	{0x1C7, 0x1C8, 0x1C6, 0x1C6},
	{0x1C9, 0x1CA, 0x1C9, 0x1CA},
	{0x1C5, 0x1CB, 0x1C6, 0x1CC},
	{0x1CD, 0x1CE, 0x1C9, 0x1CA},
	{0x1C5, 0x1C5, 0x1CF, 0x1D0},
	{0x1C9, 0x1C8, 0x1C9, 0x1D0},
	{0x1C7, 0x1CA, 0x1CF, 0x1CA},
	{0x1D1, 0x1C5, 0x1C9, 8},
	{0x1C5, 0x1C5, 8, 8},
	{0x1C5, 0x1D2, 8, 0x1CA},
	{0x1C9, 8, 0x1C9, 8},
	{8, 8, 8, 8},
	{8, 0x1CA, 8, 0x1CA},
	{0x1C9, 8, 0x1D3, 0x1C6},
	{8, 8, 0x1C6, 0x1C6},
	{8, 0x1CA, 0x1C6, 0x1D4},
	{0x1D1, 0x1D2, 0x1D3, 0x1D4},
	{0x1D5, 0x1D6, 0x1D7, 0x1D8},
	{0x1D1, 0x1C5, 0x1C9, 0x1D0},
	{0x1C5, 0x1D2, 0x1D9, 0x1CA},
	{0x1C9, 0x1C8, 0x1D3, 0x1C6},
	{0x1DA, 0x1CA, 0x1C6, 0x1DB},
	{8, 8, 8, 0x1D0},
	{8, 8, 0x1CF, 8},
	{8, 0x1C8, 8, 8},
	{0x1DA, 8, 8, 8},
	{0x1C9, 0x1C8, 0x1DC, 0x1D0},
	{0x1C9, 0x1CA, 0x1DC, 0x1CA},
	{0x1CD, 0x1CE, 0x1DC, 0x1CA},
	{0x1C5, 0x1C5, 0x1DD, 0x1C6},
	{0x1D1, 0x1C5, 0x1DC, 8},
	{0x1DA, 8, 0x1C6, 0x1C6},
	{8, 0x1CA, 0x1DE, 0x1CA},
	{8, 0x1DF, 0x1C6, 0x1C6},
	{0x1C9, 8, 0x1C9, 0x1D0},
	{0x1C9, 8, 0x1E0, 8},
	{0x1C7, 0x1CA, 0x1DD, 0x1DB},
	{0x1D1, 0x1C5, 0x1DC, 0x1D0},
	{0x1C7, 0x1C8, 0x1CF, 0x1D0},
	{8, 0x1C8, 0x1CF, 0x1D0},
	{0x1C7, 8, 0x1CF, 0x1D0},
	{0x1C7, 0x1C8, 8, 0x1D0},
	{0x1C7, 0x1C8, 0x1CF, 8},
	{8, 8, 0x1CF, 0x1D0},
	{0x1C7, 0x1C8, 8, 8},
	{8, 0x1C8, 8, 0x1D0},
	{0x1C7, 8, 0x1CF, 8},
	{8, 0x1C8, 0x1CF, 8},
	{0x1C7, 8, 8, 0x1D0},
	{0x1C1, 0x1C2, 0x1E1, 0x1C4},
	{0x1E2, 0x1E3, 0x1E4, 0x1D8},
	{8, 8, 0x1DD, 0x1C6},
	{0x1D1, 0x1D2, 0x1E5, 0x1D4},
	{0x1C5, 0x1CB, 0x1DD, 0x1CC},
	{8, 0x1CA, 0x1DD, 0x1D4},
	{0x1C9, 0x1C8, 0x1E5, 0x1C6},
	{0x1C9, 8, 0x1E5, 0x1C6},
	{0x1C7, 0x1C8, 0x1DD, 0x1C6},
	{8, 0x1DF, 0x1DD, 0x1C6},
	{0x1DA, 8, 0x1DD, 0x1C6},
	{0x1C7, 0x1CA, 8, 0x1CA},
	{0x1C9, 8, 0x1DC, 0x1D0},
	{0x1C5, 0x1C5, 8, 0x1D0},
	{0x1C9, 0x1C8, 0x1C9, 8},
	{0x1C5, 0x1C5, 0x1CF, 8},
	{0x1C9, 0x1C8, 0x1DC, 8},

	// 1
	{0x1E6, 0x1E7, 0x1E8, 0x1E9},
	{0x1EA, 0x1EA, 0x1EB, 0x1EC},
	{0x1ED, 0x1EE, 0x1EB, 0x1EC},
	{0x1EF, 0x1F0, 0x1F1, 0x1F2},
	{0x1EA, 0x1F3, 0x1F4, 0x1F5},
	{0x1F6, 0x1F7, 0x1F1, 0x1F8},
	{0x1EA, 0x1EA, 0x1F9, 0x1FA},
	{0x1FB, 0x1EE, 0x1FC, 0x1FA},
	{0x1ED, 0x1FD, 0x1FE, 0x1FF},
	{0x200, 0x201, 0x1F1, 0x202},
	{0x201, 0x201, 0x202, 0x202},
	{0x1EA, 0x203, 0x204, 0x1FF},
	{0x1EF, 0x205, 0x1F1, 0x202},
	{0x205, 0x205, 0x202, 0x202},
	{0x202, 0x1FD, 0x204, 0x1FF},
	{0x1FB, 0x202, 0x206, 0x207},
	{0x202, 0x202, 0x208, 0x207},
	{0x202, 0x1FD, 0x208, 0x209},
	{0x20A, 0x20B, 0x206, 0x209},
	{0x20C, 0x20D, 0x20E, 0x20F},
	{0x20A, 0x1EA, 0x1FC, 0x1FA},
	{0x1EA, 0x210, 0x1F9, 0x1FF},
	{0x1FB, 0x211, 0x206, 0x207},
	{0x212, 0x1FD, 0x208, 0x209},
	{0x205, 0x205, 0x202, 0x213},
	{0x205, 0x205, 0x214, 0x202},
	{0x205, 0x215, 0x202, 0x202},
	{0x216, 0x217, 0x202, 0x202},
	{0x1FB, 0x1EE, 0x218, 0x219},
	{0x1EF, 0x21A, 0x21B, 0x21C},
	{0x21D, 0x21E, 0x21F, 0x1F8},
	{0x1EA, 0x1EA, 0x220, 0x207},
	{0x200, 0x201, 0x21F, 0x202},
	{0x212, 0x202, 0x208, 0x207},
	{0x202, 0x1FD, 0x1F9, 0x1FF},
	{0x202, 0x221, 0x208, 0x207},
	{0x1FB, 0x202, 0x1FC, 0x219},
	{0x1EF, 0x205, 0x21F, 0x202},
	{0x212, 0x1FD, 0x220, 0x209},
	{0x20A, 0x1EA, 0x218, 0x219},
	{0x222, 0x1EE, 0x1F9, 0x219},
	{0x202, 0x1EE, 0x1F9, 0x219},
	{0x222, 0x202, 0x1F9, 0x219},
	{0x223, 0x215, 0x202, 0x224},
	{0x223, 0x215, 0x214, 0x202},
	{0x225, 0x225, 0x1F9, 0x219},
	{0x226, 0x227, 0x202, 0x202},
	{0x205, 0x228, 0x202, 0x224},
	{0x226, 0x205, 0x214, 0x202},
	{0x229, 0x227, 0x214, 0x202},
	{0x226, 0x205, 0x202, 0x224},
	{0x1E6, 0x1E7, 0x22A, 0x1E9},
	{0x22B, 0x20D, 0x22C, 0x20F},
	{0x202, 0x202, 0x220, 0x207},
	{0x20A, 0x22D, 0x22E, 0x22F},
	{0x1EA, 0x1F3, 0x230, 0x1F5},
	{0x202, 0x231, 0x220, 0x22F},
	{0x1FB, 0x1EE, 0x232, 0x207},
	{0x1FB, 0x202, 0x22E, 0x207},
	{0x1ED, 0x1EE, 0x220, 0x1EC},
	{0x202, 0x221, 0x220, 0x207},
	{0x212, 0x202, 0x220, 0x207},
	{0x1ED, 0x1FD, 0x204, 0x1FF},
	{0x1FB, 0x202, 0x218, 0x219},
	{0x233, 0x233, 0x202, 0x224},
	{0x1EF, 0x215, 0x1F1, 0x202},
	{0x201, 0x201, 0x214, 0x202},
	{0x1EF, 0x234, 0x21F, 0x202},

	// 2
	{0x235, 0x236, 0x237, 0x238},
	{0x239, 0x23A, 0x23B, 0x23C},
	{0x23D, 0x23E, 0x23F, 0x23F},
	{0x240, 0x241, 0x242, 0x241},
	{0x243, 0x244, 0x245, 0x246},
	{0x247, 0x248, 0x249, 0x24A},
	{0x243, 0x24B, 0x24C, 0x24D},
	{0x24E, 0x24F, 0x250, 0x251},
	{0x252, 0x253, 0x24C, 0x254},
	{0x255, 0x256, 0x257, 0x258},
	{0x259, 0x256, 0x25A, 0x258},
	{0x25B, 0x25C, 0x25D, 0x25E},
	{0x257, 0x258, 0x257, 0x258},
	{0x25A, 0x25F, 0x25A, 0x258},
	{0x260, 0x261, 0x25A, 0x261},
	{0x262, 0x25F, 0x263, 0x264},
	{0x25A, 0x258, 0x265, 0x266},
	{0x267, 0x268, 0x23F, 0x269},
	{0x26A, 0x26B, 0x26C, 0x26D},
	{0x26E, 0x26F, 0x1D7, 0x270},
	{0x26A, 0x271, 0x250, 0x251},
	{0x243, 0x272, 0x273, 0x274},
	{0x275, 0x24F, 0x26C, 0x276},
	{0x277, 0x278, 0x245, 0x279},
	{0x267, 0x267, 0x27A, 0x27B},
	{0x25A, 0x25F, 0x27C, 0x27D},
	{0x260, 0x27E, 0x25A, 0x25F},
	{0x27F, 0x25F, 0x25A, 0x25F},
	{0x280, 0x24F, 0x281, 0x251},
	{0x257, 0x282, 0x283, 0x282},
	{0x284, 0x285, 0x286, 0x287},
	{0x23A, 0x23A, 0x288, 0x289},
	{0x28A, 0x256, 0x283, 0x28B},
	{0x28C, 0x267, 0x28D, 0x28D},
	{0x25A, 0x282, 0x28E, 0x28F},
	{0x290, 0x291, 0x292, 0x276},
	{0x293, 0x294, 0x295, 0x24D},
	{0x257, 0x258, 0x296, 0x25F},
	{0x297, 0x298, 0x299, 0x29A},
	{0x29B, 0x271, 0x29C, 0x251},
	{0x29D, 0x24F, 0x24C, 0x251},
	{0x267, 0x29E, 0x29F, 0x2A0},
	{0x2A1, 0x267, 0x29F, 0x2A0},
	{0x2A1, 0x29E, 0x2A2, 0x2A3},
	{0x2A1, 0x29E, 0x29F, 0x2A4},
	{0x267, 0x267, 0x29F, 0x2A5},
	{0x2A6, 0x2A7, 0x2A8, 0x2A9},
	{0x290, 0x24F, 0x2AA, 0x2AB},
	{0x2AC, 0x267, 0x27C, 0x2A4},
	{0x267, 0x29E, 0x29F, 0x2A4},
	{0x2AD, 0x258, 0x2AE, 0x2AF},
	{0x235, 0x236, 0x2B0, 0x2B1},
	{0x2B2, 0x2B3, 0x1E4, 0x270},
	{0x267, 0x267, 0x288, 0x289},
	{0x2B4, 0x2B5, 0x2B6, 0x2B7},
	{0x2B8, 0x2B9, 0x2BA, 0x2BB},
	{0x267, 0x2BC, 0x2BD, 0x2BE},
	{0x2BF, 0x24F, 0x2B6, 0x2C0},
	{0x2C1, 0x258, 0x2C2, 0x2C3},
	{0x252, 0x2C4, 0x2C5, 0x2C6},
	{0x267, 0x2C7, 0x2C8, 0x276},
	{0x27F, 0x25F, 0x2C9, 0x2CA},
	{0x2AD, 0x253, 0x25A, 0x282},
	{0x242, 0x258, 0x29C, 0x24D},
	{0x23A, 0x271, 0x2CB, 0x2CC},
	{0x2C1, 0x27E, 0x2CD, 0x258},
	{0x243, 0x256, 0x24C, 0x2CE},
	{0x242, 0x2CF, 0x283, 0x258},

	// 3
	{0x2D0, 0x2D1, 0x2D2, 0x2D3},
	{0x2D4, 0x2D4, 0x2D5, 0x2D6},
	{0x2D7, 0x2D8, 0x2D5, 0x2D6},
	{0x2D9, 0x2DA, 0x2DB, 0x2DC},
	{0x2DD, 0x2DE, 0x2DF, 0x2E0},
	{0x2E1, 0x2E2, 0x2DB, 0x2DC},
	{0x2D4, 0x2D4, 0x2E3, 0x2E4},
	{0x2D9, 0x2E5, 0x2DB, 0x2E4},
	{0x2E6, 0x2DA, 0x2E3, 0x2DC},
	{0x2E7, 0x2E8, 0x2DB, 0x2E9},
	{0x2D4, 0x2D4, 0x2EA, 0x2E9},
	{0x2EB, 0x2EC, 0x2EA, 0x2DC},
	{0x2D9, 0x2ED, 0x2DB, 0x2E9},
	{0x2EE, 0x2ED, 0x2EA, 0x2E9},
	{0x2EE, 0x2DA, 0x2EA, 0x2DC},
	{0x2D9, 0x2ED, 0x2EF, 0x2F0},
	{0x2EE, 0x2ED, 0x2D5, 0x2D6},
	{0x2EE, 0x2DA, 0x2F1, 0x2F2},
	{0x2F3, 0x2EC, 0x2EF, 0x2F4},
	{0x2F5, 0x2F6, 0x2F7, 0x2F8},
	{0x2E7, 0x2E8, 0x2DB, 0x2E4},
	{0x2EB, 0x2EC, 0x2E3, 0x2DC},
	{0x2D9, 0x2F9, 0x2EF, 0x2F0},
	{0x2FA, 0x2DA, 0x2F1, 0x2F2},
	{0x2EE, 0x2ED, 0x2EA, 0x2E4},
	{0x2EE, 0x2ED, 0x2E3, 0x2E9},
	{0x2EE, 0x2D8, 0x2EA, 0x2E9},
	{0x2FA, 0x2ED, 0x2EA, 0x2E9},
	{0x2D9, 0x2E5, 0x2FB, 0x2E4},
	{0x2D9, 0x2DA, 0x2FB, 0x2DC},
	{0x2E1, 0x2E2, 0x2FB, 0x2DC},
	{0x2D4, 0x2D4, 0x2FC, 0x2D6},
	{0x2E7, 0x2E8, 0x2FB, 0x2E9},
	{0x2D7, 0x2ED, 0x2D5, 0x2D6},
	{0x2EE, 0x2DA, 0x2E3, 0x2DC},
	{0x2EE, 0x2D8, 0x2D5, 0x2D6},
	{0x2D9, 0x2ED, 0x2DB, 0x2E4},
	{0x2D9, 0x2ED, 0x2FB, 0x2E9},
	{0x2D7, 0x2DA, 0x2FC, 0x2F2},
	{0x2E7, 0x2E8, 0x2FB, 0x2E4},
	{0x2D7, 0x2D8, 0x2E3, 0x2E4},
	{0x2EE, 0x2D8, 0x2E3, 0x2E4},
	{0x2D7, 0x2ED, 0x2E3, 0x2E4},
	{0x2D7, 0x2D8, 0x2EA, 0x2E4},
	{0x2D7, 0x2D8, 0x2E3, 0x2E9},
	{0x2EE, 0x2ED, 0x2E3, 0x2E4},
	{0x2D7, 0x2D8, 0x2EA, 0x2E9},
	{0x2EE, 0x2D8, 0x2EA, 0x2E4},
	{0x2D7, 0x2ED, 0x2E3, 0x2E9},
	{0x2EE, 0x2D8, 0x2E3, 0x2E9},
	{0x2D7, 0x2ED, 0x2EA, 0x2E4},
	{0x2D0, 0x2D1, 0x2FD, 0x2D3},
	{0x2F5, 0x2F6, 0x2FE, 0x2F8},
	{0x2EE, 0x2ED, 0x2FC, 0x2D6},
	{0x2F3, 0x2EC, 0x2FF, 0x2F4},
	{0x2DD, 0x2DE, 0x2FC, 0x2E0},
	{0x2EE, 0x2DA, 0x2FC, 0x2F2},
	{0x2D9, 0x2F9, 0x300, 0x2F0},
	{0x2D9, 0x2ED, 0x300, 0x2F0},
	{0x2D7, 0x2D8, 0x2FC, 0x2D6},
	{0x2EE, 0x2D8, 0x2FC, 0x2D6},
	{0x2D7, 0x2ED, 0x2FC, 0x2D6},
	{0x301, 0x2DA, 0x2EA, 0x2DC},
	{0x2D9, 0x2ED, 0x2FB, 0x2E4},
	{0x2D4, 0x2D4, 0x2EA, 0x2E4},
	{0x2D9, 0x2E5, 0x2DB, 0x302},
	{0x2D4, 0x303, 0x2E3, 0x302},
	{0x2D9, 0x2E5, 0x2FB, 0x302},

	// 4
	{0x304, 0x305, 0x306, 0x307},
	{0x308, 0x309, 0x30A, 0x30B},
	{0x30C, 0x30D, 0x30A, 0x30B},
	{0x30E, 0x30F, 0x310, 0x311},
	{0x308, 0x312, 0x30A, 0x313},
	{0x314, 0x315, 0x310, 0x311},
	{0x316, 0x317, 0x318, 0x319},
	{0x30E, 0x31A, 0x310, 0x319},
	{0x31B, 0x30F, 0x318, 0x311},
	{0x31C, 0x317, 0x310, 0x31D},
	{0x31E, 0x31F, 0x320, 0x321},
	{0x322, 0x323, 0x320, 0x311},
	{0x30E, 0x324, 0x310, 0x325},
	{0x326, 0x327, 0x320, 0x325},
	{0x328, 0x30F, 0x320, 0x311},
	{0x329, 0x32A, 0x32B, 0x30B},
	{0x32C, 0x32A, 0x30A, 0x30B},
	{0x32D, 0x32E, 0x30A, 0x32F},
	{0x31C, 0x323, 0x330, 0x32F},
	{0x331, 0x332, 0x333, 0x334},
	{0x31C, 0x317, 0x310, 0x319},
	{0x322, 0x323, 0x318, 0x311},
	{0x329, 0x30D, 0x32B, 0x30B},
	{0x335, 0x32E, 0x30A, 0x32F},
	{0x326, 0x336, 0x320, 0x337},
	{0x326, 0x327, 0x318, 0x325},
	{0x326, 0x338, 0x320, 0x325},
	{0x339, 0x327, 0x320, 0x325},
	{0x30E, 0x31A, 0x33A, 0x319},
	{0x30E, 0x30F, 0x33A, 0x311},
	{0x314, 0x315, 0x33A, 0x311},
	{0x308, 0x309, 0x33B, 0x30B},
	{0x31C, 0x317, 0x33A, 0x31D},
	{0x30C, 0x32A, 0x30A, 0x30B},
	{0x328, 0x30F, 0x318, 0x311},
	{0x32C, 0x30D, 0x30A, 0x30B},
	{0x30E, 0x336, 0x310, 0x337},
	{0x30E, 0x324, 0x33A, 0x325},
	{0x33C, 0x32E, 0x33B, 0x32F},
	{0x31C, 0x317, 0x33A, 0x319},
	{0x339, 0x33D, 0x318, 0x337},
	{0x326, 0x33D, 0x318, 0x337},
	{0x339, 0x336, 0x318, 0x337},
	{0x339, 0x33D, 0x320, 0x337},
	{0x339, 0x338, 0x318, 0x325},
	{0x326, 0x336, 0x318, 0x337},
	{0x339, 0x338, 0x320, 0x325},
	{0x326, 0x33D, 0x320, 0x337},
	{0x339, 0x327, 0x318, 0x325},
	{0x326, 0x338, 0x318, 0x325},
	{0x339, 0x336, 0x320, 0x337},
	{0x304, 0x305, 0x33E, 0x307},
	{0x331, 0x332, 0x33F, 0x334},
	{0x32C, 0x32A, 0x33B, 0x30B},
	{0x340, 0x323, 0x341, 0x342},
	{0x308, 0x312, 0x33B, 0x313},
	{0x32D, 0x32E, 0x33B, 0x32F},
	{0x343, 0x344, 0x345, 0x346},
	{0x343, 0x32A, 0x345, 0x346},
	{0x30C, 0x30D, 0x33B, 0x30B},
	{0x32C, 0x30D, 0x33B, 0x30B},
	{0x30C, 0x32A, 0x33B, 0x30B},
	{0x31B, 0x30F, 0x320, 0x311},
	{0x30E, 0x336, 0x33A, 0x337},
	{0x316, 0x317, 0x320, 0x319},
	{0x30E, 0x347, 0x310, 0x325},
	{0x316, 0x317, 0x318, 0x31D},
	{0x30E, 0x348, 0x33A, 0x325},

	//5
	{0x349, 0x34A, 0x34B, 0x34C},
	{0x34D, 0x34E, 0x34F, 0x350},
	{0x351, 0x352, 0x34F, 0x350},
	{0x353, 0x354, 0x355, 0x356},
	{0x357, 0x358, 0x359, 0x35A},
	{0x35B, 0x35C, 0x355, 0x35D},
	{0x357, 0x35E, 0x35F, 0x360},
	{0x353, 0x361, 0x355, 0x362},
	{0x363, 0x354, 0x35F, 0x356},
	{0x364, 0x365, 0x355, 0x366},
	{0x367, 0x368, 0x369, 0x36A},
	{0x36B, 0x354, 0x369, 0x356},
	{0x353, 0x36C, 0x355, 0x36A},
	{0x36D, 0x36E, 0x36F, 0x370},
	{0x371, 0x354, 0x369, 0x356},
	{0x372, 0x373, 0x374, 0x375},
	{0x376, 0x377, 0x378, 0x379},
	{0x37A, 0x37B, 0x34F, 0x37C},
	{0x37D, 0x37E, 0x37F, 0x380},
	{0x381, 0x382, 0x383, 0x384},
	{0x364, 0x385, 0x355, 0x362},
	{0x386, 0x387, 0x35F, 0x388},
	{0x372, 0x389, 0x374, 0x375},
	{0x38A, 0x37B, 0x34F, 0x37C},
	{0x36D, 0x36E, 0x36F, 0x362},
	{0x36D, 0x36E, 0x35F, 0x38B},
	{0x36D, 0x38C, 0x36F, 0x38B},
	{0x38D, 0x36E, 0x36F, 0x38B},
	{0x353, 0x361, 0x38E, 0x362},
	{0x353, 0x354, 0x38E, 0x356},
	{0x35B, 0x35C, 0x38E, 0x35D},
	{0x34D, 0x34E, 0x38F, 0x350},
	{0x390, 0x391, 0x38E, 0x366},
	{0x392, 0x377, 0x378, 0x379},
	{0x36D, 0x354, 0x35F, 0x356},
	{0x376, 0x352, 0x378, 0x379},
	{0x353, 0x393, 0x355, 0x362},
	{0x353, 0x36C, 0x38E, 0x36A},
	{0x38A, 0x37B, 0x38F, 0x37C},
	{0x364, 0x365, 0x38E, 0x362},
	{0x351, 0x352, 0x35F, 0x360},
	{0x394, 0x352, 0x35F, 0x360},
	{0x351, 0x36E, 0x35F, 0x360},
	{0x351, 0x352, 0x36F, 0x360},
	{0x351, 0x352, 0x35F, 0x370},
	{0x395, 0x371, 0x35F, 0x360},
	{0x351, 0x352, 0x36F, 0x370},
	{0x394, 0x352, 0x36F, 0x396},
	{0x351, 0x36E, 0x35F, 0x370},
	{0x395, 0x352, 0x35F, 0x370},
	{0x351, 0x36E, 0x36F, 0x360},
	{0x397, 0x389, 0x398, 0x399},
	{0x381, 0x382, 0x398, 0x384},
	{0x376, 0x377, 0x38F, 0x379},
	{0x37D, 0x37E, 0x39A, 0x39B},
	{0x357, 0x358, 0x38F, 0x35A},
	{0x37A, 0x37B, 0x38F, 0x37C},
	{0x372, 0x352, 0x39A, 0x375},
	{0x372, 0x36E, 0x39A, 0x375},
	{0x39C, 0x352, 0x38F, 0x379},
	{0x395, 0x352, 0x38F, 0x379},
	{0x39D, 0x377, 0x38F, 0x379},
	{0x39D, 0x354, 0x369, 0x356},
	{0x353, 0x393, 0x38E, 0x362},
	{0x34D, 0x34E, 0x36F, 0x360},
	{0x353, 0x352, 0x355, 0x36A},
	{0x34D, 0x34E, 0x35F, 0x370},
	{0x353, 0x352, 0x38E, 0x370},
}

// This table is for figuring out which wall segment to use, based on which
// squares around it also have walls. It is extracted from 0x5EE24 in the
// Gauntlet II ROMs
//
// Bits that get set for walls in locations relative to this one:
// up left:    0x01
// up:         0x02
// up right:   0x04
// left:       0x08
// right:      0x10
// down left:  0x20
// down:       0x40
// down right: 0x80

var wallMap = []int{
	0x12, 0x12, 0x13, 0x13, 0x12, 0x12, 0x13, 0x13,
	0x4, 0x4, 0x17, 0x11, 0x4, 0x4, 0x17, 0x11,
	0x0, 0x0, 0x16, 0x16, 0x0, 0x0, 0xF, 0xF,
	0x1, 0x1, 0x2, 0x23, 0x1, 0x1, 0x21, 0x10,
	0x36, 0x36, 0x34, 0x34, 0x36, 0x36, 0x34, 0x34,
	0x37, 0x37, 0x26, 0x38, 0x37, 0x37, 0x26, 0x38,
	0x33, 0x33, 0x39, 0x39, 0x33, 0x33, 0x3A, 0x3A,
	0x1F, 0x1F, 0x3B, 0x3C, 0x1F, 0x1F, 0x3D, 0x35,
	0x5, 0x5, 0x3, 0x3, 0x5, 0x5, 0x3, 0x3,
	0x15, 0x15, 0x8, 0x22, 0x15, 0x15, 0x8, 0x22,
	0x14, 0x14, 0x7, 0x7, 0x14, 0x14, 0x24, 0x24,
	0x6, 0x6, 0x28, 0x29, 0x6, 0x6, 0x2A, 0x2D,
	0x1E, 0x1E, 0x1D, 0x1D, 0x1E, 0x1E, 0x1D, 0x1D,
	0xB, 0xB, 0x3E, 0xE, 0xB, 0xB, 0x3E, 0xE,
	0x27, 0x27, 0x1C, 0x1C, 0x27, 0x27, 0x3F, 0x3F,
	0x40, 0x40, 0x2B, 0x2F, 0x40, 0x40, 0x32, 0x18,
	0x12, 0x12, 0x13, 0x13, 0x12, 0x12, 0x13, 0x13,
	0x4, 0x4, 0x17, 0x11, 0x4, 0x4, 0x17, 0x11,
	0x0, 0x0, 0x16, 0x16, 0x0, 0x0, 0xF, 0xF,
	0x1, 0x1, 0x2, 0x23, 0x1, 0x1, 0x21, 0x10,
	0x36, 0x36, 0x34, 0x34, 0x36, 0x36, 0x34, 0x34,
	0x37, 0x37, 0x26, 0x38, 0x37, 0x37, 0x26, 0x38,
	0x33, 0x33, 0x39, 0x39, 0x33, 0x33, 0x3A, 0x3A,
	0x1F, 0x1F, 0x3B, 0x3C, 0x1F, 0x1F, 0x3D, 0x35,
	0x5, 0x5, 0x3, 0x3, 0x5, 0x5, 0x3, 0x3,
	0x15, 0x15, 0x8, 0x22, 0x15, 0x15, 0x8, 0x22,
	0x9, 0x9, 0x41, 0x41, 0x9, 0x9, 0xC, 0xC,
	0x42, 0x42, 0x2C, 0x31, 0x42, 0x42, 0x30, 0x19,
	0x1E, 0x1E, 0x1D, 0x1D, 0x1E, 0x1E, 0x1D, 0x1D,
	0xB, 0xB, 0x3E, 0xE, 0xB, 0xB, 0x3E, 0xE,
	0x20, 0x20, 0x43, 0x43, 0x20, 0x20, 0x25, 0x25,
	0xA, 0xA, 0x2E, 0x1A, 0xA, 0xA, 0x1B, 0xD,
}

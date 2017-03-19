package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var reFloorNum = regexp.MustCompile(`^floor(\d+)`)
var reFloorColor = regexp.MustCompile(`^c(\d+)`)
var reFloorAdj = regexp.MustCompile(`^var(\d+)|(hwall)|(vwall)|(dwall)`)

func dofloor(arg string) {
	split := strings.Split(arg, "-")

	var floorNum = -1
	var floorColor = 0
	var floorAdj = 0

	// Wonder if there's a cleaner way
	for _, ss := range split {
		switch {
		case reFloorNum.MatchString(ss):
			floor, _ := strconv.ParseInt(reFloorNum.FindStringSubmatch(ss)[1], 10, 0)
			floorNum = int(floor)
		case reFloorColor.MatchString(ss):
			color, _ := strconv.ParseInt(reFloorColor.FindStringSubmatch(ss)[1], 10, 0)
			floorColor = int(color)
		case reFloorAdj.MatchString(ss):
			res := reFloorAdj.FindStringSubmatch(ss)
			if res[1] != "" {
				adj, _ := strconv.ParseInt(res[1], 10, 0)
				floorAdj += int(adj)
			}
			if res[2] != "" {
				floorAdj += 4
			}
			if res[3] != "" {
				floorAdj += 16
			}
			if res[4] != "" {
				floorAdj += 8
			}

		}

	}
	fmt.Printf("Floor number: %d   color: %d    adj: %d\n", floorNum, floorColor, floorAdj)

	// t := floorGetTiles(floorNum, floorAdj)
	stamp := floorGetStamp(floorNum, floorAdj, floorColor)

	img := blankimage(2*8, 2*8)
	writestamptoimage(img, stamp, 0, 0)
	savetopng(opts.Output, img)
}

func floorGetTiles(floorNum int, floorAdj int) []int {
	t := make([]int, 4)
	for i := 0; i < 4; i++ {
		t[i] = (floorNum * 48) + floorStamps[floorAdj][i]
	}

	return t
}

func floorGetStamp(floorNum int, floorAdj int, floorColor int) *Stamp {
	tiles := floorGetTiles(floorNum, floorAdj)
	stamp := genstamp_fromarray(tiles, 2, "floor", floorColor)

	return stamp
}

var floorStamps = [][]int{
	{0x11, 0x12, 0x13, 0x14},
	{0x15, 0x16, 0x17, 0x18},
	{0x19, 0x1A, 0x1B, 0x1C},
	{0x1D, 0x1E, 0x1F, 0x20},
	{0x21, 0x12, 0x22, 0x14},
	{0x23, 0x16, 0x24, 0x18},
	{0x25, 0x1A, 0x26, 0x1C},
	{0x27, 0x1E, 0x28, 0x20},
	{0x11, 0x12, 0x29, 0x14},
	{0x15, 0x16, 0x2A, 0x18},
	{0x19, 0x1A, 0x2B, 0x1C},
	{0x1D, 0x1E, 0x2C, 0x20},
	{0x2D, 0x12, 0x29, 0x14},
	{0x2E, 0x16, 0x2A, 0x18},
	{0x2F, 0x1A, 0x2B, 0x1C},
	{0x30, 0x1E, 0x2C, 0x20},
	{0x11, 0x12, 0x31, 0x32},
	{0x15, 0x16, 0x33, 0x34},
	{0x19, 0x1A, 0x35, 0x36},
	{0x1D, 0x1E, 0x37, 0x38},
	{0x21, 0x12, 0x39, 0x32},
	{0x23, 0x16, 0x3A, 0x34},
	{0x25, 0x1A, 0x3B, 0x36},
	{0x27, 0x1E, 0x3C, 0x38},
	{0x11, 0x12, 0x29, 0x3D},
	{0x15, 0x16, 0x2A, 0x3E},
	{0x19, 0x1A, 0x2B, 0x3F},
	{0x1D, 0x1E, 0x2C, 0x40},
	{0x2D, 0x12, 0x29, 0x3D},
	{0x2E, 0x16, 0x2A, 0x3E},
	{0x2F, 0x1A, 0x2B, 0x3F},
	{0x30, 0x1E, 0x2C, 0x40},
}

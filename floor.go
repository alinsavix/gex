package main

import (
	"fmt"
	"image/gif"
	"os"
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

	t := make([]int, 4)
	for i := 0; i < 4; i++ {
		t[i] = (floorNum * 48) + floorStamps[floorAdj][i]
	}

	// These need to be done better
	opts.PalType = "floor"
	opts.PalNum = floorColor

	img := genimage_fromarray(t, 2, 2)
	f, _ := os.OpenFile(opts.Output, os.O_WRONLY|os.O_CREATE, 0600)
	defer f.Close()
	gif.Encode(f, img, &gif.Options{NumColors: 16})

}

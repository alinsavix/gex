package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func mazeMetaPrint(maze *Maze) {
	fmt.Printf("  Encoded length: %3d bytes\n", maze.encodedbytes)
	fmt.Printf("  Wall pattern: %02d, Wall color: %02d     Floor pattern: %02d, Floor color: %02d\n",
		maze.wallpattern, maze.wallcolor, maze.floorpattern, maze.floorcolor)
	fmt.Printf("  Flags: ")
	for k, v := range mazeFlagStrings {
		if (maze.flags & k) != 0 {
			fmt.Printf("%s ", v)
		}
	}
	fmt.Printf("\n  Random food adds: %d\n", (maze.flags&LFLAG3_RANDOMFOOD_MASK)>>8)
	fmt.Printf("  Secret trick: %2d - %s\n", maze.secret, mazeSecretStrings[maze.secret])
}

var reMazeNum = regexp.MustCompile(`^maze(\d+)`)
var reMazeMeta = regexp.MustCompile(`^meta$`)

func domaze(arg string) {
	split := strings.Split(arg, "-")

	var mazeNum = -1
	var mazeMeta = 0

	for _, ss := range split {
		switch {
		case reMazeNum.MatchString(ss):
			m, _ := strconv.ParseInt(reMazeNum.FindStringSubmatch(ss)[1], 10, 0)
			mazeNum = int(m)
			if mazeNum < 0 || mazeNum > 117 {
				panic("Invalid maze number specified.")
			}

		case reMazeMeta.MatchString(ss):
			mazeMeta = 1
		}
	}

	fmt.Printf("Maze number: %d\n", mazeNum)
	maze := mazeDecompress(slapsticReadMaze(mazeNum), mazeMeta > 0)

	if opts.Verbose || mazeMeta > 0 {
		mazeMetaPrint(maze)
	}

	if mazeMeta == 0 {
		genpfimage(maze)
	}
}

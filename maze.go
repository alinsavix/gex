package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var reMazeNum = regexp.MustCompile(`^maze(\d+)`)

func mazeMetaDecode(maze *Maze) {
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

func domaze(arg string) {
	var mazenum int

	m := reMazeNum.FindStringSubmatch(arg)
	if len(m[1]) <= 0 {
		panic("Illegal maze number specified")
	} else {
		mn, _ := strconv.ParseInt(m[1], 10, 0)
		mazenum = int(mn)
	}

	fmt.Printf("Maze number: %d\n", mazenum)
	maze := mazeDecompress(slapsticReadMaze(mazenum))

	if opts.Verbose {
		mazeMetaDecode(maze)
	}

	genpfimage(maze)
}

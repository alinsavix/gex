package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var reMazeNum = regexp.MustCompile(`^maze(\d+)`)

func mazeMetaDecode(maze *Maze) {
	// for k, v := range mazeLflag1strings {
	// 	if (maze.flags1 & k) != 0 {
	// 		fmt.Printf("%s ", v)
	// 	}
	// }
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
	genpfimage(maze)

	if opts.Verbose {
		mazeMetaDecode(maze)
		//		fmt.Printf("flags1: %08b       flags2: %08b\n", maze.flags1, maze.flags2)
		//		fmt.Printf("flags3: %08b       flags4: %08b\n", maze.flags3, maze.flags4)
	}
}

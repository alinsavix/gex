package main

import "github.com/jessevdk/go-flags"
import "os"

var opts struct {
	Animate bool   `short:"a" long:"animate" description:"Animate monster"`
	PalType string `long:"pt" default:"base" description:"Palette type"`
	PalNum  int    `long:"pn" default:"0" base:"16" description:"Palette number (in hex)"`
	Tile    int    `short:"t" long:"tile" base:"16" description:"Tile number to render (in hex)"`
	DimX    int    `short:"x" default:"2" description:"X dimension, in tiles"`
	DimY    int    `short:"y" default:"2" description:"Y dimension, in tiles"`
	Output  string `short:"o" long:"output" default:"test.gif" description:"Output file"`
	Monster string `short:"m" long:"monster" description:"Monster to render"`
	Floor   int    `short:"f" long:"floor" default:"-1" base:"16" description:"Floor stamp to render (in hex)"`
	Wall    int    `short:"w" long:"wall" default:"-1" base:"16" description:"Wall stamp to render (in hex)"`
}

func gexinit() {
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
		// check(err)
	}
}

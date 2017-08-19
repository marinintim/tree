package main

import (
	"flag"
	"log"
	"path/filepath"
	"strings"

	"github.com/marinintim/tree/lib"
)

var Path string
var Graph bool
var Wide bool
var Hidden bool
var Ignore string

func init() {
	flag.BoolVar(&Graph, "graph", true, "draw tree")
	flag.BoolVar(&Wide, "wide", false, "should tree be wider than normal")
	flag.BoolVar(&Hidden, "hidden", false, "include hidden files (start with dot)")
	flag.StringVar(&Ignore, "ignore", "node_modules", "directories to ignore separated by comma")
}

func main() {
	flag.Parse()
	Path = flag.Arg(0)
	Path, err := filepath.Abs(Path)
	if err != nil {
		log.Fatalf("While Absing the argument: %v", err)
	}
	toIgnore := strings.Split(Ignore, ",")
	tree.List(Path, Graph, Wide, Hidden, toIgnore)
}

package tree

import (
	"fmt"
)

var wide bool
var graph bool

// Draw root character and call Display
// for actual rendering of the tree
func DisplayFromRoot(root *Entry, isgraph bool, iswide bool) {
	wide = iswide
	graph = isgraph

	if graph {
		fmt.Println("┬")
	}
	Display(root, []bool{})
}

// Render node and recursively call Display
// to render children.
// Levels []bool is used to keep track of whether
// we need to draw | or space -- if there are no branches
// after that at level or two up, then we need to draw space.
// After printing line levels is updated with IsLast
// from node. IsLast is tracking whether the node is the
// last node on its level below its parent.
func Display(n *Entry, levels []bool) {
	s := getTreeString(n, levels, n.IsLast)
	fmt.Println(s)

	levels = append(levels, n.IsLast)
	for _, e := range n.Children {
		Display(e, levels)
	}
}

/*
Tree looks something like this:
┬
└┬root
 ├┬dir1
 │├─sub1
 │├┬sub2
 ││├─sub1
 ││├─sub2
 ││└─sub3
 │└─sub3
 └─dir2
*/

// Forms
func getTreePrefix(levels []bool) string {
	s := ""
	for _, isLast := range levels {
		if isLast {
			s += " "
		}
		if !isLast {
			s += "│"
		}
		if wide {
			s += " "
		}

	}

	return s
}

func getTreeLeaf(isLast bool) string {
	s := ""
	if isLast {
		s += "└"
	} else {
		s += "├"
	}
	if wide {
		s += "─"
	}
	return s
}

func getTreeString(e *Entry, levels []bool, isLast bool) string {
	if !graph {
		return e.Name
	}

	s := ""

	s += getTreePrefix(levels)
	s += getTreeLeaf(isLast)

	if len(e.Children) > 0 {
		s += "┬"
	} else {
		s += "─"
	}
	if wide {
		s += "─"
	}

	if e.Parent != nil {
		s += dropPrefix(e.Name, e.Parent.Name)
	} else {
		s += e.Name
	}
	return s
}

func dropPrefix(n string, p string) string {
	return n[len(p)+1:]
}

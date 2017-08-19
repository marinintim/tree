tree
====

This document is describing [tree](https://github.com/marinintim/tree) according to [DesignDocManifesto v.1](https://marinin.xyz/projects/ddm#1).

## Purpose

tree is a command utility that lists files and directories recursively, printing them with tree-like graphics.

For example, when called upon `test` directory `tree` could output this:

```
┬
└┬/Users/mt/Workspace/src/github.com/marinintim/tree/test
 ├┬dir1
 │├─sub1
 │├┬sub2
 ││├─sub1
 ││├─sub2
 ││└─sub3
 │└─sub3
 └─dir2
```

Such tool is useful like an overview of a project. (in cases where you don't have a lot of files).

Another purpose was to play with Golang, therefore `tree` is built in Golang.

## Domain

`tree` is dealing with files and directories on one hand, and printing a tree-like structure to the terminal on the other hand.

### Files and directories

On Unix-like systems directories are files from program's perspective. Golang standard library provides `ioutil.ReadDir` call that
returns `[]os.FileInfo`, which in turns contains `Name()` and `IsDir()` that allow to implement `tree` functionality.

### Printing tree-like structure

File trees are graphs. That means we can use different graph algorithms at hand, but the DFS is already provided by Golang calls structure
when you do recursive calls, so DFS was used as an algorithm to walk the tree.


## Structures & Plan

The general plan is to:

1. Generate a tree,
2. Display a tree.

To provide somewhat terminal tool-like functionality there are different command-line flags that affect
how the tree should be generated and displayed. Namely, `-hidden` and `-ignore` affects tree generation,
while `-graph` and `-wide` affect render.

### Mental structure: Tree

Tree is not defined anywhere in the code explicitly, it's assumed
that Entries form a tree.

### Structure: Entry

`Entry` is a struct defined in `lib/tree.go` that represents a node
in final listing, containing references to its children.

We define level as a list of entries which have the same depth.
If we were to draw the tree from left to right, entries which
share the level would have the same indentation.

We store `Depth` for each Entry that starts from 0 for root Entry
and is incremented by one with each level.

### Mental Structure: Root

Root is an Entry that is created from the path supplied by user.

@TRACK(main.go): It is assumed that user supplied a path to a directory and is not verified at the moment.

### Structure: `levels []bool`

Render uses an `levels []bool` to track whether the tree has
entries somewhere under the current level, each
`bool` designates whether the branch at level `n` has 
had ended already.

### Generating a tree

The main action is contained in `expand` function (`lib/expand.go`).
It calls `ioutil.ReadDir` on the name of the current node.

It also uses supplied preferences (like `hidden` and `ignore`)
to filter out unwanted entries.

It also checks whethen the new entry represent a directory, in which
case it recurses on the new entry.

### Rendering a tree

@TRACK(lib/draw.go) TBD.

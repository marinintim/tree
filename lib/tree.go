package tree

var hidden bool
var ignore map[string]bool

type Entry struct {
	Name     string
	IsDir    bool
	Depth    uint8
	Children []*Entry
	Parent   *Entry
	IsLast   bool
}

func List(dir string, graph bool, wide bool, hidden bool, toIgnore []string) {
	ignore = make(map[string]bool)
	for _, s := range toIgnore {
		ignore[s] = true
	}

	root := Entry{dir, true, 0, nil, nil, true}
	expand(&root, hidden)
	DisplayFromRoot(&root, graph, wide)
}

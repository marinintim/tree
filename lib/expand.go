package tree

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

func expand(n *Entry, hidden bool) {
	files, err := ioutil.ReadDir(n.Name)
	if err != nil {
		log.Printf("Error while readdiring %s: %v", n.Name, err)
		return
	}

	l := len(files) - 1
	for i, f := range files {
		name := f.Name()
		if !hidden && strings.HasPrefix(name, ".") {
			continue
		}

		if _, ok := ignore[name]; ok {
			continue
		}

		e := Entry{
			filepath.Join(n.Name, f.Name()),
			f.IsDir(),
			n.Depth + 1,
			nil,
			n,
			i == l,
		}

		if f.IsDir() {
			expand(&e, hidden)
		}

		n.Children = append(n.Children, &e)
	}
}

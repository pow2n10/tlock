package treelock

import (
	"errors"
	"time"
)

type TreeStats struct {
	Total   int
	Locked  int
	Relased int
	Failed  int
}

type Tree struct {
	root  *Node
	stats TreeStats
}

func NewTree() *Tree {
	tree := new(Tree)
	tree.root = NewNode("")
	return tree
}

func (tree *Tree) Lock(path string, wait time.Time) bool {

	return false
}

func (tree *Tree) Unlock(path string) bool {
	return false
}

func (tree *Tree) Try(path string) (string, bool) {
	return "", false
}

func (tree *Tree) getNode(path string) (*Node, error) {
	ps, err := ParseLockPath(path)
	if err != nil {
		return nil, err
	}
	p := tree.root
	for _, v := range ps {
		next, ok := p.Sub[v]
		if !ok {
			return nil, errors.New("get node failed: path not found:" + path)
		}
		p = next
	}
	return p, nil
}

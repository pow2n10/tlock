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

func (tree *Tree) Lock(path string, wait time.Duration) error {
	ps, err := ParseLockPath(path)
	if err != nil {
		return err
	}
	ch := make(chan int, 0)
	go func() {
		p := tree.root
		for _, v := range ps {
			p.Mutex.Lock()
			next, ok := p.Get(v)
			if !ok {
				newNode := NewNode(v)
				p.Append(v, newNode)
				next = newNode
			}
			p.Mutex.Unlock()
			p = next
		}
		p.Mutex.Lock()
		ch <- 1
	}()
	select {
	case <-ch:
		return nil
	case <-time.After(wait):
		return errors.New("wait lock timeout")
	}
}

func (tree *Tree) Unlock(path string) error {

	ps, err := ParseLockPath(path)
	if err != nil {
		return err
	}
	p := tree.root
	for _, v := range ps {
		node, ok := p.Get(v)
		if !ok {
			return errors.New("element " + v + " not found")
		}
		p = node
	}
	p.Mutex.Unlock()
	return nil
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
		next, ok := p.Get(v)
		if !ok {
			return nil, errors.New("get node failed: path not found:" + path)
		}
		p = next
	}
	return p, nil
}

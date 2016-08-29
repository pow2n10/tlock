package treelock

import (
	"sync"
	"time"
)

type Node struct {
	sync.Mutex
	Name    string
	Sub     map[string]*Node
	Meta    string
	Expires time.Time
}

func NewNode(name string) *Node {
	node := new(Node)
	node.Name = name
	node.Sub = make(map[string]*Node, 0)
	return node
}

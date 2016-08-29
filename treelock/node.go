package treelock

import (
	"sync"
	"time"

	"github.com/pow2n10/tlock/safemap"
)

type Node struct {
	Mutex         sync.Mutex
	internalMutex sync.Mutex
	Name          string
	Sub           *safemap.SafeMap
	Meta          string
	Expires       time.Time
}

func NewNode(name string) *Node {
	node := new(Node)
	node.Name = name
	node.Sub = safemap.NewSafeMap()
	return node
}

func (node *Node) Get(name string) (*Node, bool) {

	v, ok := node.Sub.Get(name)
	if ok {
		return v.(*Node), ok
	}
	return nil, ok
}

func (node *Node) Append(name string, newNode *Node) {
	node.Sub.Set(name, newNode)
}

func (node *Node) Delete(name string) {
	node.Sub.Delete(name)
}

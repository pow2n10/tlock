package treelock

import (
	"testing"
	"time"
)

func TestSimpleLock(t *testing.T) {

	tree := NewTree()

	err := tree.Lock("lockA", 30*time.Second)

	if err != nil {
		t.Error(err)
	}

	err = tree.Unlock("lockA")

	if err != nil {
		t.Error(err)
	}

	err = tree.Lock("lockA", 5*time.Second)

	if err != nil {
		t.Error(err)
	}
}

func TestTreeLock(t *testing.T) {

	var path string
	tree := NewTree()

	path = "path"

	err := tree.Lock(path, 1*time.Second)

	if err != nil {
		t.Error(err)
	}

	path = "path/to/A"

	err = tree.Lock(path, 1*time.Second)

	if err != nil {
		t.Error(err)
	}

}

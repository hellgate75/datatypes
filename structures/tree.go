package structures

import (
	"fmt"
)

func NewTree() *Tree {
	return &Tree{}
}

func TreeFromRoot(root *RNode) *Tree {
	return &Tree{
		root: root,
	}
}

type Tree struct {
	root *RNode
}

func (t *Tree) Add(node *RNode, value interface{}) *RNode {
	if node == nil {
		if t.root != nil {
			return nil
		}
		t.root = &RNode{
			Value: value,
		}
		return t.root
	} else {
		var newNode *RNode = &RNode{
			Value: value,
		}
		node.AddNext(newNode)
		return newNode
	}
	return nil
}

func (t *Tree) FindByValue(value interface{}) []*RPath {
	if t.root != nil {
		return t.root.FindByValue(value)
	}
	return make([]*RPath, 0)
}

func (t *Tree) FindByNode(node *RNode) []*RPath {
	if t.root != nil {
		return t.root.FindByNode(node)
	}
	return make([]*RPath, 0)
}

func (t *Tree) Root() *RNode {
	return t.root
}

func (t *Tree) String() string {
	if t.root != nil {
		return fmt.Sprintf("{root : %v}", t.root) //fmt.Sprint(q.root.Flatten())
	}
	return "<empty>"
}

package structures

import "fmt"

type SampleNode struct {
	Value int
}

func (n *SampleNode) String() string {
	return fmt.Sprint(n.Value)
}

type SNode struct {
	Value interface{}
	Child *SNode
}

func (sNode *SNode) Add(node *SNode) {
	var parentNode *SNode = sNode
	for parentNode.Child != nil {
		parentNode = parentNode.Child
	}
	parentNode.Child = node
}

func (sNode *SNode) RemoveFromLast() *SNode {
	var parentNode *SNode = sNode
	var previousNode *SNode
	var lastNode *SNode
	for parentNode.Child != nil {
		previousNode = parentNode
		parentNode = parentNode.Child
	}
	if parentNode.Child == nil {
		previousNode.Child = nil
		lastNode = parentNode
	}
	return lastNode
}

func (sNode *SNode) Flatten() []*SNode {
	var flattenList []*SNode = make([]*SNode, 0)
	var parentNode *SNode = sNode
	flattenList = append(flattenList, parentNode)
	for parentNode.Child != nil {
		parentNode = parentNode.Child
		flattenList = append(flattenList, parentNode)
	}
	if parentNode.Child != nil {
		flattenList = append(flattenList, parentNode.Child)
	}
	return flattenList
}

func (sNode *SNode) HasChildren() bool {
	return sNode.Child != nil
}

func (sNode *SNode) Count() int64 {
	var count int64 = 0
	var parentNode *SNode = sNode
	for parentNode.Child != nil {
		parentNode = parentNode.Child
		count++
	}
	count++
	return count
}

type RNode struct {
	Value    interface{}
	Parent   *RNode
	Children []*RNode
}

func (rNode *RNode) AddNext(node *RNode) {
	rNode.Children = append(rNode.Children, node)
}

// Number of Node Elements siblings current
func (rNode *RNode) LevelSize() int {
	return len(rNode.Children)
}

// Node Elements siblings of current at given position or nil
func (rNode *RNode) GetNodeAt(position int) *RNode {
	var node *RNode
	if position < rNode.LevelSize() {
		node = rNode.Children[position]
	}
	return node
}

// Recursive Node Size
func (rNode *RNode) Size() int64 {
	var size int64 = int64(rNode.LevelSize())
	for i := 0; i < rNode.LevelSize(); i++ {
		size += int64(rNode.Children[i].Size())
	}
	return size + 1 //Adding current node in size
}

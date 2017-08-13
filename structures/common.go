package structures

import (
	"fmt"
	"reflect"
)

type SampleNode struct {
	Value int `json:"Value" xml:"Value" mandatory:"yes" descr:"Linked Node Value" type:"text"`
}

func (n *SampleNode) String() string {
	return fmt.Sprint(n.Value)
}

type SNode struct {
	Value interface{}
	Child *SNode
}

func (sNode *SNode) String() string {
	return fmt.Sprintf("{Value: %v, Child: %v}", sNode.Value, sNode.Child)
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

type RPath struct {
	Matches []*RNode
	Path    []*RNode
}

func (p *RPath) NodesString(nodes []*RNode) string {
	var out string = "["
	for i := 0; i < len(nodes); i++ {
		return fmt.Sprintf("{Value: %v, Children: %v}", nodes[i].Value, nodes[i].Children)

	}
	return out + "]"
}

func (p *RPath) String() string {
	return fmt.Sprintf("{Matches: %v, Paths: %v}", p.NodesString(p.Matches), p.NodesString(p.Path))
}

func (p *RPath) ContainsValue(value interface{}) bool {
	for i := 0; i < len(p.Path); i++ {
		if reflect.DeepEqual(p.Path[i].Value, value) {
			return true
		}
	}
	return false
}

func (p *RPath) FindNodes(value interface{}) []*RNode {
	var results []*RNode = make([]*RNode, 0)
	for i := 0; i < len(p.Path); i++ {
		if reflect.DeepEqual(p.Path[i].Value, value) {
			results = append(results, p.Path[i])
		}
	}
	return results
}

func (p *RPath) ContainsNode(node *RNode) bool {
	if node == nil {
		return false
	}
	for i := 0; i < len(p.Path); i++ {
		if p.Path[i] == node {
			return true
		}
	}
	return false
}

func (p *RPath) Overlaps(path *RPath) bool {
	var seekingPath *RPath = p
	var matchingPath *RPath = path
	if len(p.Path) < len(path.Path) {
		seekingPath = path
		matchingPath = p
	}
	var hops int
	for i := 0; i < len(seekingPath.Path); i++ {
		if reflect.DeepEqual(seekingPath.Path[i].Value, matchingPath.Path[i]) {
			hops++
		} else {
			break
		}
	}
	return hops > 0
}

func (p *RPath) Contains(path *RPath) bool {
	if len(p.Path) < len(path.Path) {
		return false
	}
	var seekingPath *RPath = p
	var matchingPath *RPath = path
	for i := 0; i < len(seekingPath.Path); i++ {
		if !reflect.DeepEqual(seekingPath.Path[i].Value, matchingPath.Path[i]) {
			return false
		}
	}
	return true
}

func (p *RPath) Equals(path *RPath) bool {
	if len(p.Path) != len(path.Path) {
		return false
	}
	var seekingPath *RPath = p
	var matchingPath *RPath = path
	for i := 0; i < len(seekingPath.Path); i++ {
		if !reflect.DeepEqual(seekingPath.Path[i].Value, matchingPath.Path[i]) {
			return false
		}
	}
	return true
}

type RNode struct {
	Value    interface{}
	Parent   *RNode
	Children []*RNode
	Paths    []*RPath
}

func (rNode *RNode) String() string {
	return fmt.Sprintf("{Value: %v, Children: %v, Paths: %v}", rNode.Value, rNode.Children, rNode.Paths)
}

func (rNode *RNode) updatePath(node *RNode) {
	var indexes []int = make([]int, 0)
	//Removing any path from node
	for i := 0; i < len(rNode.Paths); i++ {
		if rNode.Paths[i].ContainsValue(node.Value) {
			indexes = append(indexes, i)
		}
	}
	for i := 0; i < len(indexes); i++ {
		rNode.Paths = rNode.Paths[:indexes[i]]
		rNode.Paths = append(rNode.Paths, rNode.Paths[indexes[i]:]...)
	}
	//Rebuilding node paths
	var nodePathLen int = len(node.Paths)
	if nodePathLen > 0 {
		for i := 0; i < nodePathLen; i++ {
			var newPath *RPath = &RPath{}
			newPath.Path = append(newPath.Path, rNode)
			newPath.Path = append(newPath.Path, node.Paths[i].Path...)
			rNode.Paths = append(rNode.Paths, newPath)
		}
	} else {
		var newPath *RPath = &RPath{}
		newPath.Path = append(newPath.Path, rNode)
		newPath.Path = append(newPath.Path, node)
		rNode.Paths = append(rNode.Paths, newPath)
	}
	if node.Parent != nil {
		node.Parent.updatePath(rNode)
	}
}

func (rNode *RNode) removeAt(position int) *RNode {
	var node *RNode
	if position < rNode.LevelSize() {
		node = rNode.GetNodeAt(position)
		if node != nil {
			rNode.Children = rNode.Children[:position]
			if position < rNode.LevelSize()-1 {
				rNode.Children = append(rNode.Children, rNode.Children[position:]...)
			}
			if rNode.Parent != nil {
				rNode.Parent.updatePath(rNode)
			}
		}
	}
	return node
}

func (rNode *RNode) AddNext(node *RNode) {
	node.Parent = rNode
	rNode.Children = append(rNode.Children, node)
	var pathList []*RNode = make([]*RNode, 0)
	pathList = append(pathList, node)
	var path *RPath = &RPath{
		Path: pathList,
	}
	rNode.Paths = append(rNode.Paths, path)
	rNode.updatePath(node)
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

// List current node siblings
func (rNode *RNode) ListChildren() []*RNode {
	return rNode.Children
}

// Recursive Nodes Size
func (rNode *RNode) Size() int64 {
	var size int64 = int64(rNode.LevelSize())
	for i := 0; i < rNode.LevelSize(); i++ {
		size += int64(rNode.Children[i].Size())
	}
	return size + 1 //Adding current node in size
}

//Find path by Value
func (rNode *RNode) FindByValue(value interface{}) []*RPath {
	var resultPaths []*RPath = make([]*RPath, 0)
	for i := 0; i < len(rNode.Paths); i++ {
		if rNode.Paths[i].ContainsValue(value) {
			var resultPath RPath = *rNode.Paths[i]
			var results []*RNode = resultPath.FindNodes(value)
			if len(results) > 0 {
				resultPath.Matches = append(resultPath.Matches, results...)
			}
			resultPaths = append(resultPaths, &resultPath)
		}
	}
	return resultPaths
}

//Find path by Value
func (rNode *RNode) FindByNode(node *RNode) []*RPath {
	var resultPaths []*RPath = make([]*RPath, 0)
	for i := 0; i < len(rNode.Paths); i++ {
		if rNode.Paths[i].ContainsNode(node) {
			var resultPath RPath = *rNode.Paths[i]
			resultPath.Matches = make([]*RNode, 0)
			resultPath.Matches = append(resultPath.Matches, node)
			resultPaths = append(resultPaths, &resultPath)
		}
	}
	return resultPaths
}

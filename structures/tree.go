package structures

type Tree struct {
	root *RNode
}


func (t *Tree) String() string {
	if t.root != nil {
		return "..."//fmt.Sprint(q.root.Flatten())
	}
	return  "<empty>"
}

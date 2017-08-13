package structures

import "fmt"

// NewQueue returns a new queue
func NewQueue() *Queue {
	return &Queue{}
}
// NewQueue returns a new queue
func QueueFromRoot(root *SNode) *Queue {
	return &Queue{
		root: root,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	root *SNode
}

// Push adds a node to the queue.
func (q *Queue) Push(n interface{}) {
	if q.root == nil {
		q.root = &SNode{
			Value: n,
		}
	} else {
		q.root.Add(&SNode{
			Value: n,
		})
	}
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() interface{} {
	var value interface{}
	if q.root != nil {
		value = q.root.Value
		q.root = q.root.Child
	}
	return value
}

// Return queue size
func (q *Queue) Size() int64 {
	var size int64
	if q.root != nil {
		size = q.root.Count()
	}
	return size
}

// Return queue Root Element
func (q *Queue) Root() *SNode {
	return q.root
}

// list all nodes from the queue.
func (s *Queue) List() []interface{} {
	if s.root != nil {
		var flattenNodes []*SNode = s.root.Flatten()
		var flattenValues []interface{} = make([]interface{}, len(flattenNodes))
		for i := 0; i < len(flattenNodes); i++ {
			flattenValues[i] = flattenNodes[i].Value
		}
		return flattenValues
	}
	return make([]interface{}, 0)
}

func (q *Queue) String() string {
	if q.root != nil {
		return fmt.Sprint(q.List())
	}
	return "<empty>"
}

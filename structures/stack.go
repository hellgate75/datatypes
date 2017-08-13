package structures

import "fmt"

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	root *SNode
}

// Push adds a node to the stack.
func (s *Stack) Push(n interface{}) {
	if s.root == nil {
		s.root = &SNode{
			Value: n,
		}
	} else {
		s.root.Add(&SNode{
			Value: n,
		})
	}
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() interface{} {
	var value interface{}
	if s.root != nil {
		if s.root.HasChildren() {
			var node *SNode = s.root.RemoveFromLast()
			if node != nil {
				value = node.Value
			}
		} else {
			value = s.root.Value
			s.root = nil
		}
	}
	return value
}

// list all nodes from the stack.
func (s *Stack) List() []interface{} {
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


// Return stack size
func (q *Stack) Size() int64 {
	var size int64
	if q.root != nil {
		size = q.root.Count()
	}
	return size
}


func (s *Stack) String() string {
	if s.root != nil {
		return fmt.Sprint(s.List())
	}
	return  "<empty>"
}

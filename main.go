package main

import (
	"fmt"
	"github.com/hellgate75/datatypes/structures"
)

func main() {
	s := structures.NewStack()
	s.Push(&structures.SampleNode{0})
	s.Push(&structures.SampleNode{2})
	s.Push(&structures.SampleNode{3})
	fmt.Println("Stack Size:", s.Size())
	fmt.Println("Stack :", s.List())
	fmt.Println(s.Pop(), s.Pop(), s.Pop())
	s.Push(&structures.SampleNode{7})
	fmt.Println("Stack Size:", s.Size())
	fmt.Println("Stack :", s.List())
	fmt.Println(s.Pop(), s.Pop())
	fmt.Println("Stack :", s.List())
	fmt.Println()

	q := structures.NewQueue()
	q.Push(&structures.SampleNode{4})
	q.Push(&structures.SampleNode{5})
	q.Push(&structures.SampleNode{6})
	fmt.Println("Queue Size:", q.Size())
	fmt.Println("Queue :", q.List())
	fmt.Println(q.Pop(), q.Pop(), q.Pop())
	q.Push(&structures.SampleNode{4})
	fmt.Println("Queue Size:", q.Size())
	fmt.Println("Queue :", q.List())
	fmt.Println(q.Pop(), q.Pop())
	fmt.Println("Queue :", q.List())
}

package main

import (
	"fmt"
	"github.com/hellgate75/datatypes/structures"
	"github.com/hellgate75/datatypes/utils"
)

func main() {
	s := structures.NewStack()
	s.Push(&structures.SampleNode{0})
	s.Push(&structures.SampleNode{2})
	s.Push(&structures.SampleNode{3})
	fmt.Println("Stack Size:", s.Size())
	fmt.Println("Stack :", s.List())
	fmt.Println("Saving Stack")
	sErr := utils.SaveStack(*s, "spool/stack.yml", "yaml")
	fmt.Println(fmt.Sprintf("Error: %v", sErr))
	fmt.Println("Removing :", s.Pop(), s.Pop(), s.Pop())
	s.Push(&structures.SampleNode{7})
	fmt.Println("Stack Size:", s.Size())
	fmt.Println("Stack :", s.List())
	fmt.Println("Removing :", s.Pop(), s.Pop())
	fmt.Println("Stack :", s.List())
	fmt.Println()
	fmt.Println("Reload Stack")
	var s2 *structures.Stack
	var err error
	s2, err = utils.LoadStack("spool/stack.yml", "yaml", func(i interface{}) interface{} {
		var id map[interface{}]interface{} = i.(map[interface{}]interface{})
		val, _ := id["value"]
		intVal, _ := utils.StringToInt(fmt.Sprintf("%d", val))
		return &structures.SampleNode{
			Value: intVal,
		}
	})
	fmt.Println(fmt.Sprintf("Error %v", err))
	fmt.Println("Stack #2 :", s2.List())
	fmt.Println()

	q := structures.NewQueue()
	q.Push(&structures.SampleNode{4})
	q.Push(&structures.SampleNode{5})
	q.Push(&structures.SampleNode{6})
	fmt.Println("Queue Size:", q.Size())
	fmt.Println("Queue :", q.List())
	fmt.Println("Saving Queue")
	sErr = utils.SaveQueue(*q, "spool/queue.yml", "yaml")
	fmt.Println(fmt.Sprintf("Error: %v", sErr))
	fmt.Println("Removing :", q.Pop(), q.Pop(), q.Pop())
	q.Push(&structures.SampleNode{4})
	fmt.Println("Queue Size:", q.Size())
	fmt.Println("Queue :", q.List())
	fmt.Println("Removing :", q.Pop(), q.Pop())
	fmt.Println("Queue :", q.List())
	fmt.Println()
	fmt.Println("Reload Queue")
	var q2 *structures.Queue
	q2, err = utils.LoadQueue("spool/queue.yml", "yaml", func(i interface{}) interface{} {
		var id map[interface{}]interface{} = i.(map[interface{}]interface{})
		val, _ := id["value"]
		intVal, _ := utils.StringToInt(fmt.Sprintf("%d", val))
		return &structures.SampleNode {
			Value: intVal,
		}
	})
	fmt.Println(fmt.Sprintf("Error %v", err))
	fmt.Println("Queue #2 :", q2.List())
	fmt.Println()

	t := structures.NewTree()
	tn0 := &structures.SampleNode{0}
	tn1 := &structures.SampleNode{1}
	tn2 := &structures.SampleNode{2}
	tn3 := &structures.SampleNode{3}
	tn11 := &structures.SampleNode{4}
	tn12 := &structures.SampleNode{5}
	tn21 := &structures.SampleNode{6}
	tn22 := &structures.SampleNode{7}
	tn31 := &structures.SampleNode{8}
	tn32 := &structures.SampleNode{9}
	rootNode := t.Add(nil, tn0)
	tn1Node := t.Add(rootNode, tn1)
	tn2Node := t.Add(rootNode, tn2)
	tn3Node := t.Add(rootNode, tn3)
	_ = t.Add(tn1Node, tn11)
	_ = t.Add(tn1Node, tn12)
	_ = t.Add(tn2Node, tn21)
	_ = t.Add(tn2Node, tn22)
	_ = t.Add(tn3Node, tn31)
	_ = t.Add(tn3Node, tn32)
	fmt.Println("Tree")
	fmt.Println("tree #1 : ", t)
	fmt.Println("tree #1 : ", fmt.Sprintf("%s", utils.GetTreeYaml(*t)))
	fmt.Println()
	fmt.Println("Saving Tree")
	sErr = utils.SaveTree(*t, "spool/tree.yml", "yaml")
	fmt.Println(fmt.Sprintf("Error: %v", sErr))
	fmt.Println()
	fmt.Println("Loading Tree")
	t2, err := utils.LoadTree("spool/tree.yml", "yaml", func(i interface{}) interface{} {
		var id map[interface{}]interface{} = i.(map[interface{}]interface{})
		val, _ := id["value"]
		intVal, _ := utils.StringToInt(fmt.Sprintf("%d", val))
		return &structures.SampleNode {
			Value: intVal,
		}
	})
	fmt.Println(fmt.Sprintf("Error: %v", err))
	fmt.Println("tree #1 : ", t2)
	fmt.Println("tree #2 : ", fmt.Sprintf("%s", utils.GetTreeYaml(*t2)))
}

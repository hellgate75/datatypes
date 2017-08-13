package utils

import (
	"encoding/json"
	"errors"
	"github.com/hellgate75/datatypes/structures"
	"io/ioutil"
	"encoding/xml"
	"gopkg.in/yaml.v2"
	"fmt"
)

type SimpleNode struct {
	Node	interface{} `json:"Node" xml:"Node" mandatory:"yes" descr:"Simple Node Value" type:"text"`
	Child *SimpleNode  `json:"Child" xml:"Child" mandatory:"no" descr:"Simple Node Child" type:"text"`
}

type LinkedNode struct {
	Node	   	interface{}  	`json:"Node" xml:"Node" mandatory:"yes" descr:"Linked Node Value" type:"text"`
	Children 	[]*LinkedNode 	`json:"Children" xml:"Children" mandatory:"no" descr:"Linked Node Child" type:"text"`
}

func (ln *LinkedNode) String() string {
	return fmt.Sprintf("{Node: %v, Children: %v}", ln.Node, ln.Children )
}

func attachSimpleNodeRecursive(parent *SimpleNode, node *structures.SNode) {
	if node != nil {
		var simpleNode *SimpleNode = &SimpleNode{}
		simpleNode.Node = node.Value
		attachSimpleNodeRecursive(simpleNode, node.Child)
		parent.Child = simpleNode
	}
}

func attachSNodeRecursive(parent *structures.SNode, node *SimpleNode, convert func(interface{}) (interface{})) {
	if node != nil {
		var sNode *structures.SNode = &structures.SNode{}
		sNode.Value = convert(node.Node)
		attachSNodeRecursive(sNode, node.Child, convert)
		parent.Child = sNode
	}
}

func attachLinkedNodeRecursive(parent *LinkedNode, nodes []*structures.RNode) {
	if len(nodes) > 0 {
		for i := 0; i<len(nodes); i++ {
			var linkedNode *LinkedNode = &LinkedNode{
				Node: nodes[i].Value,
			}
			attachLinkedNodeRecursive(linkedNode, nodes[i].Children)
			parent.Children = append(parent.Children, linkedNode)
		}
	}
}

func attachRNodeRecursive(parent *structures.RNode, nodes []*LinkedNode, convert func(interface{}) (interface{})) {
	if len(nodes) > 0 {
		for i := 0; i < len(nodes); i++ {
			var rNode *structures.RNode = &structures.RNode{
				Value: convert(nodes[i].Node),
				Parent: parent,
			}
			parent.AddNext(rNode)
			attachRNodeRecursive(rNode, nodes[i].Children, convert)
		}
	}
	fmt.Println("Paths :", fmt.Sprintf("%v", parent.Paths))
}

func buildQueueNode(queue structures.Queue) *SimpleNode {
	var root *structures.SNode = queue.Root()
	var simpleNode *SimpleNode
	if root != nil {
		simpleNode = &SimpleNode{}
		simpleNode.Node = root.Value
		attachSimpleNodeRecursive(simpleNode, root.Child)
	}
	return simpleNode
}

func buildNodeQueue(simpleNode *SimpleNode, convert func(interface{}) (interface{})) *structures.Queue {
	var root *structures.SNode
	var queue *structures.Queue
	if simpleNode != nil {
		root= &structures.SNode{}
		root.Value = convert(simpleNode.Node)
		attachSNodeRecursive(root, simpleNode.Child, convert)
		queue=structures.QueueFromRoot(root)
	}
	return queue

}

func buildTreeNode(tree structures.Tree) *LinkedNode {
	var root *structures.RNode = tree.Root()
	var linkedNode *LinkedNode
	if root != nil {
		linkedNode = &LinkedNode{
			Node: root.Value,
		}
		attachLinkedNodeRecursive(linkedNode, root.Children)
	}
	return linkedNode
}

func buildNodeTree(linkedNode *LinkedNode, convert func(interface{}) (interface{})) *structures.Tree {
	var root *structures.RNode
	var tree *structures.Tree
	if linkedNode != nil {
		root=&structures.RNode{}
		root.Value = convert(linkedNode.Node)
		attachRNodeRecursive(root, linkedNode.Children, convert)
		tree=structures.TreeFromRoot(root)
	}
	return tree
}


func buildStackNode(stack structures.Stack) *SimpleNode {
	var root *structures.SNode = stack.Root()
	var simpleNode *SimpleNode
	if root != nil {
		simpleNode = &SimpleNode{}
		simpleNode.Node = root.Value
		attachSimpleNodeRecursive(simpleNode, root.Child)
	}
	return simpleNode
}
func buildNodeStack(simpleNode *SimpleNode, convert func(interface{}) (interface{})) *structures.Stack {
	var root *structures.SNode
	var stack *structures.Stack
	if simpleNode != nil {
		root= &structures.SNode{}
		root.Value = convert(simpleNode.Node)
		attachSNodeRecursive(root, simpleNode.Child, convert)
		stack=structures.StackFromRoot(root)
	}
	return stack

}


func SaveQueue(queue structures.Queue, file string, format string) error {
	var err error
	var rootNode *SimpleNode = buildQueueNode(queue)
	if rootNode != nil {
		return  ExportStructureToFile(file, format, rootNode)
	} else {
		err = errors.New("Specified Queue has no values")
	}
	return err
}

func LoadQueue(file string, format string, convert func(interface{}) (interface{})) (*structures.Queue, error) {
	var err error
	if !ExistsFile(file) {
		return nil, errors.New("File " + file + " doesn't exist!!")
	}
	if format != "json" && format != "xml" && format != "json" {
		format="yaml"
	}
	byteArray, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var element *SimpleNode=&SimpleNode{}
	if format == "json" {
		err = json.Unmarshal(byteArray, element)
	} else if format == "yaml" {
		err = yaml.Unmarshal(byteArray, element)
	} else {
		err = xml.Unmarshal(byteArray, element)
	}
	if err != nil {
		return nil, err
	}
	var queue *structures.Queue = buildNodeQueue(element, convert)
	if queue == nil {
		return queue, errors.New("Unable to convert Queue")
	}
	return queue, err
}

func SaveStack(stack structures.Stack, file string, format string) error {
	var err error
	var rootNode *SimpleNode = buildStackNode(stack)
	if rootNode != nil {
		return  ExportStructureToFile(file, format, rootNode)
	} else {
		err = errors.New("Specified Stack has no values")
	}
	return err
}


func LoadStack(file string, format string, convert func(interface{}) (interface{}) ) (*structures.Stack, error) {
	var err error
	if !ExistsFile(file) {
		return nil, errors.New("File " + file + " doesn't exist!!")
	}
	if format != "json" && format != "xml" && format != "json" {
		format="yaml"
	}
	byteArray, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var element *SimpleNode=&SimpleNode{}
	if format == "json" {
		err = json.Unmarshal(byteArray, element)
	} else if format == "yaml" {
		err = yaml.Unmarshal(byteArray, element)
	} else {
		err = xml.Unmarshal(byteArray, element)
	}
	if err != nil {
		return nil, err
	}
	var stack *structures.Stack = buildNodeStack(element, convert)
	if stack == nil {
		return stack, errors.New("Unable to convert Stack")
	}
	return stack, err
}

func GetTreeYaml(tree structures.Tree) []byte {
	var node *LinkedNode = buildTreeNode(tree)
	bytes, err := yaml.Marshal(node)
	if err != nil {
		return []byte{}
	}
	return bytes
}

func SaveTree(tree structures.Tree, file string, format string) error {
	var err error
	var rootNode *LinkedNode = buildTreeNode(tree)
	if rootNode != nil {
		return  ExportStructureToFile(file, format, rootNode)
	} else {
		err = errors.New("Specified Tree has no values")
	}
	return err
}

func LoadTree(file string, format string, convert func(interface{}) (interface{}) ) (*structures.Tree, error) {
	var err error
	if !ExistsFile(file) {
		return nil, errors.New("File " + file + " doesn't exist!!")
	}
	if format != "json" && format != "xml" && format != "json" {
		format="yaml"
	}
	byteArray, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var element *LinkedNode=&LinkedNode{}
	if format == "json" {
		err = json.Unmarshal(byteArray, element)
	} else if format == "yaml" {
		err = yaml.Unmarshal(byteArray, element)
	} else {
		err = xml.Unmarshal(byteArray, element)
	}
	if err != nil {
		return nil, err
	}
	var tree *structures.Tree = buildNodeTree(element, convert)
	if tree == nil {
		return tree, errors.New("Unable to convert Tree")
	}
	return tree, err
}


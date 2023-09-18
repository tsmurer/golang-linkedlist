package main

import "fmt"

type Node struct {
	value    string
	nextNode *Node
}

func main() {
	firstNode := &Node{
		value:    "first node!",
		nextNode: nil,
	}

	printAllNodesFrom(*firstNode)
	firstNode.pushNewNode("second")
	printAllNodesFrom(*firstNode)
	firstNode.pushNewNode("third")
	printAllNodesFrom(*firstNode)
	firstNode.popLastNode()
	printAllNodesFrom(*firstNode)

}

func (node *Node) pushNewNode(value string) {
	newNode := &Node{
		value:    value,
		nextNode: nil,
	}
	for node.nextNode != nil {
		node = node.nextNode
	}
	node.nextNode = newNode
}

func (node *Node) popLastNode() {
	for node.nextNode != nil {
		if node.nextNode.nextNode == nil {
			node.nextNode = nil
		} else {
			node = node.nextNode
		}
	}
}

func printAllNodesFrom(node Node) {
	fmt.Println("Printing all nodes from the first:")
	fmt.Println(node.value)
	for node.nextNode != nil {
		node = *node.nextNode
		fmt.Println(node.value)
	}

}
